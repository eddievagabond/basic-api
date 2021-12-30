package models

import (
	"context"
	"time"
)

type UserRepository interface {
	Get(ctx context.Context, start, count int) ([]*UserResponse, error)
	GetById(ctx context.Context, id string) (*UserResponse, error)
	Create(ctx context.Context, u *CreateUserParams) (*UserResponse, error)
	Update(ctx context.Context, u *CreateUserParams) (*UserResponse, error)
	Delete(ctx context.Context, id string) error
}

type CreateUserParams struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	CreatedAt time.Time `json:"createdAt"`
}

type User struct {
	ID             string    `json:"id"`
	Email          string    `json:"email"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	HashedPassword string    `json:"hashedPassword"`
	IsActive       bool      `json:"isActive"`
	InviteToken    string    `json:"inviteToken"`
	InvitedAt      time.Time `json:"invitedAt"`
	CreatedAt      time.Time `json:"createdAt"`
}
