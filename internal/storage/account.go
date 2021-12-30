package storage

import (
	"context"
	"fmt"

	"github.com/eddievagabond/internal/models"
)

type AccountRepository struct {
	storage *Storage
}

func NewAccountRepository(s *Storage) *AccountRepository {
	return &AccountRepository{
		storage: s,
	}
}

func (r *AccountRepository) Get(ctx context.Context, start, count int) ([]*models.Account, error) {
	rows, err := r.storage.db.QueryContext(ctx, "SELECT id, user_id, balance, currency, created_at FROM accounts OFFSET  $1 LIMIT $2", start, count)
	if err != nil {
		return nil, fmt.Errorf("error listing accounts: %s", err)
	}
	defer rows.Close()

	accounts := make([]*models.Account, 0)
	for rows.Next() {
		a := &models.Account{}
		if err := rows.Scan(&a.ID, &a.UserId, &a.Balance, &a.Currency, &a.CreatedAt); err != nil {
			return nil, err
		}

		accounts = append(accounts, a)
	}

	return accounts, nil
}

func (r *AccountRepository) GetById(ctx context.Context, id string) (*models.Account, error) {
	a := &models.Account{}
	if err := r.storage.db.QueryRowContext(ctx, "SELECT id, user_id, balance, currency, created_at FROM accounts WHERE id = $1", id).Scan(&a.ID, &a.UserId, &a.Balance, &a.Currency, &a.CreatedAt); err != nil {
		return nil, fmt.Errorf("error getting account: %s", err)
	}

	return a, nil
}

func (r *AccountRepository) Create(ctx context.Context, a *models.Account) (*models.Account, error) {
	err := r.storage.db.QueryRowContext(ctx, "INSERT INTO accounts(user_id, balance, currency) VALUES($1, $2, $3) RETURNING id, created_at", a.UserId, a.Balance, a.Currency).Scan(&a.ID, &a.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("error creating account: %s", err)
	}

	return a, nil
}

func (r *AccountRepository) Update(ctx context.Context, a *models.Account) (*models.Account, error) {
	_, err := r.storage.db.ExecContext(ctx, "UPDATE accounts SET user_id = $1, balance = $2, currency = $3 WHERE id = $4", a.UserId, a.Balance, a.Currency, a.ID)
	if err != nil {
		return nil, fmt.Errorf("error updating account: %s", err)
	}

	return a, nil
}

func (r *AccountRepository) Delete(ctx context.Context, id string) error {
	_, err := r.storage.db.ExecContext(ctx, "DELETE FROM accounts WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error deleting account: %s", err)
	}

	return nil
}
