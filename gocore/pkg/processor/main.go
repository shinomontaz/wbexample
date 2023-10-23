package processor

import (
	"fmt"
	"ws-core/pkg/common"
)

func Do(m common.Message) error {
	switch m.t {
	case 2:
		return AddPoint(m)
	}

	return fmt.Errorf("unknown message type %v", m)
}
