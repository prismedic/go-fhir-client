package fhir

import (
	"context"
	fhirmodel "github.com/samply/golang-fhir-models/fhir-models/fhir"
)

func (c *Client) CreateAccount(ctx context.Context, data *fhirmodel.Account) (*fhirmodel.Account, error) {
	return Create(ctx, c, "Account", data)
}

func (c *Client) ReadAccount(ctx context.Context, id string) (*fhirmodel.Account, error) {
	return Read[fhirmodel.Account](ctx, c, "Account", id)
}

func (c *Client) UpdateAccount(ctx context.Context, id string, data *fhirmodel.Account) (*fhirmodel.Account, error) {
	return Update(ctx, c, "Account", id, data)
}

func (c *Client) SearchAccount(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Account, error) {
	return Search[fhirmodel.Account](ctx, c, "Account", params)
}
