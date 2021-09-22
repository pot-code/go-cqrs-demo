package command

import "github.com/google/wire"

type OrderCommands struct {
	*CreateOrderHandler
	*DeleteOrderHandler
	*ConfirmOrderHandler
	*CancelOrderHandler
}

var CommandSet = wire.NewSet(
	wire.Struct(new(OrderCommands), "*"),
	NewCancelOrderHandler,
	NewConfirmOrderHandler,
	NewCreateOrderHandler,
	NewDeleteOrderHandler,
)
