package consumer

import (
	"context"
	"encoding/json"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
	"github.com/pot-code/go-cqrs-demo/ent"
	"github.com/pot-code/go-cqrs-demo/event"
	"github.com/pot-code/go-cqrs-demo/pkg/mq"
	"go.uber.org/zap"
)

type OrderCreateHandler struct {
	logger *zap.Logger
	ec     *ent.Client
}

var _ mq.EventHandler = &OrderCreateHandler{}

func NewOrderCreateHandler(
	c *mq.KafkaEventConsumer,
	logger *zap.Logger,
	ec *ent.Client,
) *OrderCreateHandler {
	self := &OrderCreateHandler{logger, ec}
	c.RegisterHandler(new(event.OrderCreateEvent), self)
	return self
}

func (h *OrderCreateHandler) Handle(ctx context.Context, msg *sarama.ConsumerMessage) error {
	e := new(event.OrderCreateEvent)
	if err := json.Unmarshal(msg.Value, e); err != nil {
		return errors.Wrap(err, "failed to unmarshal OrderCreateEvent")
	}

	_, err := h.ec.Order.Create().
		SetID(e.ID).
		SetCustomerID(e.CustomerID).
		SetSellerID(e.SellerID).
		SetNote(e.Note).
		SetStatus(e.Status).
		SetCreatedAt(e.CreatedAt).
		Save(ctx)

	return errors.Wrap(err, "failed to save Order")
}
