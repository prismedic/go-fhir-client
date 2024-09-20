package fhir_test

import (
	"context"
	"testing"

	fhir "github.com/prismedic/go-fhir-client"
	fhirmodel "github.com/samply/golang-fhir-models/fhir-models/fhir"
)

func TestSearch(t *testing.T) {
	client := fhir.NewClient(&fhir.Config{
		BaseURL: "https://hapi.fhir.org/baseR4",
	})
	ctx := context.Background()
	t.Run("Search Patient", func(t *testing.T) {
		params := fhir.SearchParams{}
		params.SetId("596423")
		bundle, resources, err := fhir.Search[fhirmodel.Patient](ctx, client, "Patient", params)
		if err != nil {
			t.Fatalf("Failed to search patients: %v", err)
		}
		if len(resources) != 1 {
			t.Fatalf("Expected 1 resource, got %d", len(resources))
		}
		if len(bundle.Entry) != 1 {
			t.Fatalf("Expected 1 entry, got %d", len(bundle.Entry))
		}
		if *resources[0].Id != "596423" {
			t.Errorf("Expected resource ID to be 596423, got %s", *resources[0].Id)
		}
		if len(resources[0].Name) != 1 {
			t.Fatalf("Expected 1 name, got %d", len(resources[0].Name))
		}
		if len(resources[0].Name[0].Given) != 1 {
			t.Fatalf("Expected 1 given name, got %d", len(resources[0].Name[0].Given))
		}
		if resources[0].Name[0].Given[0] != "Caleb" {
			t.Errorf("Expected given name to be Caleb, got %s", resources[0].Name[0].Given[0])
		}
		if *resources[0].Name[0].Family != "Cushing" {
			t.Errorf("Expected family name to be Cushing, got %s", *resources[0].Name[0].Family)
		}
	})
}
