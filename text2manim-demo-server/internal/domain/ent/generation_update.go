// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"text2manim-demo-server/internal/domain/ent/generation"
	"text2manim-demo-server/internal/domain/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GenerationUpdate is the builder for updating Generation entities.
type GenerationUpdate struct {
	config
	hooks    []Hook
	mutation *GenerationMutation
}

// Where appends a list predicates to the GenerationUpdate builder.
func (gu *GenerationUpdate) Where(ps ...predicate.Generation) *GenerationUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetRequestID sets the "request_id" field.
func (gu *GenerationUpdate) SetRequestID(s string) *GenerationUpdate {
	gu.mutation.SetRequestID(s)
	return gu
}

// SetNillableRequestID sets the "request_id" field if the given value is not nil.
func (gu *GenerationUpdate) SetNillableRequestID(s *string) *GenerationUpdate {
	if s != nil {
		gu.SetRequestID(*s)
	}
	return gu
}

// SetPrompt sets the "prompt" field.
func (gu *GenerationUpdate) SetPrompt(s string) *GenerationUpdate {
	gu.mutation.SetPrompt(s)
	return gu
}

// SetNillablePrompt sets the "prompt" field if the given value is not nil.
func (gu *GenerationUpdate) SetNillablePrompt(s *string) *GenerationUpdate {
	if s != nil {
		gu.SetPrompt(*s)
	}
	return gu
}

