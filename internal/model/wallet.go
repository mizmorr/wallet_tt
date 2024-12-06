package model

import "github.com/google/uuid"

type Wallet struct {
	ID        uuid.UUID `pg:"id, notnull, pk"`
	Amount    int64     `pg:"amount, notnull"`
	Operation string    `pg:"operation, notnull"`
}

type WalletRequest struct {
	ID uuid.UUID `json:"id"`
}
