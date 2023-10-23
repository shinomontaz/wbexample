package main

import (
	"encoding/json"
	"log"
	"ws-core/pkg/common"
	"ws-core/pkg/processor"
)

func main() {

	defer mqConn.Close()
	ch, err := mqConn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		qlisten, // queue
		"",      // consumer
		false,   // auto-ack
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)
	if err != nil {
		log.Fatal(err)
	}

	var forever chan struct{}

	var mess common.Message
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			// split messages by types to processing
			mess = common.Message{}
			err := json.Unmarshal([]byte(d.Body), &mess)
			if err != nil {
				log.Fatal(err)
			}
			processor.Process(mess)
		}
	}()

	log.Printf("Waiting for messages")
	<-forever
}
