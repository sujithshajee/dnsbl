package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"net"

	"github.com/google/uuid"
	"github.com/sujithshajee/dnsbl/app/ent"
	ipaddr "github.com/sujithshajee/dnsbl/app/ent/ip"
	"github.com/sujithshajee/dnsbl/app/ent/task"
	"github.com/sujithshajee/dnsbl/app/graphql/graph/generated"
)

func (r *mutationResolver) Enqueue(ctx context.Context, ip []string) ([]*ent.Task, error) {
	tsk := []*ent.Task{}
	client := ent.FromContext(ctx)

	for _, i := range ip {
		if pip := net.ParseIP(i); pip == nil {
			return nil, fmt.Errorf("invalid IP address: %s", i)
		}

		op, err := client.Task.Create().
			SetIpaddress(i).
			SetType(task.TypeIPDNSBL).
			SetStatus(task.StatusWAITING).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("enqueing operation for '%s': %w", i, err)
		}

		tsk = append(tsk, op)
	}

	return tsk, nil
}

func (r *queryResolver) Node(ctx context.Context, id uuid.UUID) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) GetIPDetails(ctx context.Context, ip string) (*ent.IP, error) {
	return r.client.IP.Query().Where(ipaddr.IPAddressEQ(ip)).Only(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
