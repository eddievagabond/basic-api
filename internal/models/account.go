package models

import (
	"context"
	"time"
)

type AccountRepository interface {
	Get(ctx context.Context, start, count int) ([]*Account, error)
	GetById(ctx context.Context, id string) (*Account, error)
	Create(ctx context.Context, a *Account) (*Account, error)
	Update(ctx context.Context, a *Account) (*Account, error)
	Delete(ctx context.Context, id string) error
}

type Account struct {
	ID        string    `json:"id"`
	UserId    string    `json:"userId"`
	Balance   float64   `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"createdAt"`
}
