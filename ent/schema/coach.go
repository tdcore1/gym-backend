package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Coach holds the schema definition for the Coach entity.
type Coach struct {
	ent.Schema
}

// Fields of the Coach.
func (Coach) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("price"),
	}
}

// Edges of the Coach.
func (Coach) Edges() []ent.Edge {
	return nil
}
