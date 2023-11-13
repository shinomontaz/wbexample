package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"
	"ws-core/internal/common"
	"ws-core/pkg/distancer"
	"ws-core/pkg/utils"

	oddg "github.com/oddg/hungarian-algorithm"
	"github.com/sirupsen/logrus"
)

const KEY_TRUCKS string = "laravel_database_trucks"
const MAXTIME int = 60 * 60 * 24 // number of seconds in day

var (
	truck_counter int
	positions     map[int]common.Coord
	currArea      common.Area
	stop_chan     chan struct{}
	chan_started  bool
	speedUp       int
	lock          sync.Mutex
)

func GenerateFleet(m common.MessageGenerateFleet, ctx context.Context) error {
	var err error
	currArea = m.Area
	trucks = []*common.Truck{}
	for i := 0; i < m.Num; i++ {

		Long := m.Area.Min[0] + rs.Float64()*(m.Area.Max[0]-m.Area.Min[0])
		Lat := m.Area.Min[1] + rs.Float64()*(m.Area.Max[1]-m.Area.Min[1])

		trucks = append(trucks, &common.Truck{
			Id: truck_counter,
			Coord: common.Coord{
				Lat:  Lat,
				Long: Long,
			},
		})
		truck_counter += 1
	}

	var jtrucks []byte
	jtrucks, err = json.Marshal(trucks)
	if err != nil {
		return err
	}
	err = rDb.Set(ctx, KEY_TRUCKS, string(jtrucks), 0).Err()
	return err
}

func setDestination(t *common.Truck, a common.Area) {
	targetLong := a.Min[0] + rs.Float64()*(a.Max[0]-a.Min[0])
	targetLat := a.Min[1] + rs.Float64()*(a.Max[1]-a.Min[1])
	t.Destination = common.Coord{
		Lat:  targetLat,
		Long: targetLong,
	}
	t.HasDestination = true
}

func StartFleet(ctx context.Context) error {
	if chan_started && len(trucks) > 0 {
		return nil
	}
	ticker := time.NewTicker(2 * time.Second)
	chan_started = true
	stop_chan = Run(ticker, ctx)
	return nil
}

func StopFleet(ctx context.Context) error {
	return nil
}

func Run(ticker *time.Ticker, ctx context.Context) chan struct{} {
	stop := make(chan struct{}, 1)
	currtime := time.Now()
	go func() {
		for {
			select {
			case tm := <-ticker.C:
				diff := tm.Sub(currtime)
				update(diff.Seconds(), ctx)
				currtime = tm
			case <-stop:
				return
			}
		}
	}()

	return stop
}

func update(dt float64, ctx context.Context) {
	lock.Lock()
	defer lock.Unlock()

	// update fleet positions by time
	for _, t := range trucks {
		updateTruck(dt, t, currArea)
		// TODO use WaitGroup
	}

	var err error
	var jtrucks []byte
	jtrucks, err = json.Marshal(trucks)
	if err != nil {
		logrus.Error(err)
	}
	err = rDb.Set(ctx, KEY_TRUCKS, string(jtrucks), 0).Err()
	if err != nil {
		logrus.Error(err)
	}
}

func MakeAssignment(ctx context.Context) error {
	lock.Lock()
	defer lock.Unlock()

	ctrucks := make([]common.Coord, len(trucks))
	cwhs := make([]common.Coord, len(pts))
	for idx, t := range trucks {
		ctrucks[idx] = t.Coord
	}
	for idx, wh := range pts {
		cwhs[idx] = common.Coord{Lat: wh.Lat, Long: wh.Long}
	}

	fmatrix, err := distancer.GetMatrix(ctrucks, cwhs)
	if err != nil {
		return err
	}

	// Force matrix to be square
	fmatrix = makeSquare(fmatrix)

	// convert to ints
	n := len(fmatrix)
	matrix := make([][]int, n)
	for i, row := range fmatrix {
		matrix[i] = make([]int, n)
		for j, el := range row {
			matrix[i][j] = int(el)
		}
	}
	solution, err := oddg.Solve(matrix)
	if err != nil {
		return err
	}

	// Solution is an array `a` such that each row `i` is matched to column `a[i]`
	for truck_n, wh_n := range solution {
		if truck_n >= len(trucks) {
			break
		}
		if wh_n >= len(pts) {
			continue
		}

		trucks[truck_n].PointId = wh_n
		pts[wh_n].TruckId = truck_n
	}

	fmt.Println("solution: ", solution)
	fmt.Println("trucks: ", trucks)
	fmt.Println("pts: ", pts)

	var jtrucks []byte
	jtrucks, err = json.Marshal(trucks)
	if err != nil {
		logrus.Error(err)
	}
	err = rDb.Set(ctx, KEY_TRUCKS, string(jtrucks), 0).Err()
	if err != nil {
		logrus.Error(err)
	}

	return nil
}

func updateTruck(dt float64, t *common.Truck, a common.Area) {
	currPos := t.Coord
	if !t.HasDestination {
		setDestination(t, a)
		return
	}

	dist := utils.Dist(currPos, t.Destination)
	if dist < 50 {
		setDestination(t, a)
		return
	}

	speed := 10 * 27.7 // meters per second 100km per hour * 10
	Delta := speed * dt / float64(dist)

	t.Coord.Lat += Delta * (t.Destination.Lat - currPos.Lat)
	t.Coord.Long += Delta * (t.Destination.Long - currPos.Long)
}

func makeSquare(m [][]float32) [][]float32 {
	rows := len(m)
	if rows == 0 {
		return m
	}
	if rows == 1 {
		// find min and return
		min := m[0][0]
		for _, el := range m[0] {
			if el < min {
				min = el
			}
		}

		return [][]float32{{min}}
	}
	cols := len(m[0])
	if rows == cols {
		return m
	}
	max := m[0][0]
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if m[i][j] > max {
				max = m[i][j]
			}
		}
	}
	if rows < cols {
		for i := rows - 1; i < cols; i++ {
			m[i] = make([]float32, cols)
			for j := 0; j < cols; j++ {
				m[i][j] = max
			}
		}
		return m
	}

	for i := 0; i < rows; i++ {
		for j := cols; j < rows; j++ {
			m[i] = append(m[i], max)
		}
	}

	return m
}
