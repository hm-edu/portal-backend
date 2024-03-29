// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hm-edu/eab-rest-interface/ent/eabkey"
	"github.com/hm-edu/eab-rest-interface/ent/predicate"
)

// EABKeyUpdate is the builder for updating EABKey entities.
type EABKeyUpdate struct {
	config
	hooks    []Hook
	mutation *EABKeyMutation
}

// Where appends a list predicates to the EABKeyUpdate builder.
func (eku *EABKeyUpdate) Where(ps ...predicate.EABKey) *EABKeyUpdate {
	eku.mutation.Where(ps...)
	return eku
}

// SetUser sets the "user" field.
func (eku *EABKeyUpdate) SetUser(s string) *EABKeyUpdate {
	eku.mutation.SetUser(s)
	return eku
}

// SetEabKey sets the "eabKey" field.
func (eku *EABKeyUpdate) SetEabKey(s string) *EABKeyUpdate {
	eku.mutation.SetEabKey(s)
	return eku
}

// SetComment sets the "comment" field.
func (eku *EABKeyUpdate) SetComment(s string) *EABKeyUpdate {
	eku.mutation.SetComment(s)
	return eku
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (eku *EABKeyUpdate) SetNillableComment(s *string) *EABKeyUpdate {
	if s != nil {
		eku.SetComment(*s)
	}
	return eku
}

// ClearComment clears the value of the "comment" field.
func (eku *EABKeyUpdate) ClearComment() *EABKeyUpdate {
	eku.mutation.ClearComment()
	return eku
}

// Mutation returns the EABKeyMutation object of the builder.
func (eku *EABKeyUpdate) Mutation() *EABKeyMutation {
	return eku.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eku *EABKeyUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(eku.hooks) == 0 {
		affected, err = eku.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EABKeyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			eku.mutation = mutation
			affected, err = eku.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eku.hooks) - 1; i >= 0; i-- {
			if eku.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eku.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eku.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eku *EABKeyUpdate) SaveX(ctx context.Context) int {
	affected, err := eku.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eku *EABKeyUpdate) Exec(ctx context.Context) error {
	_, err := eku.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eku *EABKeyUpdate) ExecX(ctx context.Context) {
	if err := eku.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eku *EABKeyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   eabkey.Table,
			Columns: eabkey.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: eabkey.FieldID,
			},
		},
	}
	if ps := eku.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eku.mutation.User(); ok {
		_spec.SetField(eabkey.FieldUser, field.TypeString, value)
	}
	if value, ok := eku.mutation.EabKey(); ok {
		_spec.SetField(eabkey.FieldEabKey, field.TypeString, value)
	}
	if value, ok := eku.mutation.Comment(); ok {
		_spec.SetField(eabkey.FieldComment, field.TypeString, value)
	}
	if eku.mutation.CommentCleared() {
		_spec.ClearField(eabkey.FieldComment, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eku.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{eabkey.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// EABKeyUpdateOne is the builder for updating a single EABKey entity.
type EABKeyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EABKeyMutation
}

// SetUser sets the "user" field.
func (ekuo *EABKeyUpdateOne) SetUser(s string) *EABKeyUpdateOne {
	ekuo.mutation.SetUser(s)
	return ekuo
}

// SetEabKey sets the "eabKey" field.
func (ekuo *EABKeyUpdateOne) SetEabKey(s string) *EABKeyUpdateOne {
	ekuo.mutation.SetEabKey(s)
	return ekuo
}

// SetComment sets the "comment" field.
func (ekuo *EABKeyUpdateOne) SetComment(s string) *EABKeyUpdateOne {
	ekuo.mutation.SetComment(s)
	return ekuo
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (ekuo *EABKeyUpdateOne) SetNillableComment(s *string) *EABKeyUpdateOne {
	if s != nil {
		ekuo.SetComment(*s)
	}
	return ekuo
}

// ClearComment clears the value of the "comment" field.
func (ekuo *EABKeyUpdateOne) ClearComment() *EABKeyUpdateOne {
	ekuo.mutation.ClearComment()
	return ekuo
}

// Mutation returns the EABKeyMutation object of the builder.
func (ekuo *EABKeyUpdateOne) Mutation() *EABKeyMutation {
	return ekuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ekuo *EABKeyUpdateOne) Select(field string, fields ...string) *EABKeyUpdateOne {
	ekuo.fields = append([]string{field}, fields...)
	return ekuo
}

// Save executes the query and returns the updated EABKey entity.
func (ekuo *EABKeyUpdateOne) Save(ctx context.Context) (*EABKey, error) {
	var (
		err  error
		node *EABKey
	)
	if len(ekuo.hooks) == 0 {
		node, err = ekuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EABKeyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ekuo.mutation = mutation
			node, err = ekuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ekuo.hooks) - 1; i >= 0; i-- {
			if ekuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ekuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ekuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*EABKey)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EABKeyMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ekuo *EABKeyUpdateOne) SaveX(ctx context.Context) *EABKey {
	node, err := ekuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ekuo *EABKeyUpdateOne) Exec(ctx context.Context) error {
	_, err := ekuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ekuo *EABKeyUpdateOne) ExecX(ctx context.Context) {
	if err := ekuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ekuo *EABKeyUpdateOne) sqlSave(ctx context.Context) (_node *EABKey, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   eabkey.Table,
			Columns: eabkey.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: eabkey.FieldID,
			},
		},
	}
	id, ok := ekuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EABKey.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ekuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, eabkey.FieldID)
		for _, f := range fields {
			if !eabkey.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != eabkey.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ekuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ekuo.mutation.User(); ok {
		_spec.SetField(eabkey.FieldUser, field.TypeString, value)
	}
	if value, ok := ekuo.mutation.EabKey(); ok {
		_spec.SetField(eabkey.FieldEabKey, field.TypeString, value)
	}
	if value, ok := ekuo.mutation.Comment(); ok {
		_spec.SetField(eabkey.FieldComment, field.TypeString, value)
	}
	if ekuo.mutation.CommentCleared() {
		_spec.ClearField(eabkey.FieldComment, field.TypeString)
	}
	_node = &EABKey{config: ekuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ekuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{eabkey.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
