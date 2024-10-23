package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Generation struct {
	ent.Schema
}

func (Generation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("request_id").Optional(),
		field.String("prompt").Optional(),
		field.Enum("status").
			Values(
				"unspecified",
				"pending",
				"processing",
				"completed",
				"failed",
			).
			Default("unspecified"),
		field.String("video_url").Optional(),
		field.String("script_url").Optional(),
		field.String("error_message").Optional(),
		field.String("email").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}
