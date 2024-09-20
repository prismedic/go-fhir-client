package fhir

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type Config struct {
	BaseURL string
	Client  *http.Client
}

type Client struct {
	httpClient *resty.Client
}

func NewClient(config *Config) *Client {
	if config.Client == nil {
		config.Client = http.DefaultClient
	}
	return &Client{
		httpClient: resty.
			NewWithClient(config.Client).
			SetBaseURL(config.BaseURL),
	}
}

// implementation of https://hl7.org/fhir/R4/http.html#create
func Create[T any](ctx context.Context, client *Client, resourceType string, data *T) (*T, error) {
	var result T
	var fhirError FHIRError
	resp, err := client.httpClient.R().
		SetContext(ctx).
		SetBody(data).
		SetResult(&result).
		SetError(&fhirError).
		Post(fmt.Sprintf("/%s", resourceType))
	if err != nil {
		return nil, fmt.Errorf("failed to send request for creating %s: %w", resourceType, err)
	}
	if resp.IsError() {
		switch resp.StatusCode() {
		case http.StatusBadRequest:
			return nil, (*BadRequestError)(&fhirError)
		case http.StatusUnprocessableEntity:
			return nil, (*UnprocessableEntityError)(&fhirError)
		case http.StatusNotFound:
			return nil, (*NotFoundError)(&fhirError)
		default:
			return nil, fmt.Errorf("HTTP error %s: %w", resp.Status(), &fhirError)
		}
	}
	return &result, nil
}

// implementation of https://hl7.org/fhir/R4/http.html#read
func Read[T any](ctx context.Context, client *Client, resourceType string, id string) (*T, error) {
	var result T
	var fhirError FHIRError
	resp, err := client.httpClient.R().
		SetContext(ctx).
		SetResult(&result).
		SetError(&fhirError).
		Get(fmt.Sprintf("/%s/%s", resourceType, id))
	if err != nil {
		return nil, fmt.Errorf("failed to send request for reading %s: %w", resourceType, err)
	}
	if resp.IsError() {
		switch resp.StatusCode() {
		case http.StatusNotFound:
			return nil, (*NotFoundError)(&fhirError)
		case http.StatusGone:
			return nil, (*GoneError)(&fhirError)
		default:
			return nil, fmt.Errorf("HTTP error %s: %w", resp.Status(), &fhirError)
		}
	}
	return &result, nil
}

// implementation of https://hl7.org/fhir/R4/http.html#update
func Update[T any](ctx context.Context, client *Client, resourceType string, id string, data *T) (*T, error) {
	var result T
	var fhirError FHIRError
	resp, err := client.httpClient.R().
		SetContext(ctx).
		SetBody(data).
		SetResult(&result).
		SetError(&fhirError).
		Put(fmt.Sprintf("/%s/%s", resourceType, id))
	if err != nil {
		return nil, fmt.Errorf("failed to send request for updating %s: %w", resourceType, err)
	}
	if resp.IsError() {
		switch resp.StatusCode() {
		case http.StatusBadRequest:
			return nil, (*BadRequestError)(&fhirError)
		case http.StatusUnauthorized:
			return nil, (*NotAuthorizedError)(&fhirError)
		case http.StatusUnprocessableEntity:
			return nil, (*UnprocessableEntityError)(&fhirError)
		case http.StatusNotFound:
			return nil, (*NotFoundError)(&fhirError)
		default:
			return nil, fmt.Errorf("HTTP error %s: %w", resp.Status(), &fhirError)
		}
	}
	return &result, nil
}

func String(s string) *string {
	return &s
}
