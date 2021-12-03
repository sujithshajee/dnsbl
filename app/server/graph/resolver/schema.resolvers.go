package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"net"

	"github.com/google/uuid"
	"github.com/sujithshajee/dnsbl/app/ent"
	ipadder "github.com/sujithshajee/dnsbl/app/ent/ip"
	"github.com/sujithshajee/dnsbl/app/ent/task"
	"github.com/sujithshajee/dnsbl/app/server/graph/generated"
	"github.com/sujithshajee/dnsbl/app/server/graph/model"
)

func (r *mutationResolver) Enqueue(ctx context.Context, ip []string) ([]*ent.Task, error) {
	ops := []*ent.Task{}
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

		ops = append(ops, op)
	}

	return ops, nil
}

func (r *queryResolver) Node(ctx context.Context, id uuid.UUID) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) GetIPDetails(ctx context.Context, ip string) (*ent.IP, error) {
	return r.client.IP.Query().Where(ipadder.IPAddressEQ(ip)).Only(ctx)
}

func (r *taskResolver) Type(ctx context.Context, obj *ent.Task) (model.TaskType, error) {
	return model.TaskType(obj.Type.String()), nil
}

func (r *taskResolver) Status(ctx context.Context, obj *ent.Task) (model.TaskStatus, error) {
	return model.TaskStatus(obj.Type.String()), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Task returns generated.TaskResolver implementation.
func (r *Resolver) Task() generated.TaskResolver { return &taskResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type taskResolver struct{ *Resolver }
