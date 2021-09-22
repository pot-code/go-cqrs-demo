package domain

import (
	"time"

	"github.com/pkg/errors"
	"github.com/pot-code/go-cqrs-demo/ent"
)

type Order ent.Order

var (
	ErrUnableToConfirm = errors.New("unable to confirm order")
	ErrUnableToCancel  = errors.New("unable to cancel order")
	ErrUnableToDelete  = errors.New("unable to delete order")
	ErrDuplicatedOrder = errors.New("duplicated order")
	ErrOrderNotFound   = errors.New("order not found")
)

const (
	StatusCreated   = "created"
	StatusCancelled = "cancelled"
	StatusConfirmed = "confirmed"
	StatusDeleted   = "deleted"
)

func (o *Order) Confirm() error {
	if o.Status != StatusCreated {
		return ErrUnableToConfirm
	}

	o.Status = StatusConfirmed
	o.UpdatedAt = time.Now()
	return nil
}

func (o *Order) Cancel() error {
	if o.Status != StatusDeleted {
		return ErrUnableToCancel
	}

	o.Status = StatusCancelled
	o.UpdatedAt = time.Now()
	return nil
}

func (o *Order) Delete() error {
	if o.Status == StatusConfirmed {
		return ErrUnableToDelete
	}

	o.Status = StatusDeleted
	return nil
}
