package reporitory

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/mizmorr/wallet/internal/model"
	"github.com/mizmorr/wallet/store/pg"
)

type WalletRepository struct {
	db *pg.DB
}

func NewWalletRepository(db *pg.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (repo *WalletRepository) Upsert(ctx context.Context, wallet *model.Wallet) (uuid.UUID, error) {
	query := `INSERT INTO wallets (id, operation, amount)
		VALUES ($1, $2, $3)
		ON CONFLICT (id)
		DO UPDATE SET operation = COALESCE(NULLIF(EXCLUDED.operation,''), wallets.operation), amount = EXCLUDED.amount;
	`
	_, err := repo.db.Exec(context.Background(), query, wallet.ID, wallet.Operation, wallet.Amount)
	if err != nil {
		return uuid.Nil, err
	}

	return wallet.ID, nil
}

func (repo *WalletRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Wallet, error) {
	var (
		query  = `SELECT id,operation,amount FROM wallets WHERE id = $1`
		wallet model.Wallet
	)

	err := repo.db.QueryRow(context.Background(), query, id).Scan(&wallet.ID, &wallet.Operation, &wallet.Amount)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &wallet, nil
}
