package integration

import (
	"context"
	"testing"

	"github.com/google/uuid"
	logger "github.com/mizmorr/loggerm"
	"github.com/mizmorr/wallet/internal/model"
	reporitory "github.com/mizmorr/wallet/internal/repository"
	"github.com/mizmorr/wallet/internal/service"
	"github.com/mizmorr/wallet/store"
	"github.com/mizmorr/wallet/store/pg"
	"github.com/stretchr/testify/assert"
)

func TestNewWalletService(t *testing.T) {
	// Setup

	ctx := context.WithValue(context.Background(), loggerKey, logger.Get("debug"))

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
	ctx := context.WithValue(context.Background(), loggerKey, logger.Get("debug"))
	// Test
	_, err := service.NewWalletService(nil, ctx)

	// Verify
	assert.Error(t, err)
}

func TestServiceGet(t *testing.T) {
	// Setup
	ctx := context.WithValue(context.Background(), loggerKey, logger.Get("debug"))

	db, err := pg.Dial(ctx)
	assert.NoError(t, err)

	repo := reporitory.NewWalletRepository(db)
	assert.NotNil(t, repo)

	id := uuid.New()
	wallet := &model.Wallet{
		ID:     id,
		Amount: 100,
	}
	_, err = repo.Create(ctx, wallet)
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
	assert.Equal(t, wallet.ToWeb(), walletNew)

	// Clean up
	_, err = db.Exec(context.Background(), "DELETE FROM wallets WHERE id=$1", id)
	assert.NoError(t, err)
}

func TestServiceDeposit(t *testing.T) {
	// Setup
	var (
		ctx                  = context.WithValue(context.Background(), loggerKey, logger.Get("debug"))
		expectedAmount int64 = 150
	)
	db, err := pg.Dial(ctx)
	assert.NoError(t, err)

	repo := reporitory.NewWalletRepository(db)
	assert.NotNil(t, repo)

	id := uuid.New()

	wallet := &model.Wallet{
		ID:     id,
		Amount: 100,
	}

	_, err = repo.Create(ctx, wallet)
	assert.NoError(t, err)

	store, err := store.New(ctx)
	assert.NotNil(t, store)
	assert.NoError(t, err)

	service, err := service.NewWalletService(store, ctx)
	assert.NoError(t, err)
	walletRequest := &model.WalletRequest{
		ID:        id,
		Operation: "Deposit",
		Amount:    50,
	}

	// Test
	err = service.Deposit(ctx, walletRequest)

	// Verify
	assert.NoError(t, err)

	walletNew, err := repo.GetByID(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, wallet.ID, walletNew.ID)
	assert.Equal(t, expectedAmount, walletNew.Amount)

	// Clean up
	_, err = db.Exec(context.Background(), "DELETE FROM wallets WHERE id=$1", id)
	assert.NoError(t, err)
}

func TestServiceWithdraw(t *testing.T) {
	// Setup
	var (
		expectedAmount int64 = 50
		ctx                  = context.WithValue(context.Background(), loggerKey, logger.Get("debug"))
	)

	db, err := pg.Dial(ctx)
	assert.NoError(t, err)

	repo := reporitory.NewWalletRepository(db)
	assert.NotNil(t, repo)

	id := uuid.New()
	wallet := &model.Wallet{
		ID:     id,
		Amount: 100,
	}
	_, err = repo.Create(ctx, wallet)
	assert.NoError(t, err)

	store, err := store.New(ctx)
	assert.NotNil(t, store)
	assert.NoError(t, err)

	service, err := service.NewWalletService(store, ctx)
	assert.NoError(t, err)

	walletRequest := &model.WalletRequest{
		ID:        id,
		Amount:    50,
		Operation: "Withdraw",
	}

	// Test
	err = service.Withdraw(ctx, walletRequest)

	// Verify
	assert.NoError(t, err)

	walletNew, err := repo.GetByID(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, expectedAmount, walletNew.Amount)

	// Clean up
	_, err = db.Exec(context.Background(), "DELETE FROM wallets WHERE id=$1", id)
	assert.NoError(t, err)
}
