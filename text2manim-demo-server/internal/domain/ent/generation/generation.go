// Code generated by ent, DO NOT EDIT.

package generation

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the generation type in the database.
	Label = "generation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRequestID holds the string denoting the request_id field in the database.
	FieldRequestID = "request_id"
	// FieldPrompt holds the string denoting the prompt field in the database.
	FieldPrompt = "prompt"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldVideoURL holds the string denoting the video_url field in the database.
	FieldVideoURL = "video_url"
	// FieldScriptURL holds the string denoting the script_url field in the database.
	FieldScriptURL = "script_url"
	// FieldErrorMessage holds the string denoting the error_message field in the database.
	FieldErrorMessage = "error_message"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the generation in the database.
	Table = "generations"
)

// Columns holds all SQL columns for generation fields.
var Columns = []string{
	FieldID,
	FieldRequestID,
	FieldPrompt,
	FieldStatus,
	FieldVideoURL,
	FieldScriptURL,
	FieldErrorMessage,
	FieldEmail,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the "status" enum field.
type Status string

// StatusUnspecified is the default value of the Status enum.
const DefaultStatus = StatusUnspecified

// Status values.
const (
	StatusUnspecified Status = "unspecified"
	StatusPending     Status = "pending"
	StatusProcessing  Status = "processing"
	StatusCompleted   Status = "completed"
	StatusFailed      Status = "failed"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusUnspecified, StatusPending, StatusProcessing, StatusCompleted, StatusFailed:
		return nil
	default:
		return fmt.Errorf("generation: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Generation queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByRequestID orders the results by the request_id field.
func ByRequestID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRequestID, opts...).ToFunc()
}

// ByPrompt orders the results by the prompt field.
func ByPrompt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrompt, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByVideoURL orders the results by the video_url field.
func ByVideoURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVideoURL, opts...).ToFunc()
}

// ByScriptURL orders the results by the script_url field.
func ByScriptURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScriptURL, opts...).ToFunc()
}

// ByErrorMessage orders the results by the error_message field.
func ByErrorMessage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldErrorMessage, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}
