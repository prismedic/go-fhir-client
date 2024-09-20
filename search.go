package fhir

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/samply/golang-fhir-models/fhir-models/fhir"
)

type ParamPrefix string

const (
	PrefixNone               ParamPrefix = ""
	PrefixEqual              ParamPrefix = "eq"
	PrefixNotEqual           ParamPrefix = "ne"
	PrefixGreaterThan        ParamPrefix = "gt"
	PrefixGreaterThanOrEqual ParamPrefix = "ge"
	PrefixLessThan           ParamPrefix = "lt"
	PrefixLessThanOrEqual    ParamPrefix = "le"
	PrefixStartsAfter        ParamPrefix = "sa"
	PrefixEndsBefore         ParamPrefix = "eb"
	PrefixApproximate        ParamPrefix = "ap"
)

type ParamModifier string

const (
	ModifierNone     ParamModifier = ""
	ModifierExact    ParamModifier = "exact"
	ModifierContains ParamModifier = "contains"
	ModifierBelow    ParamModifier = "below"
	ModifierAbove    ParamModifier = "above"
)

type SearchParams url.Values

func (p SearchParams) SetParam(key string, modifier ParamModifier, prefix ParamPrefix, value string) {
	if modifier != ModifierNone {
		url.Values(p).Set(fmt.Sprintf("%s:%s", key, string(modifier)), fmt.Sprintf("%s%s", string(prefix), value))
	} else {
		url.Values(p).Set(key, fmt.Sprintf("%s%s", string(prefix), value))
	}
}

func (p SearchParams) SetId(id string) {
	p.SetParam("_id", ModifierNone, PrefixNone, id)
}

func (p SearchParams) SetLastUpdated(prefix ParamPrefix, lastUpdated time.Time) {
	p.SetParam("_lastUpdated", ModifierNone, prefix, lastUpdated.Format(time.RFC3339))
}

func (p SearchParams) SetTag(tag string) {
	p.SetParam("_tag", ModifierNone, PrefixNone, tag)
}

func (p SearchParams) SetProfile(modifier ParamModifier, uri url.URL) {
	p.SetParam("_profile", modifier, PrefixNone, uri.String())
}

func (p SearchParams) SetSecurity(security string) {
	p.SetParam("_security", ModifierNone, PrefixNone, security)
}

func (p SearchParams) SetText(modifier ParamModifier, text string) {
	p.SetParam("_text", modifier, PrefixNone, text)
}

func (p SearchParams) SetContent(modifier ParamModifier, content string) {
	p.SetParam("_content", modifier, PrefixNone, content)
}

func (p SearchParams) SetList(modifier ParamModifier, list string) {
	p.SetParam("_list", modifier, PrefixNone, list)
}

func (p SearchParams) SetSort(sort string) {
	p.SetParam("_sort", ModifierNone, PrefixNone, sort)
}

func (p SearchParams) SetCount(count int) {
	p.SetParam("_count", ModifierNone, PrefixNone, fmt.Sprintf("%d", count))
}

func (p SearchParams) SetSummary(summary bool) {
	p.SetParam("_summary", ModifierNone, PrefixNone, fmt.Sprintf("%t", summary))
}

type ContainedParam string

const (
	ContainedTrue  ContainedParam = "true"
	ContainedFalse ContainedParam = "false"
	ContainedBoth  ContainedParam = "both"
)

func (p SearchParams) SetContained(contained ContainedParam) {
	p.SetParam("_contained", ModifierNone, PrefixNone, string(contained))
}

type ContainedTypeParam string

const (
	ContainedTypeContainer ContainedTypeParam = "container"
	ContainedTypeContained ContainedTypeParam = "contained"
)

func (p SearchParams) SetContainedType(containedType ContainedTypeParam) {
	p.SetParam("_containedType", ModifierNone, PrefixNone, string(containedType))
}

func Search[T any](ctx context.Context, client *Client, resourceType string, params SearchParams) (*fhir.Bundle, []*T, error) {
	var result fhir.Bundle
	var fhirError FHIRError
	resp, err := client.httpClient.R().
		SetContext(ctx).
		SetQueryParamsFromValues(url.Values(params)).
		SetResult(&result).
		SetError(&fhirError).
		Get(fmt.Sprintf("/%s", resourceType))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to send request for searching %s: %w", resourceType, err)
	}
	if resp.IsError() {
		switch resp.StatusCode() {
		case http.StatusBadRequest:
			return nil, nil, (*BadRequestError)(&fhirError)
		case http.StatusUnauthorized:
			return nil, nil, (*NotAuthorizedError)(&fhirError)
		case http.StatusNotFound:
			return nil, nil, (*NotFoundError)(&fhirError)
		default:
			return nil, nil, fmt.Errorf("HTTP error %s: %w", resp.Status(), &fhirError)
		}
	}
	resources := make([]*T, 0, len(result.Entry))
	for _, entry := range result.Entry {
		var resource T
		if err := json.Unmarshal(entry.Resource, &resource); err != nil {
			return nil, nil, fmt.Errorf("failed to unmarshal resource: %w", err)
		}
		resources = append(resources, &resource)
	}
	return &result, resources, nil
}
