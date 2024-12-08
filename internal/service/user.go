package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/mizmorr/wallet/internal/model"
	"github.com/mizmorr/wallet/store"
	"github.com/pkg/errors"
)

type WalletService struct {
	store *store.Store
	ctx   context.Context
}

func NewWalletService(store *store.Store, ctx context.Context) (*WalletService, error) {
	if store == nil {
		return nil, errors.New("store is nil")
	}
	return &WalletService{
		store: store,
		ctx:   ctx,
	}, nil
}

func (svc *WalletService) Get(ctx context.Context, id uuid.UUID) (*model.WalletResponse, error) {
	wallet, err := svc.store.Wallet.GetByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get wallet by id")
	}

	return wallet.ToWeb(), nil
}

func (svc *WalletService) Deposit(ctx context.Context, wallet *model.WalletRequest) error {
	walletDB := wallet.ToDB()

	err := svc.store.Wallet.Deposit(ctx, walletDB)
	if err != nil {
		return errors.Wrap(err, "failed to make deposit")
	}
	return nil
}

func (svc *WalletService) Withdraw(ctx context.Context, wallet *model.WalletRequest) error {
	walletDB := wallet.ToDB()

	err := svc.store.Wallet.Withdraw(ctx, walletDB)
	if err != nil {
		return errors.Wrap(err, "failed to make withdraw")
	}
	return nil
}
