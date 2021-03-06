// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/pot-code/go-cqrs-demo/ent/order"
	"github.com/pot-code/go-cqrs-demo/ent/orderlineitem"
	"github.com/pot-code/go-cqrs-demo/ent/predicate"
)

// OrderLineItemQuery is the builder for querying OrderLineItem entities.
type OrderLineItemQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OrderLineItem
	// eager-loading edges.
	withOwner *OrderQuery
	withFKs   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderLineItemQuery builder.
func (oliq *OrderLineItemQuery) Where(ps ...predicate.OrderLineItem) *OrderLineItemQuery {
	oliq.predicates = append(oliq.predicates, ps...)
	return oliq
}

// Limit adds a limit step to the query.
func (oliq *OrderLineItemQuery) Limit(limit int) *OrderLineItemQuery {
	oliq.limit = &limit
	return oliq
}

// Offset adds an offset step to the query.
func (oliq *OrderLineItemQuery) Offset(offset int) *OrderLineItemQuery {
	oliq.offset = &offset
	return oliq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oliq *OrderLineItemQuery) Unique(unique bool) *OrderLineItemQuery {
	oliq.unique = &unique
	return oliq
}

// Order adds an order step to the query.
func (oliq *OrderLineItemQuery) Order(o ...OrderFunc) *OrderLineItemQuery {
	oliq.order = append(oliq.order, o...)
	return oliq
}

