package core

import (
	"encoding/json"
	"ws-core/internal/common"
	"ws-core/internal/processor"

	"github.com/sirupsen/logrus"
)

func Run() {

	defer mqConn.Close()
	ch, err := mqConn.Channel()
	if err != nil {
		logrus.Fatal(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		mqcfg.Listenqueue, // queue
		"",                // consumer
		false,             // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	if err != nil {
		logrus.Fatal(err)
	}

	var forever chan struct{}

	var mess common.Message
	go func() {
		for d := range msgs {
			logrus.Debug("Received a message: %s", d.Body)
			// split messages by types to processing
			mess = common.Message{}
			err := json.Unmarshal([]byte(d.Body), &mess)
			if err != nil {
				logrus.Fatal(err)
			}
			processor.Do(mess, []byte(d.Body))
		}
	}()

	logrus.Debug("Waiting for messages")
	<-forever
}
