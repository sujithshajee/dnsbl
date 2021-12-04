package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"net"

	"github.com/google/uuid"
	"github.com/sujithshajee/dnsbl/app/ent"
	"github.com/sujithshajee/dnsbl/app/ent/appquery"
	ipadder "github.com/sujithshajee/dnsbl/app/ent/ip"
	"github.com/sujithshajee/dnsbl/app/ent/task"
	"github.com/sujithshajee/dnsbl/app/server/graph/generated"
	"github.com/sujithshajee/dnsbl/app/server/graph/model"
)

func (r *appQueryResolver) IP(ctx context.Context, obj *ent.AppQuery) (*ent.IP, error) {
	return obj.Edges.IpaddressOrErr()
}

func (r *appQueryResolver) Responses(ctx context.Context, obj *ent.AppQuery) (*ent.AppResponseConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *appResponseResolver) Code(ctx context.Context, obj *ent.AppResponse) (string, error) {
	rsp, err := obj.QueryQuery().
		Order(
			ent.Desc(appquery.FieldCreatedAt),
		).
		QueryResponses().
		First(ctx)
	if err != nil {
		return "", fmt.Errorf("unable to find response for IP: %w", err)
	}

	return rsp.Responsecode, nil
}

func (r *appResponseResolver) Description(ctx context.Context, obj *ent.AppResponse) (string, error) {
	rsp, err := obj.QueryQuery().
		Order(
			ent.Desc(appquery.FieldCreatedAt),
		).
		QueryResponses().
		First(ctx)
	if err != nil {
		return "", fmt.Errorf("unable to find response for IP: %w", err)
	}

	return rsp.Codedescription, nil
}

func (r *iPResolver) ResponseCode(ctx context.Context, obj *ent.IP) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *iPResolver) Queries(ctx context.Context, obj *ent.IP) (*ent.AppQueryConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Enqueue(ctx context.Context, ip []string) ([]*ent.Task, error) {
	tsk := []*ent.Task{}
	client := ent.FromContext(ctx)

	for _, i := range ip {
		if pip := net.ParseIP(i); pip == nil {
			return nil, fmt.Errorf("invalid IP address: %s", i)
		}

		tk, err := client.Task.Create().
			SetIpaddress(i).
			SetType(task.TypeIPDNSBL).
			SetStatus(task.StatusWAITING).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("enqueing operation for '%s': %w", i, err)
		}

		tsk = append(tsk, tk)
	}

	return tsk, nil
}

func (r *queryResolver) Node(ctx context.Context, id uuid.UUID) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) GetIPDetails(ctx context.Context, ip string) (*ent.IP, error) {
	return r.client.IP.Query().Where(ipadder.IPAddressEQ(ip)).Only(ctx)
}

func (r *taskResolver) Type(ctx context.Context, obj *ent.Task) (model.TaskType, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *taskResolver) Status(ctx context.Context, obj *ent.Task) (model.TaskStatus, error) {
	panic(fmt.Errorf("not implemented"))
}

// AppQuery returns generated.AppQueryResolver implementation.
func (r *Resolver) AppQuery() generated.AppQueryResolver { return &appQueryResolver{r} }

// AppResponse returns generated.AppResponseResolver implementation.
func (r *Resolver) AppResponse() generated.AppResponseResolver { return &appResponseResolver{r} }

// IP returns generated.IPResolver implementation.
func (r *Resolver) IP() generated.IPResolver { return &iPResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Task returns generated.TaskResolver implementation.
func (r *Resolver) Task() generated.TaskResolver { return &taskResolver{r} }

type appQueryResolver struct{ *Resolver }
type appResponseResolver struct{ *Resolver }
type iPResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type taskResolver struct{ *Resolver }