// SetStatus sets the "status" field.
func (gu *GenerationUpdate) SetStatus(s string) *GenerationUpdate {
	gu.mutation.SetStatus(s)
	return gu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (gu *GenerationUpdate) SetNillableStatus(s *string) *GenerationUpdate {
	if s != nil {
		gu.SetStatus(*s)
	}
	return gu
}

// SetVideoURL sets the "video_url" field.
func (gu *GenerationUpdate) SetVideoURL(s string) *GenerationUpdate {
	gu.mutation.SetVideoURL(s)
	return gu
}

// SetNillableVideoURL sets the "video_url" field if the given value is not nil.
func (gu *GenerationUpdate) SetNillableVideoURL(s *string) *GenerationUpdate {
	if s != nil {
		gu.SetVideoURL(*s)
	}
	return gu
}

// SetScriptURL sets the "script_url" field.
func (gu *GenerationUpdate) SetScriptURL(s string) *GenerationUpdate {
	gu.mutation.SetScriptURL(s)
	return gu
}

// SetNillableScriptURL sets the "script_url" field if the given value is not nil.
func (gu *GenerationUpdate) SetNillableScriptURL(s *string) *GenerationUpdate {
	if s != nil {
		gu.SetScriptURL(*s)
	}
	return gu
}

// SetErrorMessage sets the "error_message" field.
func (gu *GenerationUpdate) SetErrorMessage(s string) *GenerationUpdate {
	gu.mutation.SetErrorMessage(s)
	return gu
}

// SetNillableErrorMessage sets the "error_message" field if the given value is not nil.
func (gu *GenerationUpdate) SetNillableErrorMessage(s *string) *GenerationUpdate {
	if s != nil {
		gu.SetErrorMessage(*s)
	}
	return gu
}

// SetEmail sets the "email" field.
func (gu *GenerationUpdate) SetEmail(s string) *GenerationUpdate {
	gu.mutation.SetEmail(s)
	return gu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (gu *GenerationUpdate) SetNillableEmail(s *string) *GenerationUpdate {
	if s != nil {
		gu.SetEmail(*s)
	}
	return gu
}

// SetUpdatedAt sets the "updated_at" field.
func (gu *GenerationUpdate) SetUpdatedAt(t time.Time) *GenerationUpdate {
	gu.mutation.SetUpdatedAt(t)
	return gu
}

// Mutation returns the GenerationMutation object of the builder.
func (gu *GenerationUpdate) Mutation() *GenerationMutation {
	return gu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GenerationUpdate) Save(ctx context.Context) (int, error) {
	gu.defaults()
	return withHooks(ctx, gu.sqlSave, gu.mutation, gu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GenerationUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GenerationUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GenerationUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gu *GenerationUpdate) defaults() {
	if _, ok := gu.mutation.UpdatedAt(); !ok {
		v := generation.UpdateDefaultUpdatedAt()
		gu.mutation.SetUpdatedAt(v)
	}
}

func (gu *GenerationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(generation.Table, generation.Columns, sqlgraph.NewFieldSpec(generation.FieldID, field.TypeUUID))
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.RequestID(); ok {
		_spec.SetField(generation.FieldRequestID, field.TypeString, value)
	}
	if value, ok := gu.mutation.Prompt(); ok {
		_spec.SetField(generation.FieldPrompt, field.TypeString, value)
	}
	if value, ok := gu.mutation.Status(); ok {
		_spec.SetField(generation.FieldStatus, field.TypeString, value)
	}
	if value, ok := gu.mutation.VideoURL(); ok {
		_spec.SetField(generation.FieldVideoURL, field.TypeString, value)
	}
	if value, ok := gu.mutation.ScriptURL(); ok {
		_spec.SetField(generation.FieldScriptURL, field.TypeString, value)
	}
	if value, ok := gu.mutation.ErrorMessage(); ok {
		_spec.SetField(generation.FieldErrorMessage, field.TypeString, value)
	}
	if value, ok := gu.mutation.Email(); ok {
		_spec.SetField(generation.FieldEmail, field.TypeString, value)
	}
	if value, ok := gu.mutation.UpdatedAt(); ok {
		_spec.SetField(generation.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{generation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gu.mutation.done = true
	return n, nil
}

// GenerationUpdateOne is the builder for updating a single Generation entity.
type GenerationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GenerationMutation
}

// SetRequestID sets the "request_id" field.
func (guo *GenerationUpdateOne) SetRequestID(s string) *GenerationUpdateOne {
	guo.mutation.SetRequestID(s)
	return guo
}

// SetNillableRequestID sets the "request_id" field if the given value is not nil.
func (guo *GenerationUpdateOne) SetNillableRequestID(s *string) *GenerationUpdateOne {
	if s != nil {
		guo.SetRequestID(*s)
	}
	return guo
}

// SetPrompt sets the "prompt" field.
func (guo *GenerationUpdateOne) SetPrompt(s string) *GenerationUpdateOne {
	guo.mutation.SetPrompt(s)
	return guo
}

// SetNillablePrompt sets the "prompt" field if the given value is not nil.
func (guo *GenerationUpdateOne) SetNillablePrompt(s *string) *GenerationUpdateOne {
	if s != nil {
		guo.SetPrompt(*s)
	}
	return guo
}

// SetStatus sets the "status" field.
func (guo *GenerationUpdateOne) SetStatus(s string) *GenerationUpdateOne {
	guo.mutation.SetStatus(s)
	return guo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (guo *GenerationUpdateOne) SetNillableStatus(s *string) *GenerationUpdateOne {
	if s != nil {
		guo.SetStatus(*s)
	}
	return guo
}

// SetVideoURL sets the "video_url" field.
func (guo *GenerationUpdateOne) SetVideoURL(s string) *GenerationUpdateOne {
	guo.mutation.SetVideoURL(s)
	return guo
}

// SetNillableVideoURL sets the "video_url" field if the given value is not nil.
func (guo *GenerationUpdateOne) SetNillableVideoURL(s *string) *GenerationUpdateOne {
	if s != nil {
		guo.SetVideoURL(*s)
	}
	return guo
}

// SetScriptURL sets the "script_url" field.
func (guo *GenerationUpdateOne) SetScriptURL(s string) *GenerationUpdateOne {
	guo.mutation.SetScriptURL(s)
	return guo
}

// SetNillableScriptURL sets the "script_url" field if the given value is not nil.
func (guo *GenerationUpdateOne) SetNillableScriptURL(s *string) *GenerationUpdateOne {
	if s != nil {
		guo.SetScriptURL(*s)
	}
	return guo
}

// SetErrorMessage sets the "error_message" field.
func (guo *GenerationUpdateOne) SetErrorMessage(s string) *GenerationUpdateOne {
	guo.mutation.SetErrorMessage(s)
	return guo
}

// SetNillableErrorMessage sets the "error_message" field if the given value is not nil.
func (guo *GenerationUpdateOne) SetNillableErrorMessage(s *string) *GenerationUpdateOne {
	if s != nil {
		guo.SetErrorMessage(*s)
	}
	return guo
}

// SetEmail sets the "email" field.
func (guo *GenerationUpdateOne) SetEmail(s string) *GenerationUpdateOne {
	guo.mutation.SetEmail(s)
	return guo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (guo *GenerationUpdateOne) SetNillableEmail(s *string) *GenerationUpdateOne {
	if s != nil {
		guo.SetEmail(*s)
	}
	return guo
}

// SetUpdatedAt sets the "updated_at" field.
func (guo *GenerationUpdateOne) SetUpdatedAt(t time.Time) *GenerationUpdateOne {
	guo.mutation.SetUpdatedAt(t)
	return guo
}

// Mutation returns the GenerationMutation object of the builder.
func (guo *GenerationUpdateOne) Mutation() *GenerationMutation {
	return guo.mutation
}

// Where appends a list predicates to the GenerationUpdate builder.
func (guo *GenerationUpdateOne) Where(ps ...predicate.Generation) *GenerationUpdateOne {
	guo.mutation.Where(ps...)
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GenerationUpdateOne) Select(field string, fields ...string) *GenerationUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Generation entity.
func (guo *GenerationUpdateOne) Save(ctx context.Context) (*Generation, error) {
	guo.defaults()
	return withHooks(ctx, guo.sqlSave, guo.mutation, guo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GenerationUpdateOne) SaveX(ctx context.Context) *Generation {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GenerationUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GenerationUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (guo *GenerationUpdateOne) defaults() {
	if _, ok := guo.mutation.UpdatedAt(); !ok {
		v := generation.UpdateDefaultUpdatedAt()
		guo.mutation.SetUpdatedAt(v)
	}
}

func (guo *GenerationUpdateOne) sqlSave(ctx context.Context) (_node *Generation, err error) {
	_spec := sqlgraph.NewUpdateSpec(generation.Table, generation.Columns, sqlgraph.NewFieldSpec(generation.FieldID, field.TypeUUID))
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Generation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, generation.FieldID)
		for _, f := range fields {
			if !generation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != generation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.RequestID(); ok {
		_spec.SetField(generation.FieldRequestID, field.TypeString, value)
	}
	if value, ok := guo.mutation.Prompt(); ok {
		_spec.SetField(generation.FieldPrompt, field.TypeString, value)
	}
	if value, ok := guo.mutation.Status(); ok {
		_spec.SetField(generation.FieldStatus, field.TypeString, value)
	}
	if value, ok := guo.mutation.VideoURL(); ok {
		_spec.SetField(generation.FieldVideoURL, field.TypeString, value)
	}
	if value, ok := guo.mutation.ScriptURL(); ok {
		_spec.SetField(generation.FieldScriptURL, field.TypeString, value)
	}
	if value, ok := guo.mutation.ErrorMessage(); ok {
		_spec.SetField(generation.FieldErrorMessage, field.TypeString, value)
	}
	if value, ok := guo.mutation.Email(); ok {
		_spec.SetField(generation.FieldEmail, field.TypeString, value)
	}
	if value, ok := guo.mutation.UpdatedAt(); ok {
		_spec.SetField(generation.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Generation{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{generation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	guo.mutation.done = true
	return _node, nil
}