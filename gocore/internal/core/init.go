package core

import (
	"fmt"
	"log"
	"os"

	"ws-core/internal/processor"
	"ws-core/pkg/config"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var (
	mqConn *amqp.Connection
	rDb    *redis.Client
	mqcfg  config.Rabbit
	rcfg   config.Redis
	logcfg config.Log
)

func init() {
	initLog()
	initMq()
	initRedis()
	processor.Init(rDb)
}

func initMq() {
	config.Init(&mqcfg)
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/%s", mqcfg.User, mqcfg.Pass, mqcfg.Host, mqcfg.Port, mqcfg.Vhost)

	logrus.Debug("rabbit connect url: ", url)
	var err error
	mqConn, err = amqp.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
}

func initRedis() {
	config.Init(&rcfg)
	logrus.Debug(fmt.Sprintf("redis config: %v", rcfg))
	rDb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", rcfg.Host, rcfg.Port),
		DB:   rcfg.Db,
	})
}

func initLog() {
	config.Init(&logcfg)
	logrus.SetLevel(logrus.Level(logcfg.Level))
	logrus.SetOutput(os.Stdout)
}
