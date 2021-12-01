package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"github.com/sujithshajee/dnsbl/app/ent/mixin"
)

type AppQuery struct {
	ent.Schema
}

func (AppQuery) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("responses", AppResponse.Type),
		edge.From("ipaddress", IP.Type).
			Ref("queries").
			Annotations(entgql.Bind()).
			Unique().
			Required(),
	}
}

func (AppQuery) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.UUIDMixin{},
		mixin.TimeMixin{},
	}
}
