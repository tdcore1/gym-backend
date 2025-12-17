package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Wallet holds the schema definition for the Wallet entity.
type Wallet struct {
	ent.Schema
}

// Fields of the Wallet.
func (Wallet) Fields() []ent.Field {
	return []ent.Field{
		field.Int("amount"),
		field.Int("userid"),
	}
}

// Edges of the Wallet.
func (Wallet) Edges() []ent.Edge {
	return nil
}
