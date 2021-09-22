package query

import (
	"context"

	"github.com/pot-code/go-cqrs-demo/ent"
	"github.com/pot-code/go-cqrs-demo/internal/order/domain"
	"github.com/pot-code/go-cqrs-demo/internal/order/dto"
)

type QueryById struct {
	dto *dto.QueryOrderByIdReq
}

type QueryByIdHandler struct {
	client *ent.Client
}

func NewQueryByIdHandler(ec *ent.Client) *QueryByIdHandler {
	return &QueryByIdHandler{ec}
}

func (qh *QueryByIdHandler) Handle(ctx context.Context, c *QueryById) (*domain.Order, error) {
	return nil, nil
}
