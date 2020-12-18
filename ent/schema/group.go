package schema

import "github.com/facebook/ent"

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

func (Group) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ConfirmMixin{},
	}
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return nil
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return nil
}
