package processor

import (
	"encoding/json"
	"fmt"
	"ws-core/internal/common"

	"github.com/sirupsen/logrus"
)

func Do(m common.Message, orig []byte) error {
	switch m.T {
	case 2:
		mess := common.MessageAddPoint{}
		err := json.Unmarshal(orig, &mess)
		if err != nil {
			logrus.Warn(err)
			return err
		}
		return AddPoint(mess)
	}

	return fmt.Errorf("unknown message type %v", m)
}
