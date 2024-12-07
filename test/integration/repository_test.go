package integration

import (
	"context"
	"testing"

	"github.com/google/uuid"
	logger "github.com/mizmorr/loggerm"
	"github.com/mizmorr/wallet/internal/model"
	reporitory "github.com/mizmorr/wallet/internal/repository"
	"github.com/mizmorr/wallet/pkg/types"
	"github.com/mizmorr/wallet/store/pg"
	"github.com/stretchr/testify/assert"
)

const loggerKey types.ContextKey = "logger"

func TestNewWalletRepository(t *testing.T) {
	// Setup
	ctx := context.WithValue(context.Background(), loggerKey, logger.Get("debug"))
	db, err := pg.Dial(ctx)
	assert.NoError(t, err)
	// Test
	repo := reporitory.NewWalletRepository(db)
	// Verify
	assert.NotNil(t, repo)
}

func TestDeposit(t *testing.T) {
	// Setup
	var expectedAmount int64 = 200
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

	// Test
	err = repo.Deposit(ctx, wallet)
	assert.NoError(t, err)

	// Verify
	walletNew, err := repo.GetByID(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, expectedAmount, walletNew.Amount)

	// Clean up
	_, err = db.Exec(context.Background(), "DELETE FROM wallets WHERE id=$1", id)
	assert.NoError(t, err)
}

func TestWithdraw(t *testing.T) {
	// Setup
	var expectedAmount int64 = 50
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
	walletForWithdraw := &model.Wallet{
		ID:     id,
		Amount: 50,
	}
	_, err = repo.Create(ctx, wallet)
	assert.NoError(t, err)

	// Test
	err = repo.Withdraw(ctx, walletForWithdraw)
	assert.NoError(t, err)

	// Verify
	walletNew, err := repo.GetByID(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, expectedAmount, walletNew.Amount)

	// Clean up
	_, err = db.Exec(context.Background(), "DELETE FROM wallets WHERE id=$1", id)
	assert.NoError(t, err)
}

func TestGetByID(t *testing.T) {
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
	returnedId, err := repo.Create(ctx, wallet)
	assert.NoError(t, err)
	assert.NotNil(t, id)
	assert.Equal(t, id, returnedId)

	// Test
	walletNew, err := repo.GetByID(ctx, id)

	// Verify
	assert.NoError(t, err)
	assert.Equal(t, wallet, walletNew)

	// Clean up
	_, err = db.Exec(context.Background(), "DELETE FROM wallets WHERE id=$1", id)
	assert.NoError(t, err)
}

func TestWithdrawFailed(t *testing.T) {
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
	walletForWithdraw := &model.Wallet{
		ID:     id,
		Amount: 150,
	}
	_, err = repo.Create(ctx, wallet)
	assert.NoError(t, err)

	// Test
	err = repo.Withdraw(ctx, walletForWithdraw)
	assert.Error(t, err)

	// Clean up
	_, err = db.Exec(context.Background(), "DELETE FROM wallets WHERE id=$1", id)
	assert.NoError(t, err)
}
