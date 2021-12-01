package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sujithshajee/dnsbl/app/server/graph/generated"
	"github.com/sujithshajee/dnsbl/app/server/graph/model"
)

func (r *mutationResolver) Enqueue(ctx context.Context, ip []string) ([]*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Node(ctx context.Context, id uuid.UUID) (model.Node, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetIPDetails(ctx context.Context, ip string) (*model.IP, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
