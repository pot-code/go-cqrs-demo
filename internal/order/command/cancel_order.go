package command

import (
	"context"

	"go.uber.org/zap"
)

type CancelOrderCommand struct {
	ID string
}

func NewCancelOrderCommand(id string) *CancelOrderCommand {
	return &CancelOrderCommand{id}
}

type CancelOrderHandler struct {
	logger *zap.Logger
}

func NewCancelOrderHandler(logger *zap.Logger) *CancelOrderHandler {
	return &CancelOrderHandler{logger}
}

func (h *CancelOrderHandler) Handle(ctx context.Context, command *CancelOrderCommand) error {
	return nil
}
