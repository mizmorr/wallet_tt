package model

import "github.com/google/uuid"

type Wallet struct {
	ID     uuid.UUID `pg:"id, notnull, pk"`
	Amount int64     `pg:"amount, notnull"`
}

type WalletRequest struct {
	ID        uuid.UUID `json:"id" binding:"required"`
	Amount    int64     `json:"amount" validate:"gt=0"`
	Operation string    `json:"operation" validate:"oneof=deposit withdraw"`
}

type WalletResponse struct {
	ID     uuid.UUID `json:"id"`
	Amount int64     `json:"amount"`
}

func (w *Wallet) ToWeb() *WalletResponse {
	return &WalletResponse{
		ID:     w.ID,
		Amount: w.Amount,
	}
}

func (w *WalletRequest) ToDB() *Wallet {
	return &Wallet{
		ID:     w.ID,
		Amount: w.Amount,
	}
}
