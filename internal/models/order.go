package models

import (
	"fmt"
	"time"
)

type OrderStatus string

const (
	OrderProcessing OrderStatus = "processing"
	OrderShipped    OrderStatus = "shipped"
	OrderDelivered  OrderStatus = "delivered"
	OrderCancelled  OrderStatus = "cancelled"
	OrderReturned   OrderStatus = "resturned"
)

func (s OrderStatus) IsValid() error {
	switch s {
	case OrderProcessing, OrderShipped, OrderDelivered, OrderCancelled, OrderReturned:
		return nil
	default:
		return fmt.Errorf("invalid order status: %s", s)
	}
}

type Order struct {
	ID        string      `json:"id"`
	UserId    string      `json:"userId"`
	Status    OrderStatus `json:"status"`
	CreatedAt time.Time   `json:"createdAt"`
}
