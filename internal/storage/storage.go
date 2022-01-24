package storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/eddievagabond/internal/util"
	_ "github.com/lib/pq"
)

var (
	maxIdleConnections    = 5
	maxOpenConnections    = 10
	connectionTimeout     = time.Second * 5
	connectionMaxLifetime = time.Second * 30
)

type Queries struct {
	UserRepository     *UserRepository
	ProductRepository  *ProductRepository
	TransferRepository *TransferRepository
	EntryRepository    *EntryRepository
	AccountRepository  *AccountRepository
}

type Storage struct {
	db *sql.DB
	*Queries
}

func NewStorage(config *util.Configuration) (*Storage, error) {
	databaseURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.DBHost, config.DBPort, config.DBUser, config.DBPass, config.DBName, "disable")

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
		UserRepository:     NewUserRepository(s),
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
