package models

import (
	"context"
	"time"
)

type UserRepository interface {
	Get(ctx context.Context, start, count int) ([]*User, error)
	GetById(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, u *User) (*User, error)
	Update(ctx context.Context, u *User) (*User, error)
	Delete(ctx context.Context, id string) error
}

type User struct {
	ID             string    `json:"id"`
	Email          string    `json:"email"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	HashedPassword string    `json:"hashedPassword"`
	CreatedAt      time.Time `json:"createdAt"`
}
