package storage

import (
	"context"
	"fmt"

	"github.com/eddievagabond/internal/models"
)

type EntryRepository struct {
	storage *Storage
}

func NewEntryRepository(s *Storage) *EntryRepository {
	return &EntryRepository{
		storage: s,
	}
}

func (r *EntryRepository) Create(ctx context.Context, e *models.Entry) (*models.Entry, error) {
	err := r.storage.db.QueryRowContext(ctx, "INSERT INTO entries(account_id, amount) VALUES($1, $2) RETURNING id", e.AccountID, e.Amount).Scan(&e.ID)
	if err != nil {
		return nil, fmt.Errorf("error creating entry: %s", err)
	}
	return e, nil
}
