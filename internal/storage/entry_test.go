package storage

import (
	"context"
	"testing"
	"time"

	"github.com/eddievagabond/internal/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var e = &models.Entry{
	ID:        uuid.New().String(),
	AccountID: uuid.New().String(),
	Amount:    100.50,
	CreatedAt: time.Now(),
}

func TestEntryCreate(t *testing.T) {
	db, mock := newMockDB()
	s := &Storage{
		db: db,
	}
	entryRepo := NewEntryRepository(s)

	query := "INSERT INTO entries\\(account_id, amount\\) VALUES\\(\\$1, \\$2\\) RETURNING id, created_at"
	rows := sqlmock.NewRows([]string{"id", "created_at"}).AddRow(e.ID, e.CreatedAt)
	mock.ExpectQuery(query).WithArgs(e.AccountID, e.Amount).WillReturnRows(rows)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(time.Millisecond * 10)
		cancel()
	}()
	result, err := entryRepo.Create(ctx, e)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.Equal(t, e, result)
}
