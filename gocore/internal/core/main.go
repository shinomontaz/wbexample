package core

import (
	"encoding/json"
	"fmt"
	"ws-core/internal/common"
	"ws-core/internal/processor"

	amqp "github.com/rabbitmq/amqp091-go"
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

	limit := make(chan struct{}, 3)

	go func() {
		for d := range msgs {
			limit <- struct{}{}
			go func(d amqp.Delivery) {
				defer func() { <-limit }()                // last defer
				defer func(d amqp.Delivery) { Ack(d) }(d) // first defer

				logrus.Debug(fmt.Sprintf("Received a message: %s", string(d.Body)))
				// split messages by types to processing
				mess = common.Message{}
				err := json.Unmarshal(d.Body, &mess)
				if err != nil {
					logrus.Fatal(err)
				}
				logrus.Debug(fmt.Sprintf("Message umarshalled: %v", mess))
				err = processor.Do(mess, d.Body)
				if err != nil {
					//log this message to errors
					logrus.Error(err)
				}
			}(d)
		}
	}()

	logrus.Debug("Waiting for messages")
	<-forever
}

func Ack(d amqp.Delivery) {
	if err := d.Ack(false); err != nil { // Ack only this message
		logrus.Error(err)
	}
}
