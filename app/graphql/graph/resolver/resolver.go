package resolver

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/99designs/gqlgen/graphql"

	"github.com/sujithshajee/dnsbl/app/ent"
	"github.com/sujithshajee/dnsbl/app/graphql/graph/generated"
)

// Resolver is the root resolver
type Resolver struct {
	client *ent.Client
}

// NewSchema generates a new executable schema for graphql
func NewSchema(c *ent.Client) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{client: c},
	})
}