// QueryOwner chains the current query on the "owner" edge.
func (oliq *OrderLineItemQuery) QueryOwner() *OrderQuery {
	query := &OrderQuery{config: oliq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oliq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oliq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderlineitem.Table, orderlineitem.FieldID, selector),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orderlineitem.OwnerTable, orderlineitem.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(oliq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrderLineItem entity from the query.
// Returns a *NotFoundError when no OrderLineItem was found.
func (oliq *OrderLineItemQuery) First(ctx context.Context) (*OrderLineItem, error) {
	nodes, err := oliq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orderlineitem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oliq *OrderLineItemQuery) FirstX(ctx context.Context) *OrderLineItem {
	node, err := oliq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrderLineItem ID from the query.
// Returns a *NotFoundError when no OrderLineItem ID was found.
func (oliq *OrderLineItemQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = oliq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orderlineitem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oliq *OrderLineItemQuery) FirstIDX(ctx context.Context) string {
	id, err := oliq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrderLineItem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one OrderLineItem entity is not found.
// Returns a *NotFoundError when no OrderLineItem entities are found.
func (oliq *OrderLineItemQuery) Only(ctx context.Context) (*OrderLineItem, error) {
	nodes, err := oliq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orderlineitem.Label}
	default:
		return nil, &NotSingularError{orderlineitem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oliq *OrderLineItemQuery) OnlyX(ctx context.Context) *OrderLineItem {
	node, err := oliq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrderLineItem ID in the query.
// Returns a *NotSingularError when exactly one OrderLineItem ID is not found.
// Returns a *NotFoundError when no entities are found.
func (oliq *OrderLineItemQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = oliq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orderlineitem.Label}
	default:
		err = &NotSingularError{orderlineitem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oliq *OrderLineItemQuery) OnlyIDX(ctx context.Context) string {
	id, err := oliq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderLineItems.
func (oliq *OrderLineItemQuery) All(ctx context.Context) ([]*OrderLineItem, error) {
	if err := oliq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return oliq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (oliq *OrderLineItemQuery) AllX(ctx context.Context) []*OrderLineItem {
	nodes, err := oliq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrderLineItem IDs.
func (oliq *OrderLineItemQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := oliq.Select(orderlineitem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oliq *OrderLineItemQuery) IDsX(ctx context.Context) []string {
	ids, err := oliq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oliq *OrderLineItemQuery) Count(ctx context.Context) (int, error) {
	if err := oliq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return oliq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (oliq *OrderLineItemQuery) CountX(ctx context.Context) int {
	count, err := oliq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oliq *OrderLineItemQuery) Exist(ctx context.Context) (bool, error) {
	if err := oliq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return oliq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (oliq *OrderLineItemQuery) ExistX(ctx context.Context) bool {
	exist, err := oliq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderLineItemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oliq *OrderLineItemQuery) Clone() *OrderLineItemQuery {
	if oliq == nil {
		return nil
	}
	return &OrderLineItemQuery{
		config:     oliq.config,
		limit:      oliq.limit,
		offset:     oliq.offset,
		order:      append([]OrderFunc{}, oliq.order...),
		predicates: append([]predicate.OrderLineItem{}, oliq.predicates...),
		withOwner:  oliq.withOwner.Clone(),
		// clone intermediate query.
		sql:  oliq.sql.Clone(),
		path: oliq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (oliq *OrderLineItemQuery) WithOwner(opts ...func(*OrderQuery)) *OrderLineItemQuery {
	query := &OrderQuery{config: oliq.config}
	for _, opt := range opts {
		opt(query)
	}
	oliq.withOwner = query
	return oliq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OrderID string `json:"order_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OrderLineItem.Query().
//		GroupBy(orderlineitem.FieldOrderID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (oliq *OrderLineItemQuery) GroupBy(field string, fields ...string) *OrderLineItemGroupBy {
	group := &OrderLineItemGroupBy{config: oliq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := oliq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return oliq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OrderID string `json:"order_id,omitempty"`
//	}
//
//	client.OrderLineItem.Query().
//		Select(orderlineitem.FieldOrderID).
//		Scan(ctx, &v)
//
func (oliq *OrderLineItemQuery) Select(fields ...string) *OrderLineItemSelect {
	oliq.fields = append(oliq.fields, fields...)
	return &OrderLineItemSelect{OrderLineItemQuery: oliq}
}

func (oliq *OrderLineItemQuery) prepareQuery(ctx context.Context) error {
	for _, f := range oliq.fields {
		if !orderlineitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oliq.path != nil {
		prev, err := oliq.path(ctx)
		if err != nil {
			return err
		}
		oliq.sql = prev
	}
	return nil
}

func (oliq *OrderLineItemQuery) sqlAll(ctx context.Context) ([]*OrderLineItem, error) {
	var (
		nodes       = []*OrderLineItem{}
		withFKs     = oliq.withFKs
		_spec       = oliq.querySpec()
		loadedTypes = [1]bool{
			oliq.withOwner != nil,
		}
	)
	if oliq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, orderlineitem.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &OrderLineItem{config: oliq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, oliq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := oliq.withOwner; query != nil {
		ids := make([]string, 0, len(nodes))
		nodeids := make(map[string][]*OrderLineItem)
		for i := range nodes {
			if nodes[i].order_order_items == nil {
				continue
			}
			fk := *nodes[i].order_order_items
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(order.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "order_order_items" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Owner = n
			}
		}
	}

	return nodes, nil
}

func (oliq *OrderLineItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oliq.querySpec()
	return sqlgraph.CountNodes(ctx, oliq.driver, _spec)
}

func (oliq *OrderLineItemQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := oliq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (oliq *OrderLineItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderlineitem.Table,
			Columns: orderlineitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: orderlineitem.FieldID,
			},
		},
		From:   oliq.sql,
		Unique: true,
	}
	if unique := oliq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := oliq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderlineitem.FieldID)
		for i := range fields {
			if fields[i] != orderlineitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oliq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oliq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oliq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oliq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oliq *OrderLineItemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oliq.driver.Dialect())
	t1 := builder.Table(orderlineitem.Table)
	columns := oliq.fields
	if len(columns) == 0 {
		columns = orderlineitem.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oliq.sql != nil {
		selector = oliq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range oliq.predicates {
		p(selector)
	}
	for _, p := range oliq.order {
		p(selector)
	}
	if offset := oliq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oliq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrderLineItemGroupBy is the group-by builder for OrderLineItem entities.
type OrderLineItemGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (oligb *OrderLineItemGroupBy) Aggregate(fns ...AggregateFunc) *OrderLineItemGroupBy {
	oligb.fns = append(oligb.fns, fns...)
	return oligb
}

// Scan applies the group-by query and scans the result into the given value.
func (oligb *OrderLineItemGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := oligb.path(ctx)
	if err != nil {
		return err
	}
	oligb.sql = query
	return oligb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (oligb *OrderLineItemGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := oligb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (oligb *OrderLineItemGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(oligb.fields) > 1 {
		return nil, errors.New("ent: OrderLineItemGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := oligb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (oligb *OrderLineItemGroupBy) StringsX(ctx context.Context) []string {
	v, err := oligb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (oligb *OrderLineItemGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = oligb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderlineitem.Label}
	default:
		err = fmt.Errorf("ent: OrderLineItemGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (oligb *OrderLineItemGroupBy) StringX(ctx context.Context) string {
	v, err := oligb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (oligb *OrderLineItemGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(oligb.fields) > 1 {
		return nil, errors.New("ent: OrderLineItemGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := oligb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (oligb *OrderLineItemGroupBy) IntsX(ctx context.Context) []int {
	v, err := oligb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (oligb *OrderLineItemGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = oligb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderlineitem.Label}
	default:
		err = fmt.Errorf("ent: OrderLineItemGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (oligb *OrderLineItemGroupBy) IntX(ctx context.Context) int {
	v, err := oligb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (oligb *OrderLineItemGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(oligb.fields) > 1 {
		return nil, errors.New("ent: OrderLineItemGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := oligb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (oligb *OrderLineItemGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := oligb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (oligb *OrderLineItemGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = oligb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderlineitem.Label}
	default:
		err = fmt.Errorf("ent: OrderLineItemGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (oligb *OrderLineItemGroupBy) Float64X(ctx context.Context) float64 {
	v, err := oligb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (oligb *OrderLineItemGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(oligb.fields) > 1 {
		return nil, errors.New("ent: OrderLineItemGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := oligb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (oligb *OrderLineItemGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := oligb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (oligb *OrderLineItemGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = oligb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderlineitem.Label}
	default:
		err = fmt.Errorf("ent: OrderLineItemGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (oligb *OrderLineItemGroupBy) BoolX(ctx context.Context) bool {
	v, err := oligb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (oligb *OrderLineItemGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range oligb.fields {
		if !orderlineitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := oligb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oligb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (oligb *OrderLineItemGroupBy) sqlQuery() *sql.Selector {
	selector := oligb.sql.Select()
	aggregation := make([]string, 0, len(oligb.fns))
	for _, fn := range oligb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(oligb.fields)+len(oligb.fns))
		for _, f := range oligb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(oligb.fields...)...)
}

// OrderLineItemSelect is the builder for selecting fields of OrderLineItem entities.
type OrderLineItemSelect struct {
	*OrderLineItemQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (olis *OrderLineItemSelect) Scan(ctx context.Context, v interface{}) error {
	if err := olis.prepareQuery(ctx); err != nil {
		return err
	}
	olis.sql = olis.OrderLineItemQuery.sqlQuery(ctx)
	return olis.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (olis *OrderLineItemSelect) ScanX(ctx context.Context, v interface{}) {
	if err := olis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (olis *OrderLineItemSelect) Strings(ctx context.Context) ([]string, error) {
	if len(olis.fields) > 1 {
		return nil, errors.New("ent: OrderLineItemSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := olis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (olis *OrderLineItemSelect) StringsX(ctx context.Context) []string {
	v, err := olis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (olis *OrderLineItemSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = olis.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderlineitem.Label}
	default:
		err = fmt.Errorf("ent: OrderLineItemSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (olis *OrderLineItemSelect) StringX(ctx context.Context) string {
	v, err := olis.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (olis *OrderLineItemSelect) Ints(ctx context.Context) ([]int, error) {
	if len(olis.fields) > 1 {
		return nil, errors.New("ent: OrderLineItemSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := olis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (olis *OrderLineItemSelect) IntsX(ctx context.Context) []int {
	v, err := olis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (olis *OrderLineItemSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = olis.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderlineitem.Label}
	default:
		err = fmt.Errorf("ent: OrderLineItemSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (olis *OrderLineItemSelect) IntX(ctx context.Context) int {
	v, err := olis.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (olis *OrderLineItemSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(olis.fields) > 1 {
		return nil, errors.New("ent: OrderLineItemSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := olis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (olis *OrderLineItemSelect) Float64sX(ctx context.Context) []float64 {
	v, err := olis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (olis *OrderLineItemSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = olis.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderlineitem.Label}
	default:
		err = fmt.Errorf("ent: OrderLineItemSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (olis *OrderLineItemSelect) Float64X(ctx context.Context) float64 {
	v, err := olis.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (olis *OrderLineItemSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(olis.fields) > 1 {
		return nil, errors.New("ent: OrderLineItemSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := olis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (olis *OrderLineItemSelect) BoolsX(ctx context.Context) []bool {
	v, err := olis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (olis *OrderLineItemSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = olis.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderlineitem.Label}
	default:
		err = fmt.Errorf("ent: OrderLineItemSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (olis *OrderLineItemSelect) BoolX(ctx context.Context) bool {
	v, err := olis.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (olis *OrderLineItemSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := olis.sql.Query()
	if err := olis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
