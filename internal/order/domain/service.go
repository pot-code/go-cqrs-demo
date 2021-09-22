package domain

import (
	"context"
	"time"

	"github.com/pot-code/go-cqrs-demo/event"
	"github.com/pot-code/go-cqrs-demo/internal/order/config"
	"github.com/pot-code/go-cqrs-demo/internal/order/dto"
)

type OrderService struct {
	pub  event.Publisher
	repo OrderRepository
	et   *config.EventTopicConfig
}

func NewOrderService(pub event.Publisher, repo OrderRepository, et *config.EventTopicConfig) *OrderService {
	return &OrderService{pub: pub, repo: repo, et: et}
}

func (a *OrderService) CreateOrder(ctx context.Context, dto *dto.CreateOrderDto) error {
	id := dto.ID
	o, err := a.repo.GetById(ctx, id)
	if err != nil {
		return err
	}
	if o != nil {
		return ErrDuplicatedOrder
	}

	order := &Order{
		ID:         dto.ID,
		CustomerID: dto.CustomerID,
		SellerID:   dto.SellerID,
		Note:       dto.Note,
		Status:     StatusCreated,
		CreatedAt:  time.Now(),
	}
	if err := a.repo.Save(ctx, order); err != nil {
		return err
	}

	e := &event.OrderCreateEvent{
		ID:         order.ID,
		CustomerID: order.CustomerID,
		SellerID:   order.SellerID,
		Note:       order.Note,
		Status:     order.Status,
		CreatedAt:  order.CreatedAt,
	}
	return a.pub.Publish(ctx, a.et.OrderCreated, dto.ID, e)
}

func (a *OrderService) ConfirmOrder(ctx context.Context, id string) error {
	o, err := a.repo.GetById(ctx, id)
	if err != nil {
		return err
	}
	if o == nil {
		return ErrOrderNotFound
	}

	if err := o.Confirm(); err != nil {
		return err
	}
	if err := a.repo.Save(ctx, o); err != nil {
		return err
	}

	e := &event.OrderConfirmedEvent{
		ID:         o.ID,
		CustomerID: o.CustomerID,
		SellerID:   o.SellerID,
		Note:       o.Note,
		Status:     o.Status,
		UpdatedAt:  o.UpdatedAt,
	}
	return a.pub.Publish(ctx, a.et.OrderConfirmed, o.ID, e)
}
