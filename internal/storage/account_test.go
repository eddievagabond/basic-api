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

var a = &models.Account{
	ID:        uuid.New().String(),
	Owner:     "Owner 1",
	Balance:   10000.50,
	Currency:  "USD",
	CreatedAt: time.Now(),
}

func TestAccountGet(t *testing.T) {
	db, mock := newMockDB()
	s := &Storage{
		db: db,
	}
	accountRepo := NewAccountRepository(s)

	query := "SELECT id, owner, balance, currency, created_at FROM accounts OFFSET \\$1  LIMIT \\$2"
	rows := sqlmock.NewRows([]string{"id", "owner", "balance", "currency", "created_at"}).
		AddRow(a.ID, a.Owner, a.Balance, a.Currency, a.CreatedAt)

	mock.ExpectQuery(query).WithArgs(0, 1).WillReturnRows(rows)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(time.Millisecond * 10)
		cancel()
	}()

	a, err := accountRepo.Get(ctx, 0, 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(a))
}

func TestAccountGetById(t *testing.T) {
	db, mock := newMockDB()
	s := &Storage{
		db: db,
	}
	accountRepo := NewAccountRepository(s)

	query := "SELECT id, owner, balance, currency, created_at FROM accounts WHERE id = \\$1"
	rows := sqlmock.NewRows([]string{"id", "owner", "balance", "currency", "created_at"}).
		AddRow(a.ID, a.Owner, a.Balance, a.Currency, a.CreatedAt)

	mock.ExpectQuery(query).WithArgs(a.ID).WillReturnRows(rows)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(time.Millisecond * 10)
		cancel()
	}()

	result, err := accountRepo.GetById(ctx, a.ID)
	assert.NoError(t, err)
	assert.Equal(t, result, a)
}

func TestAccountCreate(t *testing.T) {
	db, mock := newMockDB()
	s := &Storage{
		db: db,
	}
	accountRepo := NewAccountRepository(s)

	query := "INSERT INTO accounts\\(owner, balance, currency\\) VALUES\\(\\$1, \\$2, \\$3\\) RETURNING id, created_at"
	rows := sqlmock.NewRows([]string{"id", "created_at"}).AddRow(a.ID, a.CreatedAt)
	mock.ExpectQuery(query).WithArgs(a.Owner, a.Balance, a.Currency).WillReturnRows(rows)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(time.Millisecond * 10)
		cancel()
	}()

	result, err := accountRepo.Create(ctx, a)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.Equal(t, result, a)
}

func TestAccountUpdate(t *testing.T) {
	db, mock := newMockDB()
	s := &Storage{
		db: db,
	}
	accountRepo := NewAccountRepository(s)

	query := "UPDATE accounts SET owner = \\$1, balance = \\$2, currency = \\$3 WHERE id = \\$4"
	mock.ExpectExec(query).WithArgs(a.Owner, a.Balance, a.Currency, a.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(time.Millisecond * 10)
		cancel()
	}()

	result, err := accountRepo.Update(ctx, a)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.Equal(t, result, a)
}

func TestAccountDelete(t *testing.T) {
	db, mock := newMockDB()
	s := &Storage{
		db: db,
	}
	accountRepo := NewAccountRepository(s)

	query := "DELETE FROM accounts WHERE id = \\$1"

	mock.ExpectExec(query).WithArgs(p.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(time.Millisecond * 10)
		cancel()
	}()

	err := accountRepo.Delete(ctx, p.ID)
	assert.NoError(t, err)
}
