package main

import (
	"flag"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	mqConn  *amqp.Connection
	qlisten string
)

func init() {
	//	purl := flag.String("url", "127.0.0.1:", "Rabbit url to connect")
	phost := flag.String("host", "127.0.0.1", "Rabbit host to connect")
	pport := flag.String("port", "5672", "Rabbit port to connect")
	puser := flag.String("user", "user", "Rabbit user to connect")
	ppass := flag.String("pass", "qwe", "Rabbit password to connect")
	pvhost := flag.String("vhost", "/", "Rabbit vhost")
	plistenqueue := flag.String("queue", "cmd", "Queue to listen")

	qlisten = *plistenqueue

	flag.Parse()

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", *puser, *ppass, *phost, *pport, *pvhost)

	var err error
	mqConn, err = amqp.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
}
