package reporitory

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	logger "github.com/mizmorr/loggerm"
	"github.com/mizmorr/wallet/internal/model"
	"github.com/mizmorr/wallet/store/pg"
)

type WalletRepository struct {
	db *pg.DB
}

func NewWalletRepository(db *pg.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (repo *WalletRepository) Withdraw(ctx context.Context, wallet *model.Wallet) error {
	var (
		query = `UPDATE wallets
		SET amount = amount - $2
		WHERE id = $1 AND amount >= $2;`

		logger = logger.GetLoggerFromContext(ctx)
	)

	logger.Debug().Msg("Running withdraw on wallet with id: " + wallet.ID.String())

	cmdTag, err := repo.db.Exec(context.Background(), query, wallet.ID, wallet.Amount)
	if err != nil {
		logger.Debug().Msg("Error occurred when withdrawing: " + err.Error())
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		logger.Debug().Msg("No rows affected while updating")
		return errors.New("Insufficient funds")
	}

	logger.Debug().Msg("Withdraw is successful on wallet with id: " + wallet.ID.String())
	return nil
}

func (repo *WalletRepository) Deposit(ctx context.Context, wallet *model.Wallet) error {
	var (
		query = `UPDATE wallets
		SET amount = amount + $2
		WHERE id = $1;`
		logger = logger.GetLoggerFromContext(ctx)
	)

	logger.Debug().Msg("Making deposit on wallet with id: " + wallet.ID.String())

	cmdTag, err := repo.db.Exec(context.Background(), query, wallet.ID, wallet.Amount)
	if err != nil {
		logger.Debug().Msg("Error occurred when depositing: " + err.Error())
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		logger.Debug().Msg("No such wallet found")
		return errors.New("Wallet with specified id was not found!")
	}
	logger.Debug().Msg("Deposit is successful on wallet with id: " + wallet.ID.String())
	return nil
}

func (repo *WalletRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Wallet, error) {
	var (
		query  = `SELECT id,amount FROM wallets WHERE id = $1`
		wallet model.Wallet
		logger = logger.GetLoggerFromContext(ctx)
	)

	logger.Debug().Msg("Running get on wallet with id: " + id.String())
	err := repo.db.QueryRow(context.Background(), query, id).Scan(&wallet.ID, &wallet.Amount)
	if err != nil {
		if err == pgx.ErrNoRows {
			logger.Debug().Msg("No such wallet in database with id: " + id.String())
			return nil, errors.New("Wallet was not found")
		}

		logger.Debug().Msg("Error occurred while running get on wallet with id: " + id.String())
		return nil, err
	}

	logger.Debug().Msg("Retrieved successfully on wallet with id: " + id.String())
	return &wallet, nil
}

func (repo *WalletRepository) Create(ctx context.Context, wallet *model.Wallet) (uuid.UUID, error) {
	var (
		query  = `INSERT INTO wallets (id, amount) VALUES ($1, $2) RETURNING id;`
		logger = logger.GetLoggerFromContext(ctx)
	)

	logger.Debug().Msg("Creating new wallet with id: " + wallet.ID.String())

	err := repo.db.QueryRow(context.Background(), query, wallet.ID, wallet.Amount).Scan(&wallet.ID)
	if err != nil {
		logger.Debug().Msg("Error occurred while creating wallet: " + err.Error())
		return uuid.Nil, err
	}

	logger.Debug().Msg("Created wallet with id: " + wallet.ID.String())
	return wallet.ID, nil
}
