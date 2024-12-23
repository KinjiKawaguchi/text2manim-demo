// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"text2manim-demo-server/internal/domain/ent/generation"
	"text2manim-demo-server/internal/domain/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GenerationDelete is the builder for deleting a Generation entity.
type GenerationDelete struct {
	config
	hooks    []Hook
	mutation *GenerationMutation
}

// Where appends a list predicates to the GenerationDelete builder.
func (gd *GenerationDelete) Where(ps ...predicate.Generation) *GenerationDelete {
	gd.mutation.Where(ps...)
	return gd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gd *GenerationDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, gd.sqlExec, gd.mutation, gd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (gd *GenerationDelete) ExecX(ctx context.Context) int {
	n, err := gd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gd *GenerationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(generation.Table, sqlgraph.NewFieldSpec(generation.FieldID, field.TypeUUID))
	if ps := gd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, gd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	gd.mutation.done = true
	return affected, err
}

// GenerationDeleteOne is the builder for deleting a single Generation entity.
type GenerationDeleteOne struct {
	gd *GenerationDelete
}

// Where appends a list predicates to the GenerationDelete builder.
func (gdo *GenerationDeleteOne) Where(ps ...predicate.Generation) *GenerationDeleteOne {
	gdo.gd.mutation.Where(ps...)
	return gdo
}

// Exec executes the deletion query.
func (gdo *GenerationDeleteOne) Exec(ctx context.Context) error {
	n, err := gdo.gd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{generation.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gdo *GenerationDeleteOne) ExecX(ctx context.Context) {
	if err := gdo.Exec(ctx); err != nil {
		panic(err)
	}
}
