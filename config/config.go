package config

import (
	"sync"
	"time"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	LogLevel            string        `env:"LOG_LEVEL"`
	DatabaseURL         string        `env:"DATABASE_URL"`
	PgTimeout           time.Duration `env:"PG_TIMEOUT"`
	PgConnAttempts      int           `env:"PG_CONN_ATTEMPTS"`
	PgHealthCheckPeriod time.Duration `env:"PG_HEALTH_CHECK_PERIOD"`
	PgMaxIdleTime       time.Duration `env:"PG_MAX_IDLE_TIME"`
	HTTPPort            string        `env:"HTTP_PORT"`
}

var (
	once   sync.Once
	config Config
)

const projectDirName = "app"

func Get() *Config {
	once.Do(func() {
		err := env.Parse(&config)
		if err != nil {
			panic(err)
		}
	})

	return &config
}
