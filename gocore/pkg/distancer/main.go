package distancer

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
	"ws-core/internal/common"

	"github.com/sirupsen/logrus"
)

var (
	ourl   string
	client http.Client
)

func Init(osrmUrl string) {
	ourl = osrmUrl
	client = http.Client{
		Timeout: time.Duration(2) * time.Second,
	}
}

// http request handler.
func GetMatrix(whs, drivers []common.Coord) ([][]float32, error) {
	scoords := []string{}
	sources := make([]int, len(drivers))
	destinations := make([]int, len(whs))
	i := 0
	for _, coord := range drivers {
		scoords = append(scoords, fmt.Sprintf("%f,%f", coord.Long, coord.Lat))
		sources[i] = i
		i += 1
	}
	j := 0
	for _, coord := range whs {
		scoords = append(scoords, fmt.Sprintf("%f,%f", coord.Long, coord.Lat))
		destinations[j] = i + j
		j += 1
	}
	coords := strings.Join(scoords, ";")

	fmt.Println(scoords)

	ssources := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sources)), ";"), "[]")
	sdestinations := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(destinations)), ";"), "[]")

	// make osrm call
	url := fmt.Sprintf("http://%s/table/v1/driving/%s?sources=%s&destinations=%s", ourl, coords, ssources, sdestinations)
	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second)
	defer cancelFn()
	var oresp osrmTableResponse
	err := doRequest(ctx, url, &oresp)
	if err != nil {
		return nil, err
	}
	if len(oresp.Durations) == 0 && len(sources) > 0 && len(destinations) > 0 {
		return nil, fmt.Errorf("empty response for non empty request")
	}
	return oresp.Durations, nil
}

func doRequest(ctx context.Context, url string, out interface{}) error {
	logrus.Debug("doRequest to ", url)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	resp, err := client.Do(req)

	//	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusBadRequest {
		return fmt.Errorf("error http status code %d", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read body: %w", err)
	}

	if err := json.Unmarshal(bytes, out); err != nil {
		return fmt.Errorf("failed to unmarshal body %q: %w", bytes, err)
	}

	return nil
}
