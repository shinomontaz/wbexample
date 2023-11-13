// Package processor
//
// process rabbit messages: add wh point, generate fleet, start fleet moving, pause fleet moving.
package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"ws-core/internal/common"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var (
	rDb *redis.Client
	rs  *rand.Rand

	trucks []*common.Truck
	pts    []common.Point
)

// init this package from outside
func Init(client *redis.Client, source *rand.Rand) {
	rDb = client
	rs = source
	pts = []common.Point{}
	trucks = []*common.Truck{}

	ctx := context.Background()
	jpoints, err := rDb.Get(ctx, KEY_POINTS).Result()
	if err == redis.Nil || err != nil {
		jpoints = "[]"
	}
	json.Unmarshal([]byte(jpoints), &pts)

	jtrucks, err := rDb.Get(ctx, KEY_TRUCKS).Result()
	if err == redis.Nil || err != nil {
		jtrucks = "[]"
	}
	json.Unmarshal([]byte(jtrucks), &trucks)
}

// main entry point to process rabbit message.
func Do(m common.Message, orig []byte) error {
	ctx := context.Background()
	switch m.Type {
	case 2: // add point
		logrus.Debug(fmt.Sprintf("Message of type 2: %v", m))
		mess := common.MessageAddPoint{}
		err := json.Unmarshal(orig, &mess)
		if err != nil {
			logrus.Warn(err)
			return err
		}
		return AddPoint(mess, ctx)
	case 3: // generate fleet
		logrus.Debug(fmt.Sprintf("Message of type 3: %v", m))
		mess := common.MessageGenerateFleet{}
		err := json.Unmarshal(orig, &mess)
		if err != nil {
			logrus.Warn(err)
			return err
		}
		logrus.Debug(fmt.Sprintf("Message of GenerateFleet: %v", mess))
		return GenerateFleet(mess, ctx)
	case 1: // pause fleet
		logrus.Debug(fmt.Sprintf("Message of type 1: %v", m))
		return StopFleet(ctx)
	case 4: // start fleet
		logrus.Debug(fmt.Sprintf("Message of type 4: %v", m))
		return StartFleet(ctx)
	case 5: // make assignment
		logrus.Debug(fmt.Sprintf("Message of type 5: %v", m))
		return MakeAssignment(ctx)
	}

	return fmt.Errorf("unknown message type %v", m)
}
