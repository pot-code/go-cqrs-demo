package repo

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/pot-code/go-cqrs-demo/internal/order/domain"
)

type OrderRedisRepository struct {
	rc *redis.Client
}

var _ domain.OrderRepository = &OrderRedisRepository{}

func NewOrderRedisRepository(rc *redis.Client) *OrderRedisRepository {
	if rc == nil {
		panic("rc is nil")
	}
	return &OrderRedisRepository{rc: rc}
}

func (r *OrderRedisRepository) Save(ctx context.Context, order *domain.Order) error {
	ud, _ := json.Marshal(order)
	err := r.rc.HSet(ctx, "order", order.ID, ud).Err()
	return errors.Wrap(err, "failed to save order")
}

func (r *OrderRedisRepository) GetById(ctx context.Context, id string) (*domain.Order, error) {
	rs := r.rc.HGet(ctx, "order", id)
	if errors.Is(rs.Err(), redis.Nil) {
		return nil, nil
	}

	bs, err := rs.Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get order")
	}

	order := new(domain.Order)
	err = json.Unmarshal(bs, order)
	return order, errors.Wrap(err, "failed to parse order")
}
