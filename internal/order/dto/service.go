package dto

type CreateOrderDto struct {
	ID         string `json:"id,omitempty"`
	Note       string `validate:"omitempty,min=1,max=200" json:"note,omitempty"`
	CustomerID string `validate:"required" json:"customer_id,omitempty"`
	SellerID   string `validate:"required" json:"seller_id,omitempty"`
}
