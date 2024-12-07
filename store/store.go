package store

import (
	"context"
	"time"

	"github.com/google/uuid"
	logger "github.com/mizmorr/loggerm"
	"github.com/mizmorr/wallet/internal/model"
	reporitory "github.com/mizmorr/wallet/internal/repository"
	"github.com/mizmorr/wallet/store/pg"
	"github.com/pkg/errors"
)

type WalletRepo interface {
	Upsert(ctx context.Context, wallet *model.Wallet) (uuid.UUID, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Wallet, error)
}

type Store struct {
	PG     *pg.DB
	Wallet WalletRepo
}

var _ WalletRepo = (*reporitory.WalletRepository)(nil)

var store Store

func New(ctx context.Context) (*Store, error) {
	logger := logger.GetLoggerFromContext(ctx)

	logger.Debug().Msg("Initializing PostgreSQL store")
	pg, err := pg.Dial(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "pg.Store: failed to connect to the database")
	}

	logger.Debug().Msg("Running PostgreSQL migrations")

	if pg != nil {
		store.PG = pg
		go store.keepAlive(ctx)
		store.Wallet = reporitory.NewWalletRepository(pg)
	}
	logger.Info().Msg("PostgreSQL store initialized successfully")
	return &store, nil
}

const KeepALiveTimeout = 5

func (store *Store) keepAlive(ctx context.Context) {
	logger := logger.GetLoggerFromContext(ctx)
	for {
		time.Sleep(time.Second * KeepALiveTimeout)
		var (
			lost_connection bool
			err             error
		)

		if store.PG == nil {
			lost_connection = true
		}
		if lost_connection {
			logger.Debug().Msg("[store.keepAlive] Lost connection, is trying to reconnect...")
			store.PG, err = pg.Dial(ctx)
			if err != nil {
				logger.Err(err)
			} else {
				logger.Debug().Msg("[store.keepAlive] Connection established")
			}
		}

	}
}
