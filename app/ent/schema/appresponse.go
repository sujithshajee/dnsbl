package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/sujithshajee/dnsbl/app/ent/mixin"
)

type AppResponse struct {
	ent.Schema
}

func (AppResponse) Fields() []ent.Field {
	return []ent.Field{
		field.String("responsecode"),
		field.String("codedescription"),
	}
}

func (AppResponse) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("query", AppQuery.Type).
			Ref("responses").
			Annotations(entgql.Bind()).
			Unique().
			Required(),
	}
}

func (AppResponse) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.UUIDMixin{},
		mixin.TimeMixin{},
	}
}
