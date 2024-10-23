// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"text2manim-demo-server/internal/domain/ent/generation"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// GenerationCreate is the builder for creating a Generation entity.
type GenerationCreate struct {
	config
	mutation *GenerationMutation
	hooks    []Hook
}

// SetRequestID sets the "request_id" field.
func (gc *GenerationCreate) SetRequestID(s string) *GenerationCreate {
	gc.mutation.SetRequestID(s)
	return gc
}

// SetPrompt sets the "prompt" field.
func (gc *GenerationCreate) SetPrompt(s string) *GenerationCreate {
	gc.mutation.SetPrompt(s)
	return gc
}

// SetStatus sets the "status" field.
func (gc *GenerationCreate) SetStatus(s string) *GenerationCreate {
	gc.mutation.SetStatus(s)
	return gc
}

// SetVideoURL sets the "video_url" field.
func (gc *GenerationCreate) SetVideoURL(s string) *GenerationCreate {
	gc.mutation.SetVideoURL(s)
	return gc
}

// SetScriptURL sets the "script_url" field.
func (gc *GenerationCreate) SetScriptURL(s string) *GenerationCreate {
	gc.mutation.SetScriptURL(s)
	return gc
}

// SetErrorMessage sets the "error_message" field.
func (gc *GenerationCreate) SetErrorMessage(s string) *GenerationCreate {
	gc.mutation.SetErrorMessage(s)
	return gc
}

// SetEmail sets the "email" field.
func (gc *GenerationCreate) SetEmail(s string) *GenerationCreate {
	gc.mutation.SetEmail(s)
	return gc
}

// SetCreatedAt sets the "created_at" field.
func (gc *GenerationCreate) SetCreatedAt(t time.Time) *GenerationCreate {
	gc.mutation.SetCreatedAt(t)
	return gc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gc *GenerationCreate) SetNillableCreatedAt(t *time.Time) *GenerationCreate {
	if t != nil {
		gc.SetCreatedAt(*t)
	}
	return gc
}

// SetUpdatedAt sets the "updated_at" field.
func (gc *GenerationCreate) SetUpdatedAt(t time.Time) *GenerationCreate {
	gc.mutation.SetUpdatedAt(t)
	return gc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gc *GenerationCreate) SetNillableUpdatedAt(t *time.Time) *GenerationCreate {
	if t != nil {
		gc.SetUpdatedAt(*t)
	}
	return gc
}

// SetID sets the "id" field.
func (gc *GenerationCreate) SetID(u uuid.UUID) *GenerationCreate {
	gc.mutation.SetID(u)
	return gc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (gc *GenerationCreate) SetNillableID(u *uuid.UUID) *GenerationCreate {
	if u != nil {
		gc.SetID(*u)
	}
	return gc
}

// Mutation returns the GenerationMutation object of the builder.
func (gc *GenerationCreate) Mutation() *GenerationMutation {
	return gc.mutation
}

// Save creates the Generation in the database.
func (gc *GenerationCreate) Save(ctx context.Context) (*Generation, error) {
	gc.defaults()
	return withHooks(ctx, gc.sqlSave, gc.mutation, gc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GenerationCreate) SaveX(ctx context.Context) *Generation {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gc *GenerationCreate) Exec(ctx context.Context) error {
	_, err := gc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gc *GenerationCreate) ExecX(ctx context.Context) {
	if err := gc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gc *GenerationCreate) defaults() {
	if _, ok := gc.mutation.CreatedAt(); !ok {
		v := generation.DefaultCreatedAt()
		gc.mutation.SetCreatedAt(v)
	}
	if _, ok := gc.mutation.UpdatedAt(); !ok {
		v := generation.DefaultUpdatedAt()
		gc.mutation.SetUpdatedAt(v)
	}
	if _, ok := gc.mutation.ID(); !ok {
		v := generation.DefaultID()
		gc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gc *GenerationCreate) check() error {
	if _, ok := gc.mutation.RequestID(); !ok {
		return &ValidationError{Name: "request_id", err: errors.New(`ent: missing required field "Generation.request_id"`)}
	}
	if _, ok := gc.mutation.Prompt(); !ok {
		return &ValidationError{Name: "prompt", err: errors.New(`ent: missing required field "Generation.prompt"`)}
	}
	if _, ok := gc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Generation.status"`)}
	}
	if _, ok := gc.mutation.VideoURL(); !ok {
		return &ValidationError{Name: "video_url", err: errors.New(`ent: missing required field "Generation.video_url"`)}
	}
	if _, ok := gc.mutation.ScriptURL(); !ok {
		return &ValidationError{Name: "script_url", err: errors.New(`ent: missing required field "Generation.script_url"`)}
	}
	if _, ok := gc.mutation.ErrorMessage(); !ok {
		return &ValidationError{Name: "error_message", err: errors.New(`ent: missing required field "Generation.error_message"`)}
	}
	if _, ok := gc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Generation.email"`)}
	}
	if _, ok := gc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Generation.created_at"`)}
	}
	if _, ok := gc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Generation.updated_at"`)}
	}
	return nil
}

func (gc *GenerationCreate) sqlSave(ctx context.Context) (*Generation, error) {
	if err := gc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	gc.mutation.id = &_node.ID
	gc.mutation.done = true
	return _node, nil
}

func (gc *GenerationCreate) createSpec() (*Generation, *sqlgraph.CreateSpec) {
	var (
		_node = &Generation{config: gc.config}
		_spec = sqlgraph.NewCreateSpec(generation.Table, sqlgraph.NewFieldSpec(generation.FieldID, field.TypeUUID))
	)
	if id, ok := gc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := gc.mutation.RequestID(); ok {
		_spec.SetField(generation.FieldRequestID, field.TypeString, value)
		_node.RequestID = value
	}
	if value, ok := gc.mutation.Prompt(); ok {
		_spec.SetField(generation.FieldPrompt, field.TypeString, value)
		_node.Prompt = value
	}
	if value, ok := gc.mutation.Status(); ok {
		_spec.SetField(generation.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := gc.mutation.VideoURL(); ok {
		_spec.SetField(generation.FieldVideoURL, field.TypeString, value)
		_node.VideoURL = value
	}
	if value, ok := gc.mutation.ScriptURL(); ok {
		_spec.SetField(generation.FieldScriptURL, field.TypeString, value)
		_node.ScriptURL = value
	}
	if value, ok := gc.mutation.ErrorMessage(); ok {
		_spec.SetField(generation.FieldErrorMessage, field.TypeString, value)
		_node.ErrorMessage = value
	}
	if value, ok := gc.mutation.Email(); ok {
		_spec.SetField(generation.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := gc.mutation.CreatedAt(); ok {
		_spec.SetField(generation.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := gc.mutation.UpdatedAt(); ok {
		_spec.SetField(generation.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// GenerationCreateBulk is the builder for creating many Generation entities in bulk.
type GenerationCreateBulk struct {
	config
	err      error
	builders []*GenerationCreate
}

// Save creates the Generation entities in the database.
func (gcb *GenerationCreateBulk) Save(ctx context.Context) ([]*Generation, error) {
	if gcb.err != nil {
		return nil, gcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Generation, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GenerationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
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
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GenerationCreateBulk) SaveX(ctx context.Context) []*Generation {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcb *GenerationCreateBulk) Exec(ctx context.Context) error {
	_, err := gcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcb *GenerationCreateBulk) ExecX(ctx context.Context) {
	if err := gcb.Exec(ctx); err != nil {
		panic(err)
	}
}