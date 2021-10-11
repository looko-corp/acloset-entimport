// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{field.Int("id"), field.Int8("age"), field.String("name")}
}
func (User) Edges() []ent.Edge {
	return nil
}
func (User) Annotations() []schema.Annotation {
	return nil
}
