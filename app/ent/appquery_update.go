// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/sujithshajee/dnsbl/app/ent/appquery"
	"github.com/sujithshajee/dnsbl/app/ent/appresponse"
	"github.com/sujithshajee/dnsbl/app/ent/ip"
	"github.com/sujithshajee/dnsbl/app/ent/predicate"
)

// AppQueryUpdate is the builder for updating AppQuery entities.
type AppQueryUpdate struct {
	config
	hooks    []Hook
	mutation *AppQueryMutation
}

// Where appends a list predicates to the AppQueryUpdate builder.
func (aqu *AppQueryUpdate) Where(ps ...predicate.AppQuery) *AppQueryUpdate {
	aqu.mutation.Where(ps...)
	return aqu
}

// SetCreatedAt sets the "created_at" field.
func (aqu *AppQueryUpdate) SetCreatedAt(t time.Time) *AppQueryUpdate {
	aqu.mutation.SetCreatedAt(t)
	return aqu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aqu *AppQueryUpdate) SetNillableCreatedAt(t *time.Time) *AppQueryUpdate {
	if t != nil {
		aqu.SetCreatedAt(*t)
	}
	return aqu
}

// SetUpdatedAt sets the "updated_at" field.
func (aqu *AppQueryUpdate) SetUpdatedAt(t time.Time) *AppQueryUpdate {
	aqu.mutation.SetUpdatedAt(t)
	return aqu
}

// AddResponseIDs adds the "responses" edge to the AppResponse entity by IDs.
func (aqu *AppQueryUpdate) AddResponseIDs(ids ...uuid.UUID) *AppQueryUpdate {
	aqu.mutation.AddResponseIDs(ids...)
	return aqu
}

// AddResponses adds the "responses" edges to the AppResponse entity.
func (aqu *AppQueryUpdate) AddResponses(a ...*AppResponse) *AppQueryUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aqu.AddResponseIDs(ids...)
}

// SetIpaddressID sets the "ipaddress" edge to the IP entity by ID.
func (aqu *AppQueryUpdate) SetIpaddressID(id uuid.UUID) *AppQueryUpdate {
	aqu.mutation.SetIpaddressID(id)
	return aqu
}

// SetIpaddress sets the "ipaddress" edge to the IP entity.
func (aqu *AppQueryUpdate) SetIpaddress(i *IP) *AppQueryUpdate {
	return aqu.SetIpaddressID(i.ID)
}

// Mutation returns the AppQueryMutation object of the builder.
func (aqu *AppQueryUpdate) Mutation() *AppQueryMutation {
	return aqu.mutation
}

// ClearResponses clears all "responses" edges to the AppResponse entity.
func (aqu *AppQueryUpdate) ClearResponses() *AppQueryUpdate {
	aqu.mutation.ClearResponses()
	return aqu
}

// RemoveResponseIDs removes the "responses" edge to AppResponse entities by IDs.
func (aqu *AppQueryUpdate) RemoveResponseIDs(ids ...uuid.UUID) *AppQueryUpdate {
	aqu.mutation.RemoveResponseIDs(ids...)
	return aqu
}

// RemoveResponses removes "responses" edges to AppResponse entities.
func (aqu *AppQueryUpdate) RemoveResponses(a ...*AppResponse) *AppQueryUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aqu.RemoveResponseIDs(ids...)
}

