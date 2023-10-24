package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/sirupsen/logrus"
)

type Rabbit struct {
	Host        string `env:"RABBIT_HOST" envDefault:"127.0.0.1"`
	Port        int    `env:"RABBIT_PORT" envDefault:5672`
	User        string `env:"RABBIT_USER" envDefault:"user"`
	Pass        string `env:"RABBIT_PASS" envDefault:"qwe"`
	Vhost       string `env:"RABBIT_VHOST" envDefault:"/"`
	Listenqueue string `env:"RABBIT_QUEUE" envDefault:"cmd"`
}

type Redis struct {
	Host string `env:"REDIS_HOST" envDefault:"127.0.0.1"`
	Port int    `env:"REDIS_PORT" envDefault:6379`
	Db   int    `env:"REDIS_DB" envDefault:1`
}

type Log struct {
	Level Level `env:"LOG_LEVEL" envDefault:"debug"`
}

type Level uint32

func (l *Level) UnmarshalText(text []byte) error {
	var (
		ll uint32
		ok bool
	)
	if ll, ok = LogLevelMapping[string(text)]; !ok {
		ll = uint32(logrus.DebugLevel)
	}
	*l = Level(ll)
	return nil
}

func Init(cfg interface{}) {
	if err := env.Parse(cfg); err != nil {
		logrus.Fatalf("parse config fatal: %s", err)
	}
}

var LogLevelMapping = map[string]uint32{
	"trace":   uint32(logrus.TraceLevel),
	"debug":   uint32(logrus.DebugLevel),
	"info":    uint32(logrus.InfoLevel),
	"warning": uint32(logrus.WarnLevel),
	"error":   uint32(logrus.ErrorLevel),
	"panic":   uint32(logrus.PanicLevel),
}
