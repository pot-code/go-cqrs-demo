// Code generated by entc, DO NOT EDIT.

package order

import (
	"time"
)

const (
	// Label holds the string label denoting the order type in the database.
	Label = "order"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNote holds the string denoting the note field in the database.
	FieldNote = "note"
	// FieldCustomerID holds the string denoting the customer_id field in the database.
	FieldCustomerID = "customer_id"
	// FieldSellerID holds the string denoting the seller_id field in the database.
	FieldSellerID = "seller_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeOrderItems holds the string denoting the order_items edge name in mutations.
	EdgeOrderItems = "order_items"
	// Table holds the table name of the order in the database.
	Table = "orders"
	// OrderItemsTable is the table that holds the order_items relation/edge.
	OrderItemsTable = "order_line_items"
	// OrderItemsInverseTable is the table name for the OrderLineItem entity.
	// It exists in this package in order to avoid circular dependency with the "orderlineitem" package.
	OrderItemsInverseTable = "order_line_items"
	// OrderItemsColumn is the table column denoting the order_items relation/edge.
	OrderItemsColumn = "order_order_items"
)

// Columns holds all SQL columns for order fields.
var Columns = []string{
	FieldID,
	FieldNote,
	FieldCustomerID,
	FieldSellerID,
	FieldStatus,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)