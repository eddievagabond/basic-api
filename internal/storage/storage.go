package storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

var (
	maxIdleConnections    = 5
	maxOpenConnections    = 10
	connectionTimeout     = time.Second * 5
	connectionMaxLifetime = time.Second * 30
)

type Queries struct {
	ProductRepository  *ProductRepository
	TransferRepository *TransferRepository
	EntryRepository    *EntryRepository
	AccountRepository  *AccountRepository
}

type Storage struct {
	db *sql.DB
	*Queries
}

type Scanner interface {
	Scan(dest ...interface{}) error
}

func NewStorage() (*Storage, error) {
	c := NewConfig()
	databaseURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", c.host, c.port, c.user, c.password, c.dbName, c.SSLMode)

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %s", err)
	}

	db.SetMaxIdleConns(maxIdleConnections)
	db.SetMaxOpenConns(maxOpenConnections)
	db.SetConnMaxIdleTime(connectionTimeout)
	db.SetConnMaxLifetime(connectionMaxLifetime)

	s := &Storage{
		db: db,
	}

	q := &Queries{
		ProductRepository:  NewProductRepository(s),
		TransferRepository: NewTransferRepository(s),
		EntryRepository:    NewEntryRepository(s),
		AccountRepository:  NewAccountRepository(s),
	}

	s.Queries = q

	return s, nil
}

func (s *Storage) executeTx(ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error starting transaction: %s", err)
	}

	err = fn(tx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("error rolling back transaction: %s", rbErr)
		}
		return fmt.Errorf("error executing transaction: %s", err)
	}

	return tx.Commit()
}
