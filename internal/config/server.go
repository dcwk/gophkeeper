package config

import (
	"flag"

	"github.com/caarlos0/env"
)

type ServerConf struct {
	RunAddress             string `env:"RUN_ADDRESS"`
	DatabaseDSN            string `env:"DATABASE_URI"`
	DatabaseMaxConnections int    `env:"DATABASE_MAX_CONNECTIONS"`
	LogLevel               string `env:"LOG_LEVEL"`
	StoragePath            string `env:"STORAGE_PATH"`
}

func NewServerConf() (*ServerConf, error) {
	conf := &ServerConf{}

	flag.StringVar(&conf.LogLevel, "l", "info", "log level")
	flag.StringVar(&conf.LogLevel, "s", "/Users/ruslan.golovizin/Projects/practicum/gophkeeper/storage", "path to store binary data")
	flag.StringVar(&conf.RunAddress, "a", "localhost:8081", "server address")
	flag.StringVar(&conf.DatabaseDSN, "d", "postgres://postgres:123456@localhost:5432/gophkeeper", "database dsn")
	flag.IntVar(&conf.DatabaseMaxConnections, "mc", 100, "maximum number of concurrent connections")
	err := env.Parse(conf)
	if err != nil {
		return nil, err
	}

	return conf, err
}
