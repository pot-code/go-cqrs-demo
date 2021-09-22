package domain

import "context"

type OrderRepository interface {
	Save(ctx context.Context, order *Order) error
	GetById(ctx context.Context, id string) (*Order, error)
}
