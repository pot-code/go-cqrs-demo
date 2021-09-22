package repo

import (
	"context"

	"github.com/pkg/errors"
	"github.com/pot-code/go-cqrs-demo/ent"
	"github.com/pot-code/go-cqrs-demo/internal/order/domain"
)

type OrderPostgresRepository struct {
	ec *ent.Client
}

var _ domain.OrderRepository = &OrderPostgresRepository{}

func NewOrderPostgresRepository(ec *ent.Client) *OrderPostgresRepository {
	if ec == nil {
		panic("ec is nil")
	}
	return &OrderPostgresRepository{ec: ec}
}

func (r *OrderPostgresRepository) Save(ctx context.Context, order *domain.Order) error {
	_, err := r.ec.Order.Create().
		SetID(order.ID).
		SetNote(order.Note).
		SetSellerID(order.SellerID).
		SetCustomerID(order.CustomerID).
		SetStatus(order.Status).
		SetCreatedAt(order.CreatedAt).
		Save(ctx)
	return errors.Wrap(err, "failed to save order")
}

func (r *OrderPostgresRepository) GetById(ctx context.Context, id string) (*domain.Order, error) {
	order, err := r.ec.Order.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to get order")
	}

	return &domain.Order{
		ID:         order.ID,
		Note:       order.Note,
		CustomerID: order.CustomerID,
		SellerID:   order.SellerID,
		Status:     order.Status,
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
	}, errors.Wrap(err, "failed to parse order")
}
