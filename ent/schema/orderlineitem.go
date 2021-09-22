package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OrderLineItem holds the schema definition for the OrderLineItem entity.
type OrderLineItem struct {
	ent.Schema
}

// Fields of the OrderLineItem.
func (OrderLineItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable(),
		field.String("order_id"),
		field.String("product_id"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the OrderLineItem.
func (OrderLineItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Order.Type).Ref("order_items").Unique(),
	}
}
