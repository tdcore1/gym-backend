package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Course holds the schema definition for the Course entity.
type Course struct {
	ent.Schema
}

// Fields of the Course.
func (Course) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("price"),
		field.Int("day"),
	}
}

// Edges of the Course.
func (Course) Edges() []ent.Edge {
	return nil
}
