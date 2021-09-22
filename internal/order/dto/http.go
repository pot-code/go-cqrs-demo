package dto

type ConfirmOrderReq struct {
	ID string `param:"id"`
}

type DeleteOrderReq struct {
	ID string `param:"id"`
}

type CancelOrderReq struct {
	ID string `param:"id"`
}

type QueryOrderByIdReq struct {
	ID string `query:"id"`
}

type CreateOrderResponse struct {
	ID string `json:"id"`
}
