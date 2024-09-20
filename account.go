package fhir

import (
	"context"

	"github.com/samply/golang-fhir-models/fhir-models/fhir"
)

func (c *Client) CreateAccount(ctx context.Context, data *fhir.Account) (*fhir.Account, error) {
	return Create(ctx, c, "Account", data)
}

func (c *Client) ReadAccount(ctx context.Context, id string) (*fhir.Account, error) {
	return Read[fhir.Account](ctx, c, "Account", id)
}

func (c *Client) UpdateAccount(ctx context.Context, id string, data *fhir.Account) (*fhir.Account, error) {
	return Update(ctx, c, "Account", id, data)
}

func (c *Client) SearchAccount(ctx context.Context, params SearchParams) (*fhir.Bundle, []*fhir.Account, error) {
	return Search[fhir.Account](ctx, c, "Account", params)
}
