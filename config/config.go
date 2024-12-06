package config

import (
	"path/filepath"
	"regexp"
	"sync"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
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

const projectDirName = "wallet_tt"

func Get() *Config {
	once.Do(func() {
		absPath, err := filepath.Abs(filepath.Dir("."))
		if err != nil {
			panic(err)
		}

		projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)

		rootPath := projectName.Find([]byte(absPath))

		envPath := filepath.Join(string(rootPath), "config", "config.env")

		err = godotenv.Load(envPath)
		if err != nil {
			panic(err)
		}

		err = env.Parse(&config)
		if err != nil {
			panic(err)
		}
	})

	return &config
}
