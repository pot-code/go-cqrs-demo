package command

import (
	"context"

	"github.com/pot-code/go-cqrs-demo/event"
	"go.uber.org/zap"
)

type DeleteOrderCommand struct {
	ID string
}

func NewDeleteOrderCommand(id string) *DeleteOrderCommand {
	return &DeleteOrderCommand{id}
}

type DeleteOrderHandler struct {
	logger *zap.Logger
	pub    event.Publisher
}

func NewDeleteOrderHandler(logger *zap.Logger, pub event.Publisher) *DeleteOrderHandler {
	return &DeleteOrderHandler{logger, pub}
}

func (h *DeleteOrderHandler) Handle(ctx context.Context, command *DeleteOrderCommand) error {
	return nil
}