// ClearIpaddress clears the "ipaddress" edge to the IP entity.
func (aqu *AppQueryUpdate) ClearIpaddress() *AppQueryUpdate {
	aqu.mutation.ClearIpaddress()
	return aqu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aqu *AppQueryUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	aqu.defaults()
	if len(aqu.hooks) == 0 {
		if err = aqu.check(); err != nil {
			return 0, err
		}
		affected, err = aqu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppQueryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = aqu.check(); err != nil {
				return 0, err
			}
			aqu.mutation = mutation
			affected, err = aqu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aqu.hooks) - 1; i >= 0; i-- {
			if aqu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aqu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aqu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (aqu *AppQueryUpdate) SaveX(ctx context.Context) int {
	affected, err := aqu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aqu *AppQueryUpdate) Exec(ctx context.Context) error {
	_, err := aqu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aqu *AppQueryUpdate) ExecX(ctx context.Context) {
	if err := aqu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aqu *AppQueryUpdate) defaults() {
	if _, ok := aqu.mutation.UpdatedAt(); !ok {
		v := appquery.UpdateDefaultUpdatedAt()
		aqu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aqu *AppQueryUpdate) check() error {
	if _, ok := aqu.mutation.IpaddressID(); aqu.mutation.IpaddressCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"ipaddress\"")
	}
	return nil
}

func (aqu *AppQueryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appquery.Table,
			Columns: appquery.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appquery.FieldID,
			},
		},
	}
	if ps := aqu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aqu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: appquery.FieldCreatedAt,
		})
	}
	if value, ok := aqu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: appquery.FieldUpdatedAt,
		})
	}
	if aqu.mutation.ResponsesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   appquery.ResponsesTable,
			Columns: []string{appquery.ResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: appresponse.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aqu.mutation.RemovedResponsesIDs(); len(nodes) > 0 && !aqu.mutation.ResponsesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   appquery.ResponsesTable,
			Columns: []string{appquery.ResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: appresponse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aqu.mutation.ResponsesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   appquery.ResponsesTable,
			Columns: []string{appquery.ResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: appresponse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if aqu.mutation.IpaddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appquery.IpaddressTable,
			Columns: []string{appquery.IpaddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ip.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aqu.mutation.IpaddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appquery.IpaddressTable,
			Columns: []string{appquery.IpaddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ip.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aqu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appquery.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AppQueryUpdateOne is the builder for updating a single AppQuery entity.
type AppQueryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AppQueryMutation
}

// SetCreatedAt sets the "created_at" field.
func (aquo *AppQueryUpdateOne) SetCreatedAt(t time.Time) *AppQueryUpdateOne {
	aquo.mutation.SetCreatedAt(t)
	return aquo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aquo *AppQueryUpdateOne) SetNillableCreatedAt(t *time.Time) *AppQueryUpdateOne {
	if t != nil {
		aquo.SetCreatedAt(*t)
	}
	return aquo
}

// SetUpdatedAt sets the "updated_at" field.
func (aquo *AppQueryUpdateOne) SetUpdatedAt(t time.Time) *AppQueryUpdateOne {
	aquo.mutation.SetUpdatedAt(t)
	return aquo
}

// AddResponseIDs adds the "responses" edge to the AppResponse entity by IDs.
func (aquo *AppQueryUpdateOne) AddResponseIDs(ids ...uuid.UUID) *AppQueryUpdateOne {
	aquo.mutation.AddResponseIDs(ids...)
	return aquo
}

// AddResponses adds the "responses" edges to the AppResponse entity.
func (aquo *AppQueryUpdateOne) AddResponses(a ...*AppResponse) *AppQueryUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aquo.AddResponseIDs(ids...)
}

// SetIpaddressID sets the "ipaddress" edge to the IP entity by ID.
func (aquo *AppQueryUpdateOne) SetIpaddressID(id uuid.UUID) *AppQueryUpdateOne {
	aquo.mutation.SetIpaddressID(id)
	return aquo
}

// SetIpaddress sets the "ipaddress" edge to the IP entity.
func (aquo *AppQueryUpdateOne) SetIpaddress(i *IP) *AppQueryUpdateOne {
	return aquo.SetIpaddressID(i.ID)
}

// Mutation returns the AppQueryMutation object of the builder.
func (aquo *AppQueryUpdateOne) Mutation() *AppQueryMutation {
	return aquo.mutation
}

// ClearResponses clears all "responses" edges to the AppResponse entity.
func (aquo *AppQueryUpdateOne) ClearResponses() *AppQueryUpdateOne {
	aquo.mutation.ClearResponses()
	return aquo
}

// RemoveResponseIDs removes the "responses" edge to AppResponse entities by IDs.
func (aquo *AppQueryUpdateOne) RemoveResponseIDs(ids ...uuid.UUID) *AppQueryUpdateOne {
	aquo.mutation.RemoveResponseIDs(ids...)
	return aquo
}

// RemoveResponses removes "responses" edges to AppResponse entities.
func (aquo *AppQueryUpdateOne) RemoveResponses(a ...*AppResponse) *AppQueryUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aquo.RemoveResponseIDs(ids...)
}

// ClearIpaddress clears the "ipaddress" edge to the IP entity.
func (aquo *AppQueryUpdateOne) ClearIpaddress() *AppQueryUpdateOne {
	aquo.mutation.ClearIpaddress()
	return aquo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aquo *AppQueryUpdateOne) Select(field string, fields ...string) *AppQueryUpdateOne {
	aquo.fields = append([]string{field}, fields...)
	return aquo
}

// Save executes the query and returns the updated AppQuery entity.
func (aquo *AppQueryUpdateOne) Save(ctx context.Context) (*AppQuery, error) {
	var (
		err  error
		node *AppQuery
	)
	aquo.defaults()
	if len(aquo.hooks) == 0 {
		if err = aquo.check(); err != nil {
			return nil, err
		}
		node, err = aquo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppQueryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = aquo.check(); err != nil {
				return nil, err
			}
			aquo.mutation = mutation
			node, err = aquo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(aquo.hooks) - 1; i >= 0; i-- {
			if aquo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aquo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aquo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (aquo *AppQueryUpdateOne) SaveX(ctx context.Context) *AppQuery {
	node, err := aquo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aquo *AppQueryUpdateOne) Exec(ctx context.Context) error {
	_, err := aquo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aquo *AppQueryUpdateOne) ExecX(ctx context.Context) {
	if err := aquo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aquo *AppQueryUpdateOne) defaults() {
	if _, ok := aquo.mutation.UpdatedAt(); !ok {
		v := appquery.UpdateDefaultUpdatedAt()
		aquo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aquo *AppQueryUpdateOne) check() error {
	if _, ok := aquo.mutation.IpaddressID(); aquo.mutation.IpaddressCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"ipaddress\"")
	}
	return nil
}

func (aquo *AppQueryUpdateOne) sqlSave(ctx context.Context) (_node *AppQuery, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appquery.Table,
			Columns: appquery.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appquery.FieldID,
			},
		},
	}
	id, ok := aquo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing AppQuery.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := aquo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appquery.FieldID)
		for _, f := range fields {
			if !appquery.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appquery.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aquo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aquo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: appquery.FieldCreatedAt,
		})
	}
	if value, ok := aquo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: appquery.FieldUpdatedAt,
		})
	}
	if aquo.mutation.ResponsesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   appquery.ResponsesTable,
			Columns: []string{appquery.ResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: appresponse.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aquo.mutation.RemovedResponsesIDs(); len(nodes) > 0 && !aquo.mutation.ResponsesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   appquery.ResponsesTable,
			Columns: []string{appquery.ResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: appresponse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aquo.mutation.ResponsesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   appquery.ResponsesTable,
			Columns: []string{appquery.ResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: appresponse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if aquo.mutation.IpaddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appquery.IpaddressTable,
			Columns: []string{appquery.IpaddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ip.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aquo.mutation.IpaddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appquery.IpaddressTable,
			Columns: []string{appquery.IpaddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ip.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AppQuery{config: aquo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aquo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appquery.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
