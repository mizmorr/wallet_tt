package integration

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/mizmorr/wallet/internal/model"
	reporitory "github.com/mizmorr/wallet/internal/repository"
	"github.com/mizmorr/wallet/internal/service"
	"github.com/mizmorr/wallet/store"
	"github.com/mizmorr/wallet/store/pg"
	"github.com/stretchr/testify/assert"
)

func TestNewWalletService(t *testing.T) {
	// Setup
	ctx := context.Background()
	store, err := store.New(context.Background())

	// VerifySetup
	assert.NoError(t, err)
	assert.NotNil(t, store)

	// Test
	service, err := service.NewWalletService(store, ctx)

	// Verify
	assert.NoError(t, err)
	assert.NotNil(t, service)
}

func TestNewWalletServiceFailed(t *testing.T) {
	// Setup
	ctx := context.Background()

	// Test
	_, err := service.NewWalletService(nil, ctx)

	// Verify
	assert.Error(t, err)
}

func TestServiceGet(t *testing.T) {
	// Setup

	ctx := context.Background()

	db, err := pg.Dial(ctx)
	assert.NoError(t, err)

	repo := reporitory.NewWalletRepository(db)
	assert.NotNil(t, repo)

	id := uuid.New()
	wallet := &model.Wallet{
		ID:        id,
		Operation: "Deposit",
		Amount:    100,
	}
	_, err = repo.Upsert(ctx, wallet)
	assert.NoError(t, err)

	store, err := store.New(ctx)
	assert.NotNil(t, store)
	assert.NoError(t, err)

	service, err := service.NewWalletService(store, ctx)
	assert.NoError(t, err)

	// Test
	walletNew, err := service.Get(ctx, id)

	// Verify
	assert.NoError(t, err)
	assert.Equal(t, wallet, walletNew)

	_, err = db.Exec(context.Background(), "DELETE FROM wallets WHERE id=$1", id)
	assert.NoError(t, err)
}

func TestServiceCreate(t *testing.T) {
	// Setup
	ctx := context.Background()

	db, err := pg.Dial(ctx)
	assert.NoError(t, err)

	repo := reporitory.NewWalletRepository(db)
	assert.NotNil(t, repo)

	id := uuid.New()
	wallet := &model.Wallet{
		ID:        id,
		Operation: "Deposit",
		Amount:    100,
	}

	store, err := store.New(ctx)
	assert.NotNil(t, store)
	assert.NoError(t, err)

	service, err := service.NewWalletService(store, ctx)
	assert.NoError(t, err)

	// Test
	walletID, err := service.Upsert(ctx, wallet)

	// Verify
	assert.NoError(t, err)
	assert.Equal(t, wallet.ID, walletID)

	walletNew, err := repo.GetByID(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, wallet, walletNew)

	_, err = db.Exec(context.Background(), "DELETE FROM wallets WHERE id=$1", id)
	assert.NoError(t, err)
}

func TestServiceUpdate(t *testing.T) {
	// Setup
	ctx := context.Background()

	db, err := pg.Dial(ctx)
	assert.NoError(t, err)

	repo := reporitory.NewWalletRepository(db)
	assert.NotNil(t, repo)

	id := uuid.New()
	wallet := &model.Wallet{
		ID:        id,
		Operation: "Deposit",
		Amount:    100,
	}
	_, err = repo.Upsert(ctx, wallet)
	assert.NoError(t, err)

	store, err := store.New(ctx)
	assert.NotNil(t, store)
	assert.NoError(t, err)

	service, err := service.NewWalletService(store, ctx)
	assert.NoError(t, err)

	wallet.Operation = "Withdraw"
	wallet.Amount = 50

	// Test
	idUpsert, err := service.Upsert(ctx, wallet)

	// Verify
	assert.NoError(t, err)
	assert.Equal(t, wallet.ID, idUpsert)

	walletNew, err := repo.GetByID(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, wallet, walletNew)
	assert.Equal(t, "Withdraw", walletNew.Operation)
	assert.Equal(t, int64(50), walletNew.Amount)

	_, err = db.Exec(context.Background(), "DELETE FROM wallets WHERE id=$1", id)
	assert.NoError(t, err)
}
