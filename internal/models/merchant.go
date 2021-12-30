package models

import (
	"time"
)

type Merchant struct {
	ID          string    `json:"id"`
	AdminUserId string    `json:"adminUserId"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"createdAt"`
}
