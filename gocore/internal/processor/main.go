package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"ws-core/internal/common"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var rDb *redis.Client

func Init(client *redis.Client) {
	rDb = client
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
	}

	return fmt.Errorf("unknown message type %v", m)
}
