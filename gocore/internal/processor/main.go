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
)

func Init(client *redis.Client, source *rand.Rand) {
	rDb = client
	rs = source
}

func Do(m common.Message, orig []byte) error {
	ctx := context.Background()
	switch m.Type {
	case 2:
		logrus.Debug(fmt.Sprintf("Message of type 2: %v", m))
		mess := common.MessageAddPoint{}
		err := json.Unmarshal(orig, &mess)
		if err != nil {
			logrus.Warn(err)
			return err
		}
		return AddPoint(mess, ctx)
	case 3:
		logrus.Debug(fmt.Sprintf("Message of type 3: %v", m))
		mess := common.MessageGenerateFleet{}
		err := json.Unmarshal(orig, &mess)
		if err != nil {
			logrus.Warn(err)
			return err
		}
		logrus.Debug(fmt.Sprintf("Message of GenerateFleet: %v", mess))
		return GenerateFleet(mess, ctx)
	case 1:
		logrus.Debug(fmt.Sprintf("Message of type 1: %v", m))
		return StopFleet(ctx)
	case 4:
		logrus.Debug(fmt.Sprintf("Message of type 4: %v", m))
		return StartFleet(ctx)
	}

	return fmt.Errorf("unknown message type %v", m)
}
