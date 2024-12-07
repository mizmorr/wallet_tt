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

type WalletWeb struct {
	ID        uuid.UUID `json:"id"`
	Amount    int64     `json:"amount"`
	Operation string    `json:"operation"`
}

func (w *Wallet) ToWeb() *WalletWeb {
	return &WalletWeb{
		ID:        w.ID,
		Amount:    w.Amount,
		Operation: w.Operation,
	}
}
