package storage

import (
	"context"
	"fmt"
)

type CreateProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (s *Storage) CreateProduct(ctx context.Context, req CreateProductRequest) (*Product, error) {
	p := &Product{
		Name:  req.Name,
		Price: req.Price,
	}
	err := s.conn.QueryRowContext(ctx, "INSERT INTO products(name, price) VALUES($1, $2) RETURNING id", req.Name, req.Price).Scan(&p.ID)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *Storage) ListProducts(ctx context.Context, start, count int) ([]*Product, error) {
	rows, err := s.conn.QueryContext(ctx, "SELECT id, name, price FROM products LIMIT $1 OFFSET $2", count, start)
	if err != nil {
		return nil, fmt.Errorf("error listing products: %s", err)
	}
	defer rows.Close()

	products := make([]*Product, 0)
	for rows.Next() {
		p := &Product{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}
