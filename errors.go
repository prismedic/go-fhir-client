package fhir

import (
	"fmt"

	"github.com/samply/golang-fhir-models/fhir-models/fhir"
)

type FHIRError fhir.OperationOutcome

func (e *FHIRError) Error() string {
	if len(e.Issue) > 0 {
		if e.Issue[0].Diagnostics != nil {
			return *e.Issue[0].Diagnostics
		}
		return "unknown error"
	}
	return "unknown error"
}

type BadRequestError FHIRError

func (e *BadRequestError) Error() string {
	return fmt.Sprintf("bad request: %s", (*FHIRError)(e).Error())
}

func (e *BadRequestError) Unwrap() error {
	return (*FHIRError)(e)
}

type NotFoundError FHIRError

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("not found: %s", (*FHIRError)(e).Error())
}

func (e *NotFoundError) Unwrap() error {
	return (*FHIRError)(e)
}

type UnprocessableEntityError FHIRError

func (e *UnprocessableEntityError) Error() string {
	return fmt.Sprintf("unprocessable entity: %s", (*FHIRError)(e).Error())
}

func (e *UnprocessableEntityError) Unwrap() error {
	return (*FHIRError)(e)
}

type GoneError FHIRError

func (e *GoneError) Error() string {
	return fmt.Sprintf("gone: %s", (*FHIRError)(e).Error())
}

func (e *GoneError) Unwrap() error {
	return (*FHIRError)(e)
}

type NotAuthorizedError FHIRError

func (e *NotAuthorizedError) Error() string {
	return fmt.Sprintf("not authorized: %s", (*FHIRError)(e).Error())
}

func (e *NotAuthorizedError) Unwrap() error {
	return (*FHIRError)(e)
}

type MethodNotAllowedError FHIRError

func (e *MethodNotAllowedError) Error() string {
	return fmt.Sprintf("method not allowed: %s", (*FHIRError)(e).Error())
}

func (e *MethodNotAllowedError) Unwrap() error {
	return (*FHIRError)(e)
}
