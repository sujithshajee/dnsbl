package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sujithshajee/dnsbl/app/ent"
	"github.com/sujithshajee/dnsbl/app/ent/appquery"
	"github.com/sujithshajee/dnsbl/app/graphql/graph/generated"
)

func (r *appQueryResolver) IP(ctx context.Context, obj *ent.AppQuery) (*ent.IP, error) {
	return obj.Edges.IpaddressOrErr()
}

func (r *appQueryResolver) Responses(ctx context.Context, obj *ent.AppQuery, after *ent.Cursor, before *ent.Cursor, first *int, last *int) (*ent.AppResponseConnection, error) {
	return obj.QueryResponses().Paginate(ctx, after, first, before, last)
}

func (r *appResponseResolver) Query(ctx context.Context, obj *ent.AppResponse) (*ent.AppQuery, error) {
	return obj.Edges.QueryOrErr()
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

func (r *iPResolver) Queries(ctx context.Context, obj *ent.IP, after *ent.Cursor, before *ent.Cursor, first *int, last *int, orderBy *ent.AppQueryOrder) (*ent.AppQueryConnection, error) {
	return obj.QueryQueries().
		Paginate(ctx, after, first, before, last,
			ent.WithAppQueryOrder(orderBy),
		)
}

// DNSBLQuery returns generated.DNSBLQueryResolver implementation.
func (r *Resolver) AppQuery() generated.AppQueryResolver { return &appQueryResolver{r} }

// DNSBLResponse returns generated.DNSBLResponseResolver implementation.
func (r *Resolver) AppResponse() generated.AppResponseResolver { return &appResponseResolver{r} }

// IP returns gen.IPResolver implementation.
func (r *Resolver) IP() generated.IPResolver { return &iPResolver{r} }

type appQueryResolver struct{ *Resolver }
type appResponseResolver struct{ *Resolver }
type iPResolver struct{ *Resolver }
