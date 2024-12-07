package integration

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/mizmorr/wallet/internal/model"
	reporitory "github.com/mizmorr/wallet/internal/repository"
	"github.com/mizmorr/wallet/store/pg"
	"github.com/stretchr/testify/assert"
)

func TestNewWalletRepository(t *testing.T) {
	// Setup
	ctx := context.Background()
	db, err := pg.Dial(ctx)
	assert.NoError(t, err)
	// Test
	repo := reporitory.NewWalletRepository(db)
	// Verify
	assert.NotNil(t, repo)
}

func TestUpsert(t *testing.T) {
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
	// Test
	returnedId, err := repo.Upsert(ctx, wallet)
	assert.NoError(t, err)
	assert.NotNil(t, id)
	assert.Equal(t, id, returnedId)
	// Verify
	_, err = db.Exec(context.Background(), "DELETE FROM wallets WHERE id=$1", id)
	assert.NoError(t, err)
}

func TestGetByID(t *testing.T) {
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
	// Test
	returnedId, err := repo.Upsert(ctx, wallet)
	assert.NoError(t, err)
	assert.NotNil(t, id)
	assert.Equal(t, id, returnedId)

	walletNew, err := repo.GetByID(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, wallet, walletNew)

	// Verify
	_, err = db.Exec(context.Background(), "DELETE FROM wallets WHERE id=$1", id)
	assert.NoError(t, err)
}
