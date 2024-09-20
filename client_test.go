package fhir_test

import (
	"context"
	"errors"
	"testing"

	fhir "github.com/prismedic/go-fhir-client"
	fhirmodel "github.com/samply/golang-fhir-models/fhir-models/fhir"
)

func TestCreate(t *testing.T) {
	client := fhir.NewClient(&fhir.Config{
		BaseURL: "https://hapi.fhir.org/baseR4",
	})
	ctx := context.Background()
	t.Run("Create Patient", func(t *testing.T) {
		patient, err := fhir.Create(ctx, client, "Patient", &fhirmodel.Patient{
			Name: []fhirmodel.HumanName{{
				Given:  []string{"John"},
				Family: fhir.String("Li"),
			}},
		})
		if err != nil {
			t.Fatalf("Failed to create patient: %v", err)
		}
		if len(patient.Name) != 1 {
			t.Fatalf("Expected 1 name, got %d", len(patient.Name))
		}
		if len(patient.Name[0].Given) != 1 {
			t.Fatalf("Expected 1 given name, got %d", len(patient.Name[0].Given))
		}
		if patient.Name[0].Given[0] != "John" {
			t.Errorf("Expected given name to be John, got %s", patient.Name[0].Given[0])
		}
		if *patient.Name[0].Family != "Li" {
			t.Errorf("Expected family name to be Li, got %s", *patient.Name[0].Family)
		}
	})
}

func TestRead(t *testing.T) {
	client := fhir.NewClient(&fhir.Config{
		BaseURL: "https://hapi.fhir.org/baseR4",
	})
	ctx := context.Background()
	t.Run("Read Patient", func(t *testing.T) {
		patient, err := fhir.Read[fhirmodel.Patient](ctx, client, "Patient", "596423")
		if err != nil {
			t.Fatalf("Failed to read patient: %v", err)
		}
		if *patient.Id != "596423" {
			t.Errorf("Expected patient ID to be 596423, got %s", *patient.Id)
		}
		if len(patient.Name) != 1 {
			t.Fatalf("Expected 1 name, got %d", len(patient.Name))
		}
		if len(patient.Name[0].Given) != 1 {
			t.Fatalf("Expected 1 given name, got %d", len(patient.Name[0].Given))
		}
		if patient.Name[0].Given[0] != "Caleb" {
			t.Errorf("Expected given name to be Caleb, got %s", patient.Name[0].Given[0])
		}
		if *patient.Name[0].Family != "Cushing" {
			t.Errorf("Expected family name to be Cushing, got %s", *patient.Name[0].Family)
		}
	})
	t.Run("Read Patient that doesn't exist", func(t *testing.T) {
		patient, err := fhir.Read[fhirmodel.Patient](ctx, client, "Patient", "999999")
		if err == nil {
			t.Fatalf("Expected error, got %v", patient)
		}
		var notFoundErr *fhir.NotFoundError
		if !errors.As(err, &notFoundErr) {
			t.Fatalf("Expected NotFoundError, got %v", err)
		}
		if len(notFoundErr.Issue) < 1 {
			t.Fatalf("Expected issues, got %d", len(notFoundErr.Issue))
		}
		if notFoundErr.Issue[0].Severity != fhirmodel.IssueSeverityError {
			t.Fatalf("Expected severity to be error, got %s", notFoundErr.Issue[0].Severity)
		}
	})
}

func TestUpdate(t *testing.T) {
	client := fhir.NewClient(&fhir.Config{
		BaseURL: "https://hapi.fhir.org/baseR4",
	})
	ctx := context.Background()
	t.Run("Update Patient", func(t *testing.T) {
		patient, err := fhir.Update(ctx, client, "Patient", "596422", &fhirmodel.Patient{
			Id: fhir.String("596422"),
			Name: []fhirmodel.HumanName{{
				Given:  []string{"John"},
				Family: fhir.String("Li"),
			}},
		})
		if err != nil {
			t.Fatalf("Failed to update patient: %v", err)
		}
		if len(patient.Name) != 1 {
			t.Fatalf("Expected 1 name, got %d", len(patient.Name))
		}
		if len(patient.Name[0].Given) != 1 {
			t.Fatalf("Expected 1 given name, got %d", len(patient.Name[0].Given))
		}
		if patient.Name[0].Given[0] != "John" {
			t.Errorf("Expected given name to be John, got %s", patient.Name[0].Given[0])
		}
		if *patient.Name[0].Family != "Li" {
			t.Errorf("Expected family name to be Li, got %s", *patient.Name[0].Family)
		}
	})
}
