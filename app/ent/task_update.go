// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sujithshajee/dnsbl/app/ent/predicate"
	"github.com/sujithshajee/dnsbl/app/ent/task"
)

// TaskUpdate is the builder for updating Task entities.
type TaskUpdate struct {
	config
	hooks    []Hook
	mutation *TaskMutation
}

// Where appends a list predicates to the TaskUpdate builder.
func (tu *TaskUpdate) Where(ps ...predicate.Task) *TaskUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetCreatedAt sets the "created_at" field.
func (tu *TaskUpdate) SetCreatedAt(t time.Time) *TaskUpdate {
	tu.mutation.SetCreatedAt(t)
	return tu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableCreatedAt(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetCreatedAt(*t)
	}
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TaskUpdate) SetUpdatedAt(t time.Time) *TaskUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetType sets the "type" field.
func (tu *TaskUpdate) SetType(t task.Type) *TaskUpdate {
	tu.mutation.SetType(t)
	return tu
}

// SetIpaddress sets the "ipaddress" field.
func (tu *TaskUpdate) SetIpaddress(s string) *TaskUpdate {
	tu.mutation.SetIpaddress(s)
	return tu
}

// SetNillableIpaddress sets the "ipaddress" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableIpaddress(s *string) *TaskUpdate {
	if s != nil {
		tu.SetIpaddress(*s)
	}
	return tu
}

// ClearIpaddress clears the value of the "ipaddress" field.
func (tu *TaskUpdate) ClearIpaddress() *TaskUpdate {
	tu.mutation.ClearIpaddress()
	return tu
}

// SetStatus sets the "status" field.
func (tu *TaskUpdate) SetStatus(t task.Status) *TaskUpdate {
	tu.mutation.SetStatus(t)
	return tu
}

// SetError sets the "error" field.
func (tu *TaskUpdate) SetError(s string) *TaskUpdate {
	tu.mutation.SetError(s)
	return tu
}

// SetNillableError sets the "error" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableError(s *string) *TaskUpdate {
	if s != nil {
		tu.SetError(*s)
	}
	return tu
}

// ClearError clears the value of the "error" field.
func (tu *TaskUpdate) ClearError() *TaskUpdate {
	tu.mutation.ClearError()
	return tu
}

// SetCompletedBy sets the "completed_by" field.
func (tu *TaskUpdate) SetCompletedBy(t time.Time) *TaskUpdate {
	tu.mutation.SetCompletedBy(t)
	return tu
}

// SetNillableCompletedBy sets the "completed_by" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableCompletedBy(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetCompletedBy(*t)
	}
	return tu
}

// ClearCompletedBy clears the value of the "completed_by" field.
func (tu *TaskUpdate) ClearCompletedBy() *TaskUpdate {
	tu.mutation.ClearCompletedBy()
	return tu
}

// Mutation returns the TaskMutation object of the builder.
func (tu *TaskUpdate) Mutation() *TaskMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TaskUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tu.defaults()
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TaskUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TaskUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TaskUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TaskUpdate) defaults() {
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		v := task.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TaskUpdate) check() error {
	if v, ok := tu.mutation.GetType(); ok {
		if err := task.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if v, ok := tu.mutation.Status(); ok {
		if err := task.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (tu *TaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   task.Table,
			Columns: task.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: task.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldCreatedAt,
		})
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldUpdatedAt,
		})
	}
	if value, ok := tu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: task.FieldType,
		})
	}
	if value, ok := tu.mutation.Ipaddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldIpaddress,
		})
	}
	if tu.mutation.IpaddressCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: task.FieldIpaddress,
		})
	}
	if value, ok := tu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: task.FieldStatus,
		})
	}
	if value, ok := tu.mutation.Error(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldError,
		})
	}
	if tu.mutation.ErrorCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: task.FieldError,
		})
	}
	if value, ok := tu.mutation.CompletedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldCompletedBy,
		})
	}
	if tu.mutation.CompletedByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: task.FieldCompletedBy,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TaskUpdateOne is the builder for updating a single Task entity.
type TaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaskMutation
}

// SetCreatedAt sets the "created_at" field.
func (tuo *TaskUpdateOne) SetCreatedAt(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetCreatedAt(t)
	return tuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableCreatedAt(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetCreatedAt(*t)
	}
	return tuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TaskUpdateOne) SetUpdatedAt(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetType sets the "type" field.
func (tuo *TaskUpdateOne) SetType(t task.Type) *TaskUpdateOne {
	tuo.mutation.SetType(t)
	return tuo
}

// SetIpaddress sets the "ipaddress" field.
func (tuo *TaskUpdateOne) SetIpaddress(s string) *TaskUpdateOne {
	tuo.mutation.SetIpaddress(s)
	return tuo
}

// SetNillableIpaddress sets the "ipaddress" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableIpaddress(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetIpaddress(*s)
	}
	return tuo
}

// ClearIpaddress clears the value of the "ipaddress" field.
func (tuo *TaskUpdateOne) ClearIpaddress() *TaskUpdateOne {
	tuo.mutation.ClearIpaddress()
	return tuo
}

// SetStatus sets the "status" field.
func (tuo *TaskUpdateOne) SetStatus(t task.Status) *TaskUpdateOne {
	tuo.mutation.SetStatus(t)
	return tuo
}

// SetError sets the "error" field.
func (tuo *TaskUpdateOne) SetError(s string) *TaskUpdateOne {
	tuo.mutation.SetError(s)
	return tuo
}

// SetNillableError sets the "error" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableError(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetError(*s)
	}
	return tuo
}

// ClearError clears the value of the "error" field.
func (tuo *TaskUpdateOne) ClearError() *TaskUpdateOne {
	tuo.mutation.ClearError()
	return tuo
}

// SetCompletedBy sets the "completed_by" field.
func (tuo *TaskUpdateOne) SetCompletedBy(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetCompletedBy(t)
	return tuo
}

// SetNillableCompletedBy sets the "completed_by" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableCompletedBy(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetCompletedBy(*t)
	}
	return tuo
}

// ClearCompletedBy clears the value of the "completed_by" field.
func (tuo *TaskUpdateOne) ClearCompletedBy() *TaskUpdateOne {
	tuo.mutation.ClearCompletedBy()
	return tuo
}

// Mutation returns the TaskMutation object of the builder.
func (tuo *TaskUpdateOne) Mutation() *TaskMutation {
	return tuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TaskUpdateOne) Select(field string, fields ...string) *TaskUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Task entity.
func (tuo *TaskUpdateOne) Save(ctx context.Context) (*Task, error) {
	var (
		err  error
		node *Task
	)
	tuo.defaults()
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TaskUpdateOne) SaveX(ctx context.Context) *Task {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TaskUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TaskUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TaskUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		v := task.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TaskUpdateOne) check() error {
	if v, ok := tuo.mutation.GetType(); ok {
		if err := task.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if v, ok := tuo.mutation.Status(); ok {
		if err := task.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (tuo *TaskUpdateOne) sqlSave(ctx context.Context) (_node *Task, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   task.Table,
			Columns: task.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: task.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Task.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, task.FieldID)
		for _, f := range fields {
			if !task.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != task.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldCreatedAt,
		})
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldUpdatedAt,
		})
	}
	if value, ok := tuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: task.FieldType,
		})
	}
	if value, ok := tuo.mutation.Ipaddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldIpaddress,
		})
	}
	if tuo.mutation.IpaddressCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: task.FieldIpaddress,
		})
	}
	if value, ok := tuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: task.FieldStatus,
		})
	}
	if value, ok := tuo.mutation.Error(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldError,
		})
	}
	if tuo.mutation.ErrorCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: task.FieldError,
		})
	}
	if value, ok := tuo.mutation.CompletedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldCompletedBy,
		})
	}
	if tuo.mutation.CompletedByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: task.FieldCompletedBy,
		})
	}
	_node = &Task{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
