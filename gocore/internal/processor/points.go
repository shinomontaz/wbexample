package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"ws-core/internal/common"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

const KEY_POINTS string = "laravel_database_points"

var counter int

func AddPoint(m common.MessageAddPoint, ctx context.Context) error {
	//{"type":2,"long":37.32421874999999,"lat":55.76727004496223}
	jpoints, err := rDb.Get(ctx, KEY_POINTS).Result()
	if err == redis.Nil {
		jpoints = "[]"
	} else if err != nil {
		return err
	}

	logrus.Debug(fmt.Sprintf("Message of type 2: %v", jpoints))

	pts := []common.Point{}
	err = json.Unmarshal([]byte(jpoints), &pts)
	if err != nil {
		return err
	}

	pts = append(pts, common.Point{
		Id:   counter,
		Lat:  m.Lat,
		Long: m.Long,
	})
	var jpts []byte
	jpts, err = json.Marshal(pts)
	if err != nil {
		return err
	}
	counter += 1
	err = rDb.Set(ctx, KEY_POINTS, string(jpts), 0).Err()
	return err
}
