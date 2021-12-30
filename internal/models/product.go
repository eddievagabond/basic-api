package models

import (
	"context"
	"time"
)

type Product struct {
	ID         string    `json:"id"`
	MerchantID string    `json:"merchantId"`
	Name       string    `json:"name"`
	Price      float64   `json:"price"`
	Quantity   int       `json:"quantity"`
	CreatedAt  time.Time `json:"createdAt"`
}

type ProductRepository interface {
	Get(ctx context.Context, start, count int) ([]*Product, error)
	GetById(ctx context.Context, id string) (*Product, error)
	Create(ctx context.Context, p *Product) (*Product, error)
	Update(ctx context.Context, p *Product) (*Product, error)
	Delete(ctx context.Context, id string) error
}
