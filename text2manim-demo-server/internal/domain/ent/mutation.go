// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"text2manim-demo-server/internal/domain/ent/generation"
	"text2manim-demo-server/internal/domain/ent/predicate"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeGeneration = "Generation"
)

// GenerationMutation represents an operation that mutates the Generation nodes in the graph.
type GenerationMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	request_id    *string
	prompt        *string
	status        *generation.Status
	video_url     *string
	script_url    *string
	error_message *string
	email         *string
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Generation, error)
	predicates    []predicate.Generation
}

var _ ent.Mutation = (*GenerationMutation)(nil)

// generationOption allows management of the mutation configuration using functional options.
type generationOption func(*GenerationMutation)

// newGenerationMutation creates new mutation for the Generation entity.
func newGenerationMutation(c config, op Op, opts ...generationOption) *GenerationMutation {
	m := &GenerationMutation{
		config:        c,
		op:            op,
		typ:           TypeGeneration,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withGenerationID sets the ID field of the mutation.
func withGenerationID(id uuid.UUID) generationOption {
	return func(m *GenerationMutation) {
		var (
			err   error
			once  sync.Once
			value *Generation
		)
		m.oldValue = func(ctx context.Context) (*Generation, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Generation.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withGeneration sets the old Generation of the mutation.
func withGeneration(node *Generation) generationOption {
	return func(m *GenerationMutation) {
		m.oldValue = func(context.Context) (*Generation, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m GenerationMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m GenerationMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Generation entities.
func (m *GenerationMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *GenerationMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *GenerationMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Generation.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetRequestID sets the "request_id" field.
func (m *GenerationMutation) SetRequestID(s string) {
	m.request_id = &s
}

// RequestID returns the value of the "request_id" field in the mutation.
func (m *GenerationMutation) RequestID() (r string, exists bool) {
	v := m.request_id
	if v == nil {
		return
	}
	return *v, true
}

// OldRequestID returns the old "request_id" field's value of the Generation entity.
// If the Generation object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GenerationMutation) OldRequestID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRequestID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRequestID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRequestID: %w", err)
	}
	return oldValue.RequestID, nil
}

// ClearRequestID clears the value of the "request_id" field.
func (m *GenerationMutation) ClearRequestID() {
	m.request_id = nil
	m.clearedFields[generation.FieldRequestID] = struct{}{}
}

// RequestIDCleared returns if the "request_id" field was cleared in this mutation.
func (m *GenerationMutation) RequestIDCleared() bool {
	_, ok := m.clearedFields[generation.FieldRequestID]
	return ok
}

// ResetRequestID resets all changes to the "request_id" field.
func (m *GenerationMutation) ResetRequestID() {
	m.request_id = nil
	delete(m.clearedFields, generation.FieldRequestID)
}

// SetPrompt sets the "prompt" field.
func (m *GenerationMutation) SetPrompt(s string) {
	m.prompt = &s
}

// Prompt returns the value of the "prompt" field in the mutation.
func (m *GenerationMutation) Prompt() (r string, exists bool) {
	v := m.prompt
	if v == nil {
		return
	}
	return *v, true
}

// OldPrompt returns the old "prompt" field's value of the Generation entity.
// If the Generation object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GenerationMutation) OldPrompt(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPrompt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPrompt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPrompt: %w", err)
	}
	return oldValue.Prompt, nil
}

// ClearPrompt clears the value of the "prompt" field.
func (m *GenerationMutation) ClearPrompt() {
	m.prompt = nil
	m.clearedFields[generation.FieldPrompt] = struct{}{}
}

// PromptCleared returns if the "prompt" field was cleared in this mutation.
func (m *GenerationMutation) PromptCleared() bool {
	_, ok := m.clearedFields[generation.FieldPrompt]
	return ok
}

// ResetPrompt resets all changes to the "prompt" field.
func (m *GenerationMutation) ResetPrompt() {
	m.prompt = nil
	delete(m.clearedFields, generation.FieldPrompt)
}

// SetStatus sets the "status" field.
func (m *GenerationMutation) SetStatus(ge generation.Status) {
	m.status = &ge
}

// Status returns the value of the "status" field in the mutation.
func (m *GenerationMutation) Status() (r generation.Status, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the Generation entity.
// If the Generation object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GenerationMutation) OldStatus(ctx context.Context) (v generation.Status, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// ResetStatus resets all changes to the "status" field.
func (m *GenerationMutation) ResetStatus() {
	m.status = nil
}

// SetVideoURL sets the "video_url" field.
func (m *GenerationMutation) SetVideoURL(s string) {
	m.video_url = &s
}

// VideoURL returns the value of the "video_url" field in the mutation.
func (m *GenerationMutation) VideoURL() (r string, exists bool) {
	v := m.video_url
	if v == nil {
		return
	}
	return *v, true
}

// OldVideoURL returns the old "video_url" field's value of the Generation entity.
// If the Generation object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GenerationMutation) OldVideoURL(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldVideoURL is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldVideoURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldVideoURL: %w", err)
	}
	return oldValue.VideoURL, nil
}

// ClearVideoURL clears the value of the "video_url" field.
func (m *GenerationMutation) ClearVideoURL() {
	m.video_url = nil
	m.clearedFields[generation.FieldVideoURL] = struct{}{}
}

// VideoURLCleared returns if the "video_url" field was cleared in this mutation.
func (m *GenerationMutation) VideoURLCleared() bool {
	_, ok := m.clearedFields[generation.FieldVideoURL]
	return ok
}

// ResetVideoURL resets all changes to the "video_url" field.
func (m *GenerationMutation) ResetVideoURL() {
	m.video_url = nil
	delete(m.clearedFields, generation.FieldVideoURL)
}

// SetScriptURL sets the "script_url" field.
func (m *GenerationMutation) SetScriptURL(s string) {
	m.script_url = &s
}

// ScriptURL returns the value of the "script_url" field in the mutation.
func (m *GenerationMutation) ScriptURL() (r string, exists bool) {
	v := m.script_url
	if v == nil {
		return
	}
	return *v, true
}

// OldScriptURL returns the old "script_url" field's value of the Generation entity.
// If the Generation object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GenerationMutation) OldScriptURL(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldScriptURL is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldScriptURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldScriptURL: %w", err)
	}
	return oldValue.ScriptURL, nil
}

// ClearScriptURL clears the value of the "script_url" field.
func (m *GenerationMutation) ClearScriptURL() {
	m.script_url = nil
	m.clearedFields[generation.FieldScriptURL] = struct{}{}
}

// ScriptURLCleared returns if the "script_url" field was cleared in this mutation.
func (m *GenerationMutation) ScriptURLCleared() bool {
	_, ok := m.clearedFields[generation.FieldScriptURL]
	return ok
}

// ResetScriptURL resets all changes to the "script_url" field.
func (m *GenerationMutation) ResetScriptURL() {
	m.script_url = nil
	delete(m.clearedFields, generation.FieldScriptURL)
}

// SetErrorMessage sets the "error_message" field.
func (m *GenerationMutation) SetErrorMessage(s string) {
	m.error_message = &s
}

// ErrorMessage returns the value of the "error_message" field in the mutation.
func (m *GenerationMutation) ErrorMessage() (r string, exists bool) {
	v := m.error_message
	if v == nil {
		return
	}
	return *v, true
}

// OldErrorMessage returns the old "error_message" field's value of the Generation entity.
// If the Generation object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GenerationMutation) OldErrorMessage(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldErrorMessage is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldErrorMessage requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldErrorMessage: %w", err)
	}
	return oldValue.ErrorMessage, nil
}

