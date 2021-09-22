package port

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/pot-code/go-cqrs-demo/internal/order/command"
	"github.com/pot-code/go-cqrs-demo/internal/order/domain"
	"github.com/pot-code/go-cqrs-demo/internal/order/dto"
	"github.com/pot-code/go-cqrs-demo/internal/order/query"
	"github.com/pot-code/gobit/pkg/api"
	"github.com/pot-code/gobit/pkg/uuid"
	"github.com/pot-code/gobit/pkg/validate"
)

type OrderHandler struct {
	oc  *command.OrderCommands
	oq  *query.OrderQueries
	gid uuid.UUID
	v   *validate.ValidatorV10
}

func NewOrderHandler(oc *command.OrderCommands, oq *query.OrderQueries, gid uuid.UUID, v *validate.ValidatorV10) *OrderHandler {
	return &OrderHandler{oc, oq, gid, v}
}

func (h *OrderHandler) HandleQueryById(c echo.Context) error {
	return nil
}

func (h *OrderHandler) HandleDeleteOrder(c echo.Context) error {
	return nil
}

func (h *OrderHandler) HandleConfirmOrder(c echo.Context) error {
	req := new(dto.ConfirmOrderReq)
	if err := c.Bind(req); err != nil {
		return api.BindErrorResponse(c, err)
	}

	err := h.oc.ConfirmOrderHandler.Handle(c.Request().Context(), command.NewConfirmOrderCommand(req.ID))
	if err != nil {
		if errors.Is(err, domain.ErrOrderNotFound) || errors.Is(err, domain.ErrUnableToConfirm) {
			return api.BadRequestResponse(c, err)
		}
	}

	return err
}

func (h *OrderHandler) HandleCreateOrder(c echo.Context) error {
	req := new(dto.CreateOrderDto)
	if err := c.Bind(req); err != nil {
		return api.BindErrorResponse(c, err)
	}

	if err := h.v.Struct(req); err != nil {
		return api.ValidateFailedResponse(c, h.v.TranslateWithHttpRequest(err, c.Request()))
	}

	guid, _ := h.gid.V4()
	req.ID = guid
	err := h.oc.CreateOrderHandler.Handle(c.Request().Context(), command.NewCreateOrderCommand(req))
	if err != nil {
		return err
	}

	return api.JsonResponse(c, &dto.CreateOrderResponse{ID: guid})
}

func (h *OrderHandler) HandleCancelOrder(c echo.Context) error {
	return nil
}
