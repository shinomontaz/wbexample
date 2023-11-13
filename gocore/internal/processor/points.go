package processor

import (
	"context"
	"encoding/json"
	"ws-core/internal/common"
)

const KEY_POINTS string = "laravel_database_points"

var (
	wh_counter int
	err        error
)

func AddPoint(m common.MessageAddPoint, ctx context.Context) error {
	// if len(pts) == 0 {
	// 	jpoints, err := rDb.Get(ctx, KEY_POINTS).Result()
	// 	if err == redis.Nil {
	// 		jpoints = "[]"
	// 	} else if err != nil {
	// 		return err
	// 	}

	// 	logrus.Debug(fmt.Sprintf("Message of type 2: %v", jpoints))
	// 	err = json.Unmarshal([]byte(jpoints), &pts)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	pts = append(pts, common.Point{
		Id:   wh_counter,
		Lat:  m.Lat,
		Long: m.Long,
	})
	var (
		jpts []byte
		err  error
	)
	jpts, err = json.Marshal(pts)
	if err != nil {
		return err
	}
	wh_counter += 1
	err = rDb.Set(ctx, KEY_POINTS, string(jpts), 0).Err()
	return err
}
