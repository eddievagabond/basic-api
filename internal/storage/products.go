package storage

import (
	"context"
	"fmt"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (s *Storage) CreateProduct(ctx context.Context, req Product) (*Product, error) {
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

func (s *Storage) UpdateProduct(ctx context.Context, req Product) (*Product, error) {
	_, err := s.conn.ExecContext(ctx, "UPDATE products SET name = $1, price = $2 WHERE id = $3", req.Name, req.Price, req.ID)
	if err != nil {
		return nil, fmt.Errorf("error updating product: %s", err)
	}
	return &req, nil
}

func (s *Storage) DeleteProduct(ctx context.Context, id string) error {
	_, err := s.conn.ExecContext(ctx, "DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error deleting product: %s", err)
	}
	return nil
}

func (s *Storage) GetProduct(ctx context.Context, id string) (*Product, error) {
	p := &Product{}
	err := s.conn.QueryRowContext(ctx, "SELECT id, name, price FROM products WHERE id = $1", id).Scan(&p.ID, &p.Name, &p.Price)
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
