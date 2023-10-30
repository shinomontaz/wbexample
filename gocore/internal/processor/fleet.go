package processor

import (
	"context"
	"encoding/json"
	"time"
	"ws-core/internal/common"
	"ws-core/pkg/utils"

	"github.com/sirupsen/logrus"
)

const KEY_TRUCKS string = "laravel_database_trucks"
const KEY_POSITIONS string = "laravel_database_positions"

var (
	truck_counter int
	trucks        []*common.Truck
	positions     map[int]common.Coord
	currArea      common.Area
	stop_chan     chan struct{}
	chan_started  bool
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

func StartFleet(ctx context.Context) {
	if chan_started {
		return
	}
	ticker := time.NewTicker(2 * time.Second)
	chan_started = true
	stop_chan = Run(ticker, ctx)
}

func StopFleet() {
	if !chan_started {
		return
	}
	stop_chan <- struct{}{}
	chan_started = false
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

	speed := 27.7 // meters per second
	Delta := speed * dt / float64(dist)

	t.Coord.Lat += Delta * (t.Destination.Lat - currPos.Lat)
	t.Coord.Long += Delta * (t.Destination.Long - currPos.Long)
}
