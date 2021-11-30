// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/sujithshajee/dnsbl/app/ent/appquery"
	"github.com/sujithshajee/dnsbl/app/ent/appresponse"
	"github.com/sujithshajee/dnsbl/app/ent/ip"
)

// AppQueryCreate is the builder for creating a AppQuery entity.
type AppQueryCreate struct {
	config
	mutation *AppQueryMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (aqc *AppQueryCreate) SetCreatedAt(t time.Time) *AppQueryCreate {
	aqc.mutation.SetCreatedAt(t)
	return aqc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aqc *AppQueryCreate) SetNillableCreatedAt(t *time.Time) *AppQueryCreate {
	if t != nil {
		aqc.SetCreatedAt(*t)
	}
	return aqc
}

// SetUpdatedAt sets the "updated_at" field.
func (aqc *AppQueryCreate) SetUpdatedAt(t time.Time) *AppQueryCreate {
	aqc.mutation.SetUpdatedAt(t)
	return aqc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (aqc *AppQueryCreate) SetNillableUpdatedAt(t *time.Time) *AppQueryCreate {
	if t != nil {
		aqc.SetUpdatedAt(*t)
	}
	return aqc
}

// SetID sets the "id" field.
func (aqc *AppQueryCreate) SetID(u uuid.UUID) *AppQueryCreate {
	aqc.mutation.SetID(u)
	return aqc
}

// AddResponseIDs adds the "responses" edge to the AppResponse entity by IDs.
func (aqc *AppQueryCreate) AddResponseIDs(ids ...uuid.UUID) *AppQueryCreate {
	aqc.mutation.AddResponseIDs(ids...)
	return aqc
}

// AddResponses adds the "responses" edges to the AppResponse entity.
func (aqc *AppQueryCreate) AddResponses(a ...*AppResponse) *AppQueryCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aqc.AddResponseIDs(ids...)
}

// SetIpaddressID sets the "ipaddress" edge to the IP entity by ID.
func (aqc *AppQueryCreate) SetIpaddressID(id uuid.UUID) *AppQueryCreate {
	aqc.mutation.SetIpaddressID(id)
	return aqc
}

// SetIpaddress sets the "ipaddress" edge to the IP entity.
func (aqc *AppQueryCreate) SetIpaddress(i *IP) *AppQueryCreate {
	return aqc.SetIpaddressID(i.ID)
}

// Mutation returns the AppQueryMutation object of the builder.
func (aqc *AppQueryCreate) Mutation() *AppQueryMutation {
	return aqc.mutation
}

// Save creates the AppQuery in the database.
func (aqc *AppQueryCreate) Save(ctx context.Context) (*AppQuery, error) {
	var (
		err  error
		node *AppQuery
	)
	aqc.defaults()
	if len(aqc.hooks) == 0 {
		if err = aqc.check(); err != nil {
			return nil, err
		}
		node, err = aqc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppQueryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = aqc.check(); err != nil {
				return nil, err
			}
			aqc.mutation = mutation
			if node, err = aqc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(aqc.hooks) - 1; i >= 0; i-- {
			if aqc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aqc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aqc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (aqc *AppQueryCreate) SaveX(ctx context.Context) *AppQuery {
	v, err := aqc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aqc *AppQueryCreate) Exec(ctx context.Context) error {
	_, err := aqc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aqc *AppQueryCreate) ExecX(ctx context.Context) {
	if err := aqc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aqc *AppQueryCreate) defaults() {
	if _, ok := aqc.mutation.CreatedAt(); !ok {
		v := appquery.DefaultCreatedAt()
		aqc.mutation.SetCreatedAt(v)
	}
	if _, ok := aqc.mutation.UpdatedAt(); !ok {
		v := appquery.DefaultUpdatedAt()
		aqc.mutation.SetUpdatedAt(v)
	}
	if _, ok := aqc.mutation.ID(); !ok {
		v := appquery.DefaultID()
		aqc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aqc *AppQueryCreate) check() error {
	if _, ok := aqc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := aqc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := aqc.mutation.IpaddressID(); !ok {
		return &ValidationError{Name: "ipaddress", err: errors.New("ent: missing required edge \"ipaddress\"")}
	}
	return nil
}

func (aqc *AppQueryCreate) sqlSave(ctx context.Context) (*AppQuery, error) {
	_node, _spec := aqc.createSpec()
	if err := sqlgraph.CreateNode(ctx, aqc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (aqc *AppQueryCreate) createSpec() (*AppQuery, *sqlgraph.CreateSpec) {
	var (
		_node = &AppQuery{config: aqc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: appquery.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appquery.FieldID,
			},
		}
	)
	if id, ok := aqc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := aqc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: appquery.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := aqc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: appquery.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := aqc.mutation.ResponsesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := aqc.mutation.IpaddressIDs(); len(nodes) > 0 {
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
		_node.ip_queries = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AppQueryCreateBulk is the builder for creating many AppQuery entities in bulk.
type AppQueryCreateBulk struct {
	config
	builders []*AppQueryCreate
}

// Save creates the AppQuery entities in the database.
func (aqcb *AppQueryCreateBulk) Save(ctx context.Context) ([]*AppQuery, error) {
	specs := make([]*sqlgraph.CreateSpec, len(aqcb.builders))
	nodes := make([]*AppQuery, len(aqcb.builders))
	mutators := make([]Mutator, len(aqcb.builders))
	for i := range aqcb.builders {
		func(i int, root context.Context) {
			builder := aqcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppQueryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, aqcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, aqcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, aqcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (aqcb *AppQueryCreateBulk) SaveX(ctx context.Context) []*AppQuery {
	v, err := aqcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aqcb *AppQueryCreateBulk) Exec(ctx context.Context) error {
	_, err := aqcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aqcb *AppQueryCreateBulk) ExecX(ctx context.Context) {
	if err := aqcb.Exec(ctx); err != nil {
		panic(err)
	}
}