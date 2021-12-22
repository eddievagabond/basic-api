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

var p = &models.Product{
	ID:    uuid.New().String(),
	Name:  "Product 1",
	Price: 100.50,
}

func TestProductGet(t *testing.T) {
	db, mock := newMockDB()
	s := &Storage{
		db: db,
	}
	productRepo := NewProductRepository(s)

	query := "SELECT id, name, price FROM products OFFSET \\$1  LIMIT \\$2"
	rows := sqlmock.NewRows([]string{"id", "name", "price"}).
		AddRow("1", "Product 1", 1.0).
		AddRow("2", "Product 2", 2.0)

	mock.ExpectQuery(query).WithArgs(0, 2).WillReturnRows(rows)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(time.Millisecond * 10)
		cancel()
	}()

	p, err := productRepo.Get(ctx, 0, 2)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(p))
}

func TestProductGetById(t *testing.T) {
	db, mock := newMockDB()
	s := &Storage{
		db: db,
	}
	productRepo := NewProductRepository(s)

	query := "SELECT id, name, price FROM products WHERE id = \\$1"
	rows := sqlmock.NewRows([]string{"id", "name", "price"}).
		AddRow(p.ID, p.Name, p.Price)

	mock.ExpectQuery(query).WithArgs(p.ID).WillReturnRows(rows)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(time.Millisecond * 10)
		cancel()
	}()

	result, err := productRepo.GetById(ctx, p.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.Equal(t, result, p)
}

func TestProductCreate(t *testing.T) {
	db, mock := newMockDB()
	s := &Storage{
		db: db,
	}
	productRepo := NewProductRepository(s)

	query := "INSERT INTO products\\(name, price\\) VALUES\\(\\$1, \\$2\\) RETURNING id"
	rows := sqlmock.NewRows([]string{"id"}).AddRow(p.ID)

	mock.ExpectQuery(query).WithArgs(p.Name, p.Price).WillReturnRows(rows)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(time.Millisecond * 10)
		cancel()
	}()

	result, err := productRepo.Create(ctx, p)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.Equal(t, result, p)
}

func TestProductUpdate(t *testing.T) {
	db, mock := newMockDB()
	s := &Storage{
		db: db,
	}
	productRepo := NewProductRepository(s)

	query := "UPDATE products SET name = \\$1, price = \\$2 WHERE id = \\$3"
	mock.ExpectExec(query).WithArgs(p.Name, p.Price, p.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(time.Millisecond * 10)
		cancel()
	}()

	result, err := productRepo.Update(ctx, p)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.Equal(t, result, p)
}

func TestProductDelete(t *testing.T) {
	db, mock := newMockDB()
	s := &Storage{
		db: db,
	}
	productRepo := NewProductRepository(s)

	query := "DELETE FROM products WHERE id = \\$1"
	mock.ExpectExec(query).WithArgs(p.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(time.Millisecond * 10)
		cancel()
	}()

	err := productRepo.Delete(ctx, p.ID)
	assert.NoError(t, err)
}
