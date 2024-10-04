package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique(),
		field.String("password"),

		field.String("name").Optional(),

		field.Time("created").Default(time.Now).Optional(),
	}
}

func (User) Edges() []ent.Edge {
	return nil
}