// ClearErrorMessage clears the value of the "error_message" field.
func (m *GenerationMutation) ClearErrorMessage() {
	m.error_message = nil
	m.clearedFields[generation.FieldErrorMessage] = struct{}{}
}

// ErrorMessageCleared returns if the "error_message" field was cleared in this mutation.
func (m *GenerationMutation) ErrorMessageCleared() bool {
	_, ok := m.clearedFields[generation.FieldErrorMessage]
	return ok
}

// ResetErrorMessage resets all changes to the "error_message" field.
func (m *GenerationMutation) ResetErrorMessage() {
	m.error_message = nil
	delete(m.clearedFields, generation.FieldErrorMessage)
}

// SetEmail sets the "email" field.
func (m *GenerationMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *GenerationMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the Generation entity.
// If the Generation object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GenerationMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ClearEmail clears the value of the "email" field.
func (m *GenerationMutation) ClearEmail() {
	m.email = nil
	m.clearedFields[generation.FieldEmail] = struct{}{}
}

// EmailCleared returns if the "email" field was cleared in this mutation.
func (m *GenerationMutation) EmailCleared() bool {
	_, ok := m.clearedFields[generation.FieldEmail]
	return ok
}

// ResetEmail resets all changes to the "email" field.
func (m *GenerationMutation) ResetEmail() {
	m.email = nil
	delete(m.clearedFields, generation.FieldEmail)
}

// SetCreatedAt sets the "created_at" field.
func (m *GenerationMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *GenerationMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Generation entity.
// If the Generation object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GenerationMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *GenerationMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *GenerationMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *GenerationMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Generation entity.
// If the Generation object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GenerationMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *GenerationMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the GenerationMutation builder.
func (m *GenerationMutation) Where(ps ...predicate.Generation) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the GenerationMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *GenerationMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Generation, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *GenerationMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *GenerationMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Generation).
func (m *GenerationMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *GenerationMutation) Fields() []string {
	fields := make([]string, 0, 9)
	if m.request_id != nil {
		fields = append(fields, generation.FieldRequestID)
	}
	if m.prompt != nil {
		fields = append(fields, generation.FieldPrompt)
	}
	if m.status != nil {
		fields = append(fields, generation.FieldStatus)
	}
	if m.video_url != nil {
		fields = append(fields, generation.FieldVideoURL)
	}
	if m.script_url != nil {
		fields = append(fields, generation.FieldScriptURL)
	}
	if m.error_message != nil {
		fields = append(fields, generation.FieldErrorMessage)
	}
	if m.email != nil {
		fields = append(fields, generation.FieldEmail)
	}
	if m.created_at != nil {
		fields = append(fields, generation.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, generation.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *GenerationMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case generation.FieldRequestID:
		return m.RequestID()
	case generation.FieldPrompt:
		return m.Prompt()
	case generation.FieldStatus:
		return m.Status()
	case generation.FieldVideoURL:
		return m.VideoURL()
	case generation.FieldScriptURL:
		return m.ScriptURL()
	case generation.FieldErrorMessage:
		return m.ErrorMessage()
	case generation.FieldEmail:
		return m.Email()
	case generation.FieldCreatedAt:
		return m.CreatedAt()
	case generation.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *GenerationMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case generation.FieldRequestID:
		return m.OldRequestID(ctx)
	case generation.FieldPrompt:
		return m.OldPrompt(ctx)
	case generation.FieldStatus:
		return m.OldStatus(ctx)
	case generation.FieldVideoURL:
		return m.OldVideoURL(ctx)
	case generation.FieldScriptURL:
		return m.OldScriptURL(ctx)
	case generation.FieldErrorMessage:
		return m.OldErrorMessage(ctx)
	case generation.FieldEmail:
		return m.OldEmail(ctx)
	case generation.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case generation.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Generation field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *GenerationMutation) SetField(name string, value ent.Value) error {
	switch name {
	case generation.FieldRequestID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRequestID(v)
		return nil
	case generation.FieldPrompt:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPrompt(v)
		return nil
	case generation.FieldStatus:
		v, ok := value.(generation.Status)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case generation.FieldVideoURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetVideoURL(v)
		return nil
	case generation.FieldScriptURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetScriptURL(v)
		return nil
	case generation.FieldErrorMessage:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetErrorMessage(v)
		return nil
	case generation.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case generation.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case generation.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Generation field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *GenerationMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *GenerationMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *GenerationMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Generation numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *GenerationMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(generation.FieldRequestID) {
		fields = append(fields, generation.FieldRequestID)
	}
	if m.FieldCleared(generation.FieldPrompt) {
		fields = append(fields, generation.FieldPrompt)
	}
	if m.FieldCleared(generation.FieldVideoURL) {
		fields = append(fields, generation.FieldVideoURL)
	}
	if m.FieldCleared(generation.FieldScriptURL) {
		fields = append(fields, generation.FieldScriptURL)
	}
	if m.FieldCleared(generation.FieldErrorMessage) {
		fields = append(fields, generation.FieldErrorMessage)
	}
	if m.FieldCleared(generation.FieldEmail) {
		fields = append(fields, generation.FieldEmail)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *GenerationMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *GenerationMutation) ClearField(name string) error {
	switch name {
	case generation.FieldRequestID:
		m.ClearRequestID()
		return nil
	case generation.FieldPrompt:
		m.ClearPrompt()
		return nil
	case generation.FieldVideoURL:
		m.ClearVideoURL()
		return nil
	case generation.FieldScriptURL:
		m.ClearScriptURL()
		return nil
	case generation.FieldErrorMessage:
		m.ClearErrorMessage()
		return nil
	case generation.FieldEmail:
		m.ClearEmail()
		return nil
	}
	return fmt.Errorf("unknown Generation nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *GenerationMutation) ResetField(name string) error {
	switch name {
	case generation.FieldRequestID:
		m.ResetRequestID()
		return nil
	case generation.FieldPrompt:
		m.ResetPrompt()
		return nil
	case generation.FieldStatus:
		m.ResetStatus()
		return nil
	case generation.FieldVideoURL:
		m.ResetVideoURL()
		return nil
	case generation.FieldScriptURL:
		m.ResetScriptURL()
		return nil
	case generation.FieldErrorMessage:
		m.ResetErrorMessage()
		return nil
	case generation.FieldEmail:
		m.ResetEmail()
		return nil
	case generation.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case generation.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Generation field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *GenerationMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *GenerationMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *GenerationMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *GenerationMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *GenerationMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *GenerationMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *GenerationMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Generation unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *GenerationMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Generation edge %s", name)
}
