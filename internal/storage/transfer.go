package storage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/eddievagabond/internal/models"
)

type TransferRepository struct {
	storage *Storage
}

func NewTransferRepository(s *Storage) *TransferRepository {
	return &TransferRepository{
		storage: s,
	}
}

func (r *TransferRepository) Get(ctx context.Context, start, count int) ([]*models.Transfer, error) {
	rows, err := r.storage.db.QueryContext(ctx, "SELECT id, from_account_id, to_account_id, amount FROM transfers LIMIT $1 OFFSET $2", count, start)
	if err != nil {
		return nil, fmt.Errorf("error listing transfers: %s", err)
	}
	defer rows.Close()

	transfers := make([]*models.Transfer, 0)
	for rows.Next() {
		t := &models.Transfer{}
		if err := rows.Scan(&t.ID, &t.FromAccountID, &t.ToAccountID, &t.Amount); err != nil {
			return nil, err
		}

		transfers = append(transfers, t)
	}

	return transfers, nil
}

func (r *TransferRepository) GetById(ctx context.Context, id string) (*models.Transfer, error) {
	var t models.Transfer

	err := r.storage.db.QueryRowContext(ctx, "SELECT id, from_account_id, to_account_id, amount FROM transfers WHERE id = $1", id).Scan(&t.ID, &t.FromAccountID, &t.ToAccountID, &t.Amount)
	if err != nil {
		return nil, fmt.Errorf("error getting transfer: %s", err)
	}

	return &t, nil
}

func (r *TransferRepository) Create(ctx context.Context, t *models.Transfer) (*models.Transfer, error) {
	err := r.storage.db.QueryRowContext(ctx, "INSERT INTO transfers(from_account_id, to_account_id, amount) VALUES($1, $2, $3) RETURNING id", t.FromAccountID, t.ToAccountID, t.Amount).Scan(&t.ID)
	if err != nil {
		return nil, fmt.Errorf("error creating transfer: %s", err)
	}

	return t, nil
}

func (r *TransferRepository) TransferTx(ctx context.Context, t *models.Transfer) (*models.TransferTxResult, error) {
	var result models.TransferTxResult

	err := r.storage.executeTx(ctx, func(*sql.Tx) error {
		t, err := r.Create(ctx, t)
		if err != nil {
			return err
		}

		fe, err := r.storage.EntryRepository.Create(ctx, &models.Entry{
			AccountID: t.FromAccountID,
			Amount:    -t.Amount,
		})
		if err != nil {
			return err
		}

		te, err := r.storage.EntryRepository.Create(ctx, &models.Entry{
			AccountID: t.ToAccountID,
			Amount:    t.Amount,
		})

		if err != nil {
			return err
		}

		result.Transfer = *t
		result.FromEntry = *fe
		result.ToEntry = *te

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error transfering funds: %s", err)
	}

	return &result, nil
}
