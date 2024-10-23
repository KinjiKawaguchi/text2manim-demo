// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// GenerationsColumns holds the columns for the "generations" table.
	GenerationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "request_id", Type: field.TypeString, Unique: true},
		{Name: "prompt", Type: field.TypeString},
		{Name: "status", Type: field.TypeString},
		{Name: "video_url", Type: field.TypeString},
		{Name: "script_url", Type: field.TypeString},
		{Name: "error_message", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// GenerationsTable holds the schema information for the "generations" table.
	GenerationsTable = &schema.Table{
		Name:       "generations",
		Columns:    GenerationsColumns,
		PrimaryKey: []*schema.Column{GenerationsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GenerationsTable,
	}
)

func init() {
}