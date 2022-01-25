package models

import (
	"context"
	"fmt"
	"time"
)

type UserRepository interface {
	Get(ctx context.Context, start, count int) ([]*User, error)
	GetById(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, u *CreateUserParams, hashedPassword string) (*User, error)
}

type CreateUserParams struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (u *CreateUserParams) Validate() error {
	if u.Email == "" {
		return fmt.Errorf("email is required")
	}

	if u.Password == "" {
		return fmt.Errorf("password is required")
	}

	if u.FirstName == "" {
		return fmt.Errorf("first name is required")
	}

	if u.LastName == "" {
		return fmt.Errorf("last name is required")
	}

	return nil
}

type LoginUserParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID             string    `json:"id"`
	Email          string    `json:"email"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	HashedPassword string    `json:"hashedPassword"`
	CreatedAt      time.Time `json:"createdAt"`
}

func (u *User) Validate() error {
	if u.Email == "" {
		return fmt.Errorf("email is required")
	}

	if u.FirstName == "" {
		return fmt.Errorf("first name is required")
	}

	if u.LastName == "" {
		return fmt.Errorf("last name is required")
	}

	return nil
}

func (u *User) Sanitize() *User {
	return &User{
		ID:             u.ID,
		Email:          u.Email,
		FirstName:      u.FirstName,
		LastName:       u.LastName,
		HashedPassword: "",
		CreatedAt:      u.CreatedAt,
	}
}
