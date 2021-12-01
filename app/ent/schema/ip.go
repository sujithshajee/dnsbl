package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/sujithshajee/dnsbl/app/ent/mixin"
)

type IP struct {
	ent.Schema
}

func (IP) Fields() []ent.Field {
	return []ent.Field{
		field.String("ip_address"),
	}
}

func (IP) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("queries", AppQuery.Type).
			Annotations(entgql.Bind()),
	}
}

func (IP) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.UUIDMixin{},
		mixin.TimeMixin{},
	}
}
