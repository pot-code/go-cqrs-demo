package command

import (
	"context"

	"github.com/pot-code/go-cqrs-demo/internal/order/domain"
	"github.com/pot-code/go-cqrs-demo/internal/order/dto"
	"go.uber.org/zap"
)

type CreateOrderCommand struct {
	dto *dto.CreateOrderDto
}

func NewCreateOrderCommand(dto *dto.CreateOrderDto) *CreateOrderCommand {
	return &CreateOrderCommand{dto}
}

type CreateOrderHandler struct {
	logger *zap.Logger
	osvc   *domain.OrderService
}

func NewCreateOrderHandler(logger *zap.Logger, osvc *domain.OrderService) *CreateOrderHandler {
	return &CreateOrderHandler{logger, osvc}
}

func (h *CreateOrderHandler) Handle(ctx context.Context, command *CreateOrderCommand) error {
	return h.osvc.CreateOrder(ctx, command.dto)
}
