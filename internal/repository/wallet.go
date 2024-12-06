package reporitory

import "github.com/mizmorr/wallet/store/pg"

type WalletRepository struct {
	db *pg.DB
}

func NewWalletRepository(db *pg.DB) *WalletRepository {
	return &WalletRepository{db: db}
}
