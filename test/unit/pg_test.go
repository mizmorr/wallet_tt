package unit

import (
	"context"
	"testing"

	"github.com/mizmorr/wallet/store/pg"
	"github.com/stretchr/testify/assert"
)

func TestDial(t *testing.T) {
	// Setup
	ctx := context.Background()
	expected := &pg.DB{}
	// Test
	actual, err := pg.Dial(ctx)

	// Verify
	assert.Nil(t, err)
	assert.NotEqual(t, expected, actual)

	err = actual.Ping(ctx)
	assert.Nil(t, err)
}
