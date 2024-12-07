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

func (svc *WalletService) Get(ctx context.Context, id uuid.UUID) (*model.Wallet, error) {
	wallet, err := svc.store.Wallet.GetByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get wallet by id")
	}
	if wallet == nil {
		return nil, errors.New("wallet not found")
	}
	return wallet, nil
}

func (svc *WalletService) Upsert(ctx context.Context, wallet *model.Wallet) (uuid.UUID, error) {
	id, err := svc.store.Wallet.Upsert(ctx, wallet)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "failed to upsert wallet")
	}
	return id, nil
}
