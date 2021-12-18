package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	conn *sql.DB
}

type Scanner interface {
	Scan(dest ...interface{}) error
}

func NewStorage(databaseURL string) (*Storage, error) {
	conn, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %s", err)
	}

	return &Storage{
		conn: conn,
	}, nil
}
