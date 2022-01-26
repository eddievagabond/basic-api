package models

import (
	"context"
)

type AuthRepository interface {
	GetByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, u *CreateUserParams, hashedPassword string) (*User, error)
}

type RefreshRequestParams struct {
	RefreshToken string `json:"refreshToken"`
}

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
