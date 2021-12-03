package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sujithshajee/dnsbl/app/ent"
	"github.com/sujithshajee/dnsbl/app/ent/appquery"
	"github.com/sujithshajee/dnsbl/app/server/graph/generated"
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
	rsp, err := obj.QueryQueries().
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

func (r *iPResolver) Queries(ctx context.Context, obj *ent.IP) (*ent.AppQueryConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// AppQuery returns generated.AppQueryResolver implementation.
func (r *Resolver) AppQuery() generated.AppQueryResolver { return &appQueryResolver{r} }

// AppResponse returns generated.AppResponseResolver implementation.
func (r *Resolver) AppResponse() generated.AppResponseResolver { return &appResponseResolver{r} }

// IP returns generated.IPResolver implementation.
func (r *Resolver) IP() generated.IPResolver { return &iPResolver{r} }

type appQueryResolver struct{ *Resolver }
type appResponseResolver struct{ *Resolver }
type iPResolver struct{ *Resolver }
