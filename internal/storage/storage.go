package storage

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Storage struct {
	db                *sql.DB
	ProductRepository *ProductRepository
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

	// Maximum Idle Connections
	db.SetMaxIdleConns(5)
	// Maximum Open Connections
	db.SetMaxOpenConns(10)
	// Idle Connection Timeout
	db.SetConnMaxIdleTime(1 * time.Second)
	// Connection Lifetime
	db.SetConnMaxLifetime(30 * time.Second)

	return &Storage{
		db:                db,
		ProductRepository: NewProductRepository(db),
	}, nil
}

// func (s *Storage) executeTransaction(ctx context.Context, fn func(*sql.Tx) error) error {
// 	tx, err := s.conn.BeginTx(ctx, nil)
// 	if err != nil {
// 		return fmt.Errorf("error starting transaction: %s", err)
// 	}

// 	err = fn(tx)
// 	if err != nil {
// 		if rbErr := tx.Rollback(); rbErr != nil {
// 			return fmt.Errorf("error rolling back transaction: %s", rbErr)
// 		}
// 		return fmt.Errorf("error executing transaction: %s", err)
// 	}

// 	return tx.Commit()
// }
