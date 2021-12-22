package storage

import (
	"context"
	"fmt"

	"github.com/eddievagabond/internal/models"
)

type ProductRepository struct {
	storage *Storage
}

func NewProductRepository(s *Storage) *ProductRepository {
	return &ProductRepository{
		storage: s,
	}
}

func (r *ProductRepository) Get(ctx context.Context, start, count int) ([]*models.Product, error) {
	rows, err := r.storage.db.QueryContext(ctx, "SELECT id, name, price FROM products OFFSET $1 LIMIT $2", start, count)
	if err != nil {
		return nil, fmt.Errorf("error listing products: %s", err)
	}
	defer rows.Close()

	products := make([]*models.Product, 0)
	for rows.Next() {
		p := &models.Product{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func (r *ProductRepository) GetById(ctx context.Context, id string) (*models.Product, error) {
	p := &models.Product{}
	if err := r.storage.db.QueryRowContext(ctx, "SELECT id, name, price FROM products WHERE id = $1", id).Scan(&p.ID, &p.Name, &p.Price); err != nil {
		return nil, fmt.Errorf("error getting product: %s", err)
	}

	return p, nil
}

func (r *ProductRepository) Create(ctx context.Context, p *models.Product) (*models.Product, error) {
	err := r.storage.db.QueryRowContext(ctx, "INSERT INTO products(name, price) VALUES($1, $2) RETURNING id", p.Name, p.Price).Scan(&p.ID)
	if err != nil {
		return nil, fmt.Errorf("error creating product: %s", err)
	}

	return p, nil
}

func (r *ProductRepository) Update(ctx context.Context, p *models.Product) (*models.Product, error) {
	_, err := r.storage.db.ExecContext(ctx, "UPDATE products SET name = $1, price = $2 WHERE id = $3", p.Name, p.Price, p.ID)
	if err != nil {
		return nil, fmt.Errorf("error updating product: %s", err)
	}

	return p, nil
}

func (r *ProductRepository) Delete(ctx context.Context, id string) error {
	_, err := r.storage.db.ExecContext(ctx, "DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error deleting product: %s", err)
	}

	return nil
}
