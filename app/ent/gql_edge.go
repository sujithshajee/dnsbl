// Code generated by entc, DO NOT EDIT.

package ent

import "context"

func (aq *AppQuery) Responses(ctx context.Context) ([]*AppResponse, error) {
	result, err := aq.Edges.ResponsesOrErr()
	if IsNotLoaded(err) {
		result, err = aq.QueryResponses().All(ctx)
	}
	return result, err
}

func (aq *AppQuery) Ipaddress(ctx context.Context) (*IP, error) {
	result, err := aq.Edges.IpaddressOrErr()
	if IsNotLoaded(err) {
		result, err = aq.QueryIpaddress().Only(ctx)
	}
	return result, err
}

func (ar *AppResponse) Query(ctx context.Context) (*AppQuery, error) {
	result, err := ar.Edges.QueryOrErr()
	if IsNotLoaded(err) {
		result, err = ar.QueryQuery().Only(ctx)
	}
	return result, err
}

func (i *IP) Queries(ctx context.Context) ([]*AppQuery, error) {
	result, err := i.Edges.QueriesOrErr()
	if IsNotLoaded(err) {
		result, err = i.QueryQueries().All(ctx)
	}
	return result, err
}
