package command

import (
	"context"

	"github.com/pot-code/go-cqrs-demo/internal/order/domain"
	"go.uber.org/zap"
)

type ConfirmOrderCommand struct {
	ID string
}

func NewConfirmOrderCommand(id string) *ConfirmOrderCommand {
	return &ConfirmOrderCommand{id}
}

type ConfirmOrderHandler struct {
	logger *zap.Logger
	osvc   *domain.OrderService
}

func NewConfirmOrderHandler(logger *zap.Logger, ag *domain.OrderService) *ConfirmOrderHandler {
	return &ConfirmOrderHandler{logger, ag}
}

func (h *ConfirmOrderHandler) Handle(ctx context.Context, command *ConfirmOrderCommand) error {
	return h.osvc.ConfirmOrder(ctx, command.ID)
}
