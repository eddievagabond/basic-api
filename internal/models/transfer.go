package models

import (
	"context"
	"time"
)

type Transfer struct {
	ID            string    `json:"id"`
	FromAccountID string    `json:"fromAccountId"`
	ToAccountID   string    `json:"toAccountId"`
	Amount        float64   `json:"amount"`
	CreatedAt     time.Time `json:"createdAt"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"fromAccount"`
	ToAccount   Account  `json:"toAccount"`
	FromEntry   Entry    `json:"fromEntry"`
	ToEntry     Entry    `json:"toEntry"`
}

type TransferRepository interface {
	Get(ctx context.Context, start, count int) ([]*Transfer, error)
	GetById(ctx context.Context, id string) (*Transfer, error)
	Create(ctx context.Context, t *Transfer) (*Transfer, error)
	TransferTx(ctx context.Context, t *Transfer) (*TransferTxResult, error)
}
