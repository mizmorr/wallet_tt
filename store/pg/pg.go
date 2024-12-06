package pg

import (
	"context"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mizmorr/auth_tt/pkg/logger"
	"github.com/mizmorr/wallet/config"
)

type DB struct {
	*pgxpool.Pool
}

var (
	once       sync.Once
	pgInstance *DB
)

func Dial(ctx context.Context) (*DB, error) {
	confg := config.Get()
	log := logger.GetLoggerFromContext(ctx)

	log.Debug().Msg("Database url checking...")
	if confg.DatabaseURL == "" {
		return nil, errors.New("No database URL provided")
	}

	log.Debug().Msg("Database config parsing...")
	poolConfig, err := pgxpool.ParseConfig(confg.DatabaseURL)
	if err != nil {
		return nil, errors.Wrap(err, "Parse config failed")
	}
	poolConfig.MaxConnIdleTime = confg.PgMaxIdleTime
	poolConfig.HealthCheckPeriod = confg.PgHealthCheckPeriod
	poolConfig.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	once.Do(func() {
		var db *pgxpool.Pool
		for confg.PgConnAttempts > 0 {

			db, err = pgxpool.NewWithConfig(ctx, poolConfig)
			if err == nil {
				log.Info().Msg("Connect to pg is established")
				break
			}

			log.Error().Err(err).Msg("Failed to connect to pg, retrying...")

			time.Sleep(confg.PgTimeout)

			confg.PgConnAttempts--
		}
		if err != nil {
			panic(errors.Wrap(err, "Cannot connect to PostgreSQL database"))
		}
		pgInstance = &DB{db}
	})

	return pgInstance, nil
}
