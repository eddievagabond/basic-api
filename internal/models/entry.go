package models

import (
	"time"
)

type Entry struct {
	ID        string    `json:"id"`
	AccountID string    `json:"accountId"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}
