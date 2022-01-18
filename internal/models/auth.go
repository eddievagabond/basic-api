package models

import (
	"context"
)

type AuthRepository interface {
	GetByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, u *CreateUserParams, hashedPassword string) (*User, error)
}

type Refresh struct {
	RefreshToken string `json:"refreshToken"`
}
