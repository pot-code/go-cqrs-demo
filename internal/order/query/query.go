package query

import "github.com/google/wire"

type OrderQueries struct {
	*QueryByIdHandler
}

var QuerySet = wire.NewSet(
	wire.Struct(new(OrderQueries), "*"),
	NewQueryByIdHandler,
)
