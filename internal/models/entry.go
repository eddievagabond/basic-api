package models

import (
	"context"
	"time"
)

type Entry struct {
	ID        string    `json:"id"`
	AccountID string    `json:"accountId"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}

type EntryRepository interface {
	Create(ctx context.Context, e *Entry) (*Entry, error)
}
