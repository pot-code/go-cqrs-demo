package event

import (
	"time"
)

type OrderCreateEvent struct {
	ID         string    `json:"id,omitempty"`
	Note       string    `json:"note,omitempty"`
	CustomerID string    `json:"customer_id,omitempty"`
	SellerID   string    `json:"seller_id,omitempty"`
	Status     string    `json:"status,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

var _ Event = &OrderCreateEvent{}

func (e *OrderCreateEvent) Name() string {
	return "OrderCreatedEvent"
}
