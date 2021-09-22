package event

import "time"

type OrderConfirmedEvent struct {
	ID         string    `json:"id,omitempty"`
	Note       string    `json:"note,omitempty"`
	CustomerID string    `json:"customer_id,omitempty"`
	SellerID   string    `json:"seller_id,omitempty"`
	Status     string    `json:"status,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

var _ Event = &OrderConfirmedEvent{}

func (e *OrderConfirmedEvent) Name() string {
	return "OrderConfirmedEvent"
}
