package storage

import (
	"context"
	"fmt"

	"github.com/eddievagabond/internal/models"
)

type UserRepository struct {
	storage *Storage
}

func NewUserRepository(s *Storage) *UserRepository {
	return &UserRepository{
		storage: s,
	}
}

func (r *UserRepository) Get(ctx context.Context, start, count int) ([]*models.User, error) {
	rows, err := r.storage.db.QueryContext(ctx, "SELECT id, email, first_name, last_name, created_at FROM users OFFSET  $1 LIMIT $2", start, count)
	if err != nil {
		return nil, fmt.Errorf("error listing users: %s", err)
	}
	defer rows.Close()

	users := make([]*models.User, 0)
	for rows.Next() {
		u := &models.User{}
		if err := rows.Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

func (r *UserRepository) GetById(ctx context.Context, id string) (*models.User, error) {
	u := &models.User{}
	if err := r.storage.db.QueryRowContext(ctx, "SELECT id, email, first_name, last_name, created_at FROM users WHERE id = $1", id).Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.CreatedAt); err != nil {
		return nil, fmt.Errorf("error getting user: %s", err)
	}

	return u, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	u := &models.User{}
	if err := r.storage.db.QueryRowContext(ctx, "SELECT id, email, first_name, last_name, hashed_password, created_at FROM users WHERE email = $1", email).Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.HashedPassword, &u.CreatedAt); err != nil {
		return nil, fmt.Errorf("error getting user: %s", err)
	}

	return u, nil
}

func (r *UserRepository) Create(ctx context.Context, u *models.CreateUserParams, hashedPassword string) (*models.User, error) {
	user := &models.User{
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}

	err := r.storage.db.QueryRowContext(
		ctx,
		"INSERT INTO users(email, first_name, last_name, hashed_password) VALUES($1, $2, $3, $4) RETURNING id, created_at",
		u.Email, u.FirstName, u.LastName, hashedPassword,
	).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		return nil, fmt.Errorf("error creating user: %s", err)
	}

	return user, nil
}

func (r *UserRepository) Update(ctx context.Context, u *models.User) (*models.User, error) {
	_, err := r.storage.db.ExecContext(ctx, "UPDATE users SET email = $1, first_name = $2, last_name = $3 WHERE id = $4", u.Email, u.FirstName, u.LastName, u.ID)
	if err != nil {
		return nil, fmt.Errorf("error updating user: %s", err)
	}

	return u, nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	_, err := r.storage.db.ExecContext(ctx, "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error deleting user: %s", err)
	}

	return nil
}
