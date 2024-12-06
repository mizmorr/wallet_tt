package unit

import (
	"context"
	"testing"

	"github.com/mizmorr/wallet/store"
	"github.com/stretchr/testify/assert"
)

func TestNewStore(t *testing.T) {
	// Setup
	expected := &store.Store{}
	// Test
	actual, err := store.New(context.Background())

	// Verify
	assert.NotEqual(t, expected, actual)
	assert.Nil(t, err)
}
