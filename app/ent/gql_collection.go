// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (aq *AppQueryQuery) CollectFields(ctx context.Context, satisfies ...string) *AppQueryQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		aq = aq.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return aq
}

func (aq *AppQueryQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *AppQueryQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "ipaddress":
			aq = aq.WithIpaddress(func(query *IPQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return aq
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ar *AppResponseQuery) CollectFields(ctx context.Context, satisfies ...string) *AppResponseQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		ar = ar.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return ar
}

func (ar *AppResponseQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *AppResponseQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "query":
			ar = ar.WithQuery(func(query *AppQueryQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return ar
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (i *IPQuery) CollectFields(ctx context.Context, satisfies ...string) *IPQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		i = i.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return i
}

func (i *IPQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *IPQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "queries":
			i = i.WithQueries(func(query *AppQueryQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return i
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TaskQuery) CollectFields(ctx context.Context, satisfies ...string) *TaskQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		t = t.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return t
}

func (t *TaskQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *TaskQuery {
	return t
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) *UserQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		u = u.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return u
}

func (u *UserQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *UserQuery {
	return u
}