package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/sujithshajee/dnsbl/app/ent/mixin"
)

type Task struct {
	ent.Schema
}

func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			Values("IPDNSBL"),
		field.String("ipaddress").Optional(),
		field.Enum("status").
			Values("WAITING", "IN_PROGRESS", "DONE", "ERROR"),
		field.String("error").Optional(),
		field.Time("completed_by").Optional(),
	}
}

func (Task) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.UUIDMixin{},
		mixin.TimeMixin{},
	}
}
