package fhir

import (
	"context"
	fhirmodel "github.com/samply/golang-fhir-models/fhir-models/fhir"
)

func (c *Client) ReadAccount(ctx context.Context, id string) (*fhirmodel.Account, error) {
	return Read[fhirmodel.Account](ctx, c, "Account", id)
}

func (c *Client) UpdateAccount(ctx context.Context, id string, data *fhirmodel.Account) (*fhirmodel.Account, error) {
	return Update(ctx, c, "Account", id, data)
}

func (c *Client) CreateAccount(ctx context.Context, data *fhirmodel.Account) (*fhirmodel.Account, error) {
	return Create(ctx, c, "Account", data)
}

func (c *Client) SearchAccount(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Account, error) {
	return Search[fhirmodel.Account](ctx, c, "Account", params)
}

func (c *Client) ReadActivityDefinition(ctx context.Context, id string) (*fhirmodel.ActivityDefinition, error) {
	return Read[fhirmodel.ActivityDefinition](ctx, c, "ActivityDefinition", id)
}

func (c *Client) UpdateActivityDefinition(ctx context.Context, id string, data *fhirmodel.ActivityDefinition) (*fhirmodel.ActivityDefinition, error) {
	return Update(ctx, c, "ActivityDefinition", id, data)
}

func (c *Client) CreateActivityDefinition(ctx context.Context, data *fhirmodel.ActivityDefinition) (*fhirmodel.ActivityDefinition, error) {
	return Create(ctx, c, "ActivityDefinition", data)
}

func (c *Client) SearchActivityDefinition(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.ActivityDefinition, error) {
	return Search[fhirmodel.ActivityDefinition](ctx, c, "ActivityDefinition", params)
}

func (c *Client) ReadAdverseEvent(ctx context.Context, id string) (*fhirmodel.AdverseEvent, error) {
	return Read[fhirmodel.AdverseEvent](ctx, c, "AdverseEvent", id)
}

func (c *Client) UpdateAdverseEvent(ctx context.Context, id string, data *fhirmodel.AdverseEvent) (*fhirmodel.AdverseEvent, error) {
	return Update(ctx, c, "AdverseEvent", id, data)
}

func (c *Client) CreateAdverseEvent(ctx context.Context, data *fhirmodel.AdverseEvent) (*fhirmodel.AdverseEvent, error) {
	return Create(ctx, c, "AdverseEvent", data)
}

func (c *Client) SearchAdverseEvent(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.AdverseEvent, error) {
	return Search[fhirmodel.AdverseEvent](ctx, c, "AdverseEvent", params)
}

func (c *Client) ReadAllergyIntolerance(ctx context.Context, id string) (*fhirmodel.AllergyIntolerance, error) {
	return Read[fhirmodel.AllergyIntolerance](ctx, c, "AllergyIntolerance", id)
}

func (c *Client) UpdateAllergyIntolerance(ctx context.Context, id string, data *fhirmodel.AllergyIntolerance) (*fhirmodel.AllergyIntolerance, error) {
	return Update(ctx, c, "AllergyIntolerance", id, data)
}

func (c *Client) CreateAllergyIntolerance(ctx context.Context, data *fhirmodel.AllergyIntolerance) (*fhirmodel.AllergyIntolerance, error) {
	return Create(ctx, c, "AllergyIntolerance", data)
}

func (c *Client) SearchAllergyIntolerance(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.AllergyIntolerance, error) {
	return Search[fhirmodel.AllergyIntolerance](ctx, c, "AllergyIntolerance", params)
}

func (c *Client) ReadAppointment(ctx context.Context, id string) (*fhirmodel.Appointment, error) {
	return Read[fhirmodel.Appointment](ctx, c, "Appointment", id)
}

func (c *Client) UpdateAppointment(ctx context.Context, id string, data *fhirmodel.Appointment) (*fhirmodel.Appointment, error) {
	return Update(ctx, c, "Appointment", id, data)
}

func (c *Client) CreateAppointment(ctx context.Context, data *fhirmodel.Appointment) (*fhirmodel.Appointment, error) {
	return Create(ctx, c, "Appointment", data)
}

func (c *Client) SearchAppointment(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Appointment, error) {
	return Search[fhirmodel.Appointment](ctx, c, "Appointment", params)
}

func (c *Client) ReadAppointmentResponse(ctx context.Context, id string) (*fhirmodel.AppointmentResponse, error) {
	return Read[fhirmodel.AppointmentResponse](ctx, c, "AppointmentResponse", id)
}

func (c *Client) UpdateAppointmentResponse(ctx context.Context, id string, data *fhirmodel.AppointmentResponse) (*fhirmodel.AppointmentResponse, error) {
	return Update(ctx, c, "AppointmentResponse", id, data)
}

func (c *Client) CreateAppointmentResponse(ctx context.Context, data *fhirmodel.AppointmentResponse) (*fhirmodel.AppointmentResponse, error) {
	return Create(ctx, c, "AppointmentResponse", data)
}

func (c *Client) SearchAppointmentResponse(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.AppointmentResponse, error) {
	return Search[fhirmodel.AppointmentResponse](ctx, c, "AppointmentResponse", params)
}

func (c *Client) ReadAuditEvent(ctx context.Context, id string) (*fhirmodel.AuditEvent, error) {
	return Read[fhirmodel.AuditEvent](ctx, c, "AuditEvent", id)
}

func (c *Client) UpdateAuditEvent(ctx context.Context, id string, data *fhirmodel.AuditEvent) (*fhirmodel.AuditEvent, error) {
	return Update(ctx, c, "AuditEvent", id, data)
}

func (c *Client) CreateAuditEvent(ctx context.Context, data *fhirmodel.AuditEvent) (*fhirmodel.AuditEvent, error) {
	return Create(ctx, c, "AuditEvent", data)
}

func (c *Client) SearchAuditEvent(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.AuditEvent, error) {
	return Search[fhirmodel.AuditEvent](ctx, c, "AuditEvent", params)
}

func (c *Client) ReadBasic(ctx context.Context, id string) (*fhirmodel.Basic, error) {
	return Read[fhirmodel.Basic](ctx, c, "Basic", id)
}

func (c *Client) UpdateBasic(ctx context.Context, id string, data *fhirmodel.Basic) (*fhirmodel.Basic, error) {
	return Update(ctx, c, "Basic", id, data)
}

func (c *Client) CreateBasic(ctx context.Context, data *fhirmodel.Basic) (*fhirmodel.Basic, error) {
	return Create(ctx, c, "Basic", data)
}

func (c *Client) SearchBasic(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Basic, error) {
	return Search[fhirmodel.Basic](ctx, c, "Basic", params)
}

func (c *Client) ReadBinary(ctx context.Context, id string) (*fhirmodel.Binary, error) {
	return Read[fhirmodel.Binary](ctx, c, "Binary", id)
}

func (c *Client) UpdateBinary(ctx context.Context, id string, data *fhirmodel.Binary) (*fhirmodel.Binary, error) {
	return Update(ctx, c, "Binary", id, data)
}

func (c *Client) CreateBinary(ctx context.Context, data *fhirmodel.Binary) (*fhirmodel.Binary, error) {
	return Create(ctx, c, "Binary", data)
}

func (c *Client) SearchBinary(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Binary, error) {
	return Search[fhirmodel.Binary](ctx, c, "Binary", params)
}

func (c *Client) ReadBiologicallyDerivedProduct(ctx context.Context, id string) (*fhirmodel.BiologicallyDerivedProduct, error) {
	return Read[fhirmodel.BiologicallyDerivedProduct](ctx, c, "BiologicallyDerivedProduct", id)
}

func (c *Client) UpdateBiologicallyDerivedProduct(ctx context.Context, id string, data *fhirmodel.BiologicallyDerivedProduct) (*fhirmodel.BiologicallyDerivedProduct, error) {
	return Update(ctx, c, "BiologicallyDerivedProduct", id, data)
}

func (c *Client) CreateBiologicallyDerivedProduct(ctx context.Context, data *fhirmodel.BiologicallyDerivedProduct) (*fhirmodel.BiologicallyDerivedProduct, error) {
	return Create(ctx, c, "BiologicallyDerivedProduct", data)
}

func (c *Client) SearchBiologicallyDerivedProduct(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.BiologicallyDerivedProduct, error) {
	return Search[fhirmodel.BiologicallyDerivedProduct](ctx, c, "BiologicallyDerivedProduct", params)
}

func (c *Client) ReadBodyStructure(ctx context.Context, id string) (*fhirmodel.BodyStructure, error) {
	return Read[fhirmodel.BodyStructure](ctx, c, "BodyStructure", id)
}

func (c *Client) UpdateBodyStructure(ctx context.Context, id string, data *fhirmodel.BodyStructure) (*fhirmodel.BodyStructure, error) {
	return Update(ctx, c, "BodyStructure", id, data)
}

func (c *Client) CreateBodyStructure(ctx context.Context, data *fhirmodel.BodyStructure) (*fhirmodel.BodyStructure, error) {
	return Create(ctx, c, "BodyStructure", data)
}

func (c *Client) SearchBodyStructure(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.BodyStructure, error) {
	return Search[fhirmodel.BodyStructure](ctx, c, "BodyStructure", params)
}

func (c *Client) ReadBundle(ctx context.Context, id string) (*fhirmodel.Bundle, error) {
	return Read[fhirmodel.Bundle](ctx, c, "Bundle", id)
}

func (c *Client) UpdateBundle(ctx context.Context, id string, data *fhirmodel.Bundle) (*fhirmodel.Bundle, error) {
	return Update(ctx, c, "Bundle", id, data)
}

func (c *Client) CreateBundle(ctx context.Context, data *fhirmodel.Bundle) (*fhirmodel.Bundle, error) {
	return Create(ctx, c, "Bundle", data)
}

func (c *Client) SearchBundle(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Bundle, error) {
	return Search[fhirmodel.Bundle](ctx, c, "Bundle", params)
}

func (c *Client) ReadCapabilityStatement(ctx context.Context, id string) (*fhirmodel.CapabilityStatement, error) {
	return Read[fhirmodel.CapabilityStatement](ctx, c, "CapabilityStatement", id)
}

func (c *Client) UpdateCapabilityStatement(ctx context.Context, id string, data *fhirmodel.CapabilityStatement) (*fhirmodel.CapabilityStatement, error) {
	return Update(ctx, c, "CapabilityStatement", id, data)
}

func (c *Client) CreateCapabilityStatement(ctx context.Context, data *fhirmodel.CapabilityStatement) (*fhirmodel.CapabilityStatement, error) {
	return Create(ctx, c, "CapabilityStatement", data)
}

func (c *Client) SearchCapabilityStatement(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.CapabilityStatement, error) {
	return Search[fhirmodel.CapabilityStatement](ctx, c, "CapabilityStatement", params)
}

func (c *Client) ReadCarePlan(ctx context.Context, id string) (*fhirmodel.CarePlan, error) {
	return Read[fhirmodel.CarePlan](ctx, c, "CarePlan", id)
}

func (c *Client) UpdateCarePlan(ctx context.Context, id string, data *fhirmodel.CarePlan) (*fhirmodel.CarePlan, error) {
	return Update(ctx, c, "CarePlan", id, data)
}

func (c *Client) CreateCarePlan(ctx context.Context, data *fhirmodel.CarePlan) (*fhirmodel.CarePlan, error) {
	return Create(ctx, c, "CarePlan", data)
}

func (c *Client) SearchCarePlan(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.CarePlan, error) {
	return Search[fhirmodel.CarePlan](ctx, c, "CarePlan", params)
}

func (c *Client) ReadCareTeam(ctx context.Context, id string) (*fhirmodel.CareTeam, error) {
	return Read[fhirmodel.CareTeam](ctx, c, "CareTeam", id)
}

func (c *Client) UpdateCareTeam(ctx context.Context, id string, data *fhirmodel.CareTeam) (*fhirmodel.CareTeam, error) {
	return Update(ctx, c, "CareTeam", id, data)
}

func (c *Client) CreateCareTeam(ctx context.Context, data *fhirmodel.CareTeam) (*fhirmodel.CareTeam, error) {
	return Create(ctx, c, "CareTeam", data)
}

func (c *Client) SearchCareTeam(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.CareTeam, error) {
	return Search[fhirmodel.CareTeam](ctx, c, "CareTeam", params)
}

func (c *Client) ReadCatalogEntry(ctx context.Context, id string) (*fhirmodel.CatalogEntry, error) {
	return Read[fhirmodel.CatalogEntry](ctx, c, "CatalogEntry", id)
}

func (c *Client) UpdateCatalogEntry(ctx context.Context, id string, data *fhirmodel.CatalogEntry) (*fhirmodel.CatalogEntry, error) {
	return Update(ctx, c, "CatalogEntry", id, data)
}

func (c *Client) CreateCatalogEntry(ctx context.Context, data *fhirmodel.CatalogEntry) (*fhirmodel.CatalogEntry, error) {
	return Create(ctx, c, "CatalogEntry", data)
}

func (c *Client) SearchCatalogEntry(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.CatalogEntry, error) {
	return Search[fhirmodel.CatalogEntry](ctx, c, "CatalogEntry", params)
}

func (c *Client) ReadChargeItem(ctx context.Context, id string) (*fhirmodel.ChargeItem, error) {
	return Read[fhirmodel.ChargeItem](ctx, c, "ChargeItem", id)
}

func (c *Client) UpdateChargeItem(ctx context.Context, id string, data *fhirmodel.ChargeItem) (*fhirmodel.ChargeItem, error) {
	return Update(ctx, c, "ChargeItem", id, data)
}

func (c *Client) CreateChargeItem(ctx context.Context, data *fhirmodel.ChargeItem) (*fhirmodel.ChargeItem, error) {
	return Create(ctx, c, "ChargeItem", data)
}

func (c *Client) SearchChargeItem(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.ChargeItem, error) {
	return Search[fhirmodel.ChargeItem](ctx, c, "ChargeItem", params)
}

func (c *Client) ReadChargeItemDefinition(ctx context.Context, id string) (*fhirmodel.ChargeItemDefinition, error) {
	return Read[fhirmodel.ChargeItemDefinition](ctx, c, "ChargeItemDefinition", id)
}

func (c *Client) UpdateChargeItemDefinition(ctx context.Context, id string, data *fhirmodel.ChargeItemDefinition) (*fhirmodel.ChargeItemDefinition, error) {
	return Update(ctx, c, "ChargeItemDefinition", id, data)
}

func (c *Client) CreateChargeItemDefinition(ctx context.Context, data *fhirmodel.ChargeItemDefinition) (*fhirmodel.ChargeItemDefinition, error) {
	return Create(ctx, c, "ChargeItemDefinition", data)
}

func (c *Client) SearchChargeItemDefinition(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.ChargeItemDefinition, error) {
	return Search[fhirmodel.ChargeItemDefinition](ctx, c, "ChargeItemDefinition", params)
}

func (c *Client) ReadClaim(ctx context.Context, id string) (*fhirmodel.Claim, error) {
	return Read[fhirmodel.Claim](ctx, c, "Claim", id)
}

func (c *Client) UpdateClaim(ctx context.Context, id string, data *fhirmodel.Claim) (*fhirmodel.Claim, error) {
	return Update(ctx, c, "Claim", id, data)
}

func (c *Client) CreateClaim(ctx context.Context, data *fhirmodel.Claim) (*fhirmodel.Claim, error) {
	return Create(ctx, c, "Claim", data)
}

func (c *Client) SearchClaim(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Claim, error) {
	return Search[fhirmodel.Claim](ctx, c, "Claim", params)
}

func (c *Client) ReadClaimResponse(ctx context.Context, id string) (*fhirmodel.ClaimResponse, error) {
	return Read[fhirmodel.ClaimResponse](ctx, c, "ClaimResponse", id)
}

func (c *Client) UpdateClaimResponse(ctx context.Context, id string, data *fhirmodel.ClaimResponse) (*fhirmodel.ClaimResponse, error) {
	return Update(ctx, c, "ClaimResponse", id, data)
}

func (c *Client) CreateClaimResponse(ctx context.Context, data *fhirmodel.ClaimResponse) (*fhirmodel.ClaimResponse, error) {
	return Create(ctx, c, "ClaimResponse", data)
}

func (c *Client) SearchClaimResponse(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.ClaimResponse, error) {
	return Search[fhirmodel.ClaimResponse](ctx, c, "ClaimResponse", params)
}

func (c *Client) ReadClinicalImpression(ctx context.Context, id string) (*fhirmodel.ClinicalImpression, error) {
	return Read[fhirmodel.ClinicalImpression](ctx, c, "ClinicalImpression", id)
}

func (c *Client) UpdateClinicalImpression(ctx context.Context, id string, data *fhirmodel.ClinicalImpression) (*fhirmodel.ClinicalImpression, error) {
	return Update(ctx, c, "ClinicalImpression", id, data)
}

func (c *Client) CreateClinicalImpression(ctx context.Context, data *fhirmodel.ClinicalImpression) (*fhirmodel.ClinicalImpression, error) {
	return Create(ctx, c, "ClinicalImpression", data)
}

func (c *Client) SearchClinicalImpression(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.ClinicalImpression, error) {
	return Search[fhirmodel.ClinicalImpression](ctx, c, "ClinicalImpression", params)
}

func (c *Client) ReadCodeSystem(ctx context.Context, id string) (*fhirmodel.CodeSystem, error) {
	return Read[fhirmodel.CodeSystem](ctx, c, "CodeSystem", id)
}

func (c *Client) UpdateCodeSystem(ctx context.Context, id string, data *fhirmodel.CodeSystem) (*fhirmodel.CodeSystem, error) {
	return Update(ctx, c, "CodeSystem", id, data)
}

func (c *Client) CreateCodeSystem(ctx context.Context, data *fhirmodel.CodeSystem) (*fhirmodel.CodeSystem, error) {
	return Create(ctx, c, "CodeSystem", data)
}

func (c *Client) SearchCodeSystem(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.CodeSystem, error) {
	return Search[fhirmodel.CodeSystem](ctx, c, "CodeSystem", params)
}

func (c *Client) ReadCommunication(ctx context.Context, id string) (*fhirmodel.Communication, error) {
	return Read[fhirmodel.Communication](ctx, c, "Communication", id)
}

func (c *Client) UpdateCommunication(ctx context.Context, id string, data *fhirmodel.Communication) (*fhirmodel.Communication, error) {
	return Update(ctx, c, "Communication", id, data)
}

func (c *Client) CreateCommunication(ctx context.Context, data *fhirmodel.Communication) (*fhirmodel.Communication, error) {
	return Create(ctx, c, "Communication", data)
}

func (c *Client) SearchCommunication(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Communication, error) {
	return Search[fhirmodel.Communication](ctx, c, "Communication", params)
}

func (c *Client) ReadCommunicationRequest(ctx context.Context, id string) (*fhirmodel.CommunicationRequest, error) {
	return Read[fhirmodel.CommunicationRequest](ctx, c, "CommunicationRequest", id)
}

func (c *Client) UpdateCommunicationRequest(ctx context.Context, id string, data *fhirmodel.CommunicationRequest) (*fhirmodel.CommunicationRequest, error) {
	return Update(ctx, c, "CommunicationRequest", id, data)
}

func (c *Client) CreateCommunicationRequest(ctx context.Context, data *fhirmodel.CommunicationRequest) (*fhirmodel.CommunicationRequest, error) {
	return Create(ctx, c, "CommunicationRequest", data)
}

func (c *Client) SearchCommunicationRequest(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.CommunicationRequest, error) {
	return Search[fhirmodel.CommunicationRequest](ctx, c, "CommunicationRequest", params)
}

func (c *Client) ReadCompartmentDefinition(ctx context.Context, id string) (*fhirmodel.CompartmentDefinition, error) {
	return Read[fhirmodel.CompartmentDefinition](ctx, c, "CompartmentDefinition", id)
}

func (c *Client) UpdateCompartmentDefinition(ctx context.Context, id string, data *fhirmodel.CompartmentDefinition) (*fhirmodel.CompartmentDefinition, error) {
	return Update(ctx, c, "CompartmentDefinition", id, data)
}

func (c *Client) CreateCompartmentDefinition(ctx context.Context, data *fhirmodel.CompartmentDefinition) (*fhirmodel.CompartmentDefinition, error) {
	return Create(ctx, c, "CompartmentDefinition", data)
}

func (c *Client) SearchCompartmentDefinition(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.CompartmentDefinition, error) {
	return Search[fhirmodel.CompartmentDefinition](ctx, c, "CompartmentDefinition", params)
}

func (c *Client) ReadComposition(ctx context.Context, id string) (*fhirmodel.Composition, error) {
	return Read[fhirmodel.Composition](ctx, c, "Composition", id)
}

func (c *Client) UpdateComposition(ctx context.Context, id string, data *fhirmodel.Composition) (*fhirmodel.Composition, error) {
	return Update(ctx, c, "Composition", id, data)
}

func (c *Client) CreateComposition(ctx context.Context, data *fhirmodel.Composition) (*fhirmodel.Composition, error) {
	return Create(ctx, c, "Composition", data)
}

func (c *Client) SearchComposition(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Composition, error) {
	return Search[fhirmodel.Composition](ctx, c, "Composition", params)
}

func (c *Client) ReadConceptMap(ctx context.Context, id string) (*fhirmodel.ConceptMap, error) {
	return Read[fhirmodel.ConceptMap](ctx, c, "ConceptMap", id)
}

func (c *Client) UpdateConceptMap(ctx context.Context, id string, data *fhirmodel.ConceptMap) (*fhirmodel.ConceptMap, error) {
	return Update(ctx, c, "ConceptMap", id, data)
}

func (c *Client) CreateConceptMap(ctx context.Context, data *fhirmodel.ConceptMap) (*fhirmodel.ConceptMap, error) {
	return Create(ctx, c, "ConceptMap", data)
}

func (c *Client) SearchConceptMap(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.ConceptMap, error) {
	return Search[fhirmodel.ConceptMap](ctx, c, "ConceptMap", params)
}

func (c *Client) ReadCondition(ctx context.Context, id string) (*fhirmodel.Condition, error) {
	return Read[fhirmodel.Condition](ctx, c, "Condition", id)
}

func (c *Client) UpdateCondition(ctx context.Context, id string, data *fhirmodel.Condition) (*fhirmodel.Condition, error) {
	return Update(ctx, c, "Condition", id, data)
}

func (c *Client) CreateCondition(ctx context.Context, data *fhirmodel.Condition) (*fhirmodel.Condition, error) {
	return Create(ctx, c, "Condition", data)
}

func (c *Client) SearchCondition(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Condition, error) {
	return Search[fhirmodel.Condition](ctx, c, "Condition", params)
}

func (c *Client) ReadConsent(ctx context.Context, id string) (*fhirmodel.Consent, error) {
	return Read[fhirmodel.Consent](ctx, c, "Consent", id)
}

func (c *Client) UpdateConsent(ctx context.Context, id string, data *fhirmodel.Consent) (*fhirmodel.Consent, error) {
	return Update(ctx, c, "Consent", id, data)
}

func (c *Client) CreateConsent(ctx context.Context, data *fhirmodel.Consent) (*fhirmodel.Consent, error) {
	return Create(ctx, c, "Consent", data)
}

func (c *Client) SearchConsent(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Consent, error) {
	return Search[fhirmodel.Consent](ctx, c, "Consent", params)
}

func (c *Client) ReadContract(ctx context.Context, id string) (*fhirmodel.Contract, error) {
	return Read[fhirmodel.Contract](ctx, c, "Contract", id)
}

func (c *Client) UpdateContract(ctx context.Context, id string, data *fhirmodel.Contract) (*fhirmodel.Contract, error) {
	return Update(ctx, c, "Contract", id, data)
}

func (c *Client) CreateContract(ctx context.Context, data *fhirmodel.Contract) (*fhirmodel.Contract, error) {
	return Create(ctx, c, "Contract", data)
}

func (c *Client) SearchContract(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Contract, error) {
	return Search[fhirmodel.Contract](ctx, c, "Contract", params)
}

func (c *Client) ReadCoverage(ctx context.Context, id string) (*fhirmodel.Coverage, error) {
	return Read[fhirmodel.Coverage](ctx, c, "Coverage", id)
}

func (c *Client) UpdateCoverage(ctx context.Context, id string, data *fhirmodel.Coverage) (*fhirmodel.Coverage, error) {
	return Update(ctx, c, "Coverage", id, data)
}

func (c *Client) CreateCoverage(ctx context.Context, data *fhirmodel.Coverage) (*fhirmodel.Coverage, error) {
	return Create(ctx, c, "Coverage", data)
}

func (c *Client) SearchCoverage(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Coverage, error) {
	return Search[fhirmodel.Coverage](ctx, c, "Coverage", params)
}

func (c *Client) ReadCoverageEligibilityRequest(ctx context.Context, id string) (*fhirmodel.CoverageEligibilityRequest, error) {
	return Read[fhirmodel.CoverageEligibilityRequest](ctx, c, "CoverageEligibilityRequest", id)
}

func (c *Client) UpdateCoverageEligibilityRequest(ctx context.Context, id string, data *fhirmodel.CoverageEligibilityRequest) (*fhirmodel.CoverageEligibilityRequest, error) {
	return Update(ctx, c, "CoverageEligibilityRequest", id, data)
}

func (c *Client) CreateCoverageEligibilityRequest(ctx context.Context, data *fhirmodel.CoverageEligibilityRequest) (*fhirmodel.CoverageEligibilityRequest, error) {
	return Create(ctx, c, "CoverageEligibilityRequest", data)
}

func (c *Client) SearchCoverageEligibilityRequest(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.CoverageEligibilityRequest, error) {
	return Search[fhirmodel.CoverageEligibilityRequest](ctx, c, "CoverageEligibilityRequest", params)
}

func (c *Client) ReadCoverageEligibilityResponse(ctx context.Context, id string) (*fhirmodel.CoverageEligibilityResponse, error) {
	return Read[fhirmodel.CoverageEligibilityResponse](ctx, c, "CoverageEligibilityResponse", id)
}

func (c *Client) UpdateCoverageEligibilityResponse(ctx context.Context, id string, data *fhirmodel.CoverageEligibilityResponse) (*fhirmodel.CoverageEligibilityResponse, error) {
	return Update(ctx, c, "CoverageEligibilityResponse", id, data)
}

func (c *Client) CreateCoverageEligibilityResponse(ctx context.Context, data *fhirmodel.CoverageEligibilityResponse) (*fhirmodel.CoverageEligibilityResponse, error) {
	return Create(ctx, c, "CoverageEligibilityResponse", data)
}

func (c *Client) SearchCoverageEligibilityResponse(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.CoverageEligibilityResponse, error) {
	return Search[fhirmodel.CoverageEligibilityResponse](ctx, c, "CoverageEligibilityResponse", params)
}

func (c *Client) ReadDetectedIssue(ctx context.Context, id string) (*fhirmodel.DetectedIssue, error) {
	return Read[fhirmodel.DetectedIssue](ctx, c, "DetectedIssue", id)
}

func (c *Client) UpdateDetectedIssue(ctx context.Context, id string, data *fhirmodel.DetectedIssue) (*fhirmodel.DetectedIssue, error) {
	return Update(ctx, c, "DetectedIssue", id, data)
}

func (c *Client) CreateDetectedIssue(ctx context.Context, data *fhirmodel.DetectedIssue) (*fhirmodel.DetectedIssue, error) {
	return Create(ctx, c, "DetectedIssue", data)
}

func (c *Client) SearchDetectedIssue(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.DetectedIssue, error) {
	return Search[fhirmodel.DetectedIssue](ctx, c, "DetectedIssue", params)
}

func (c *Client) ReadDevice(ctx context.Context, id string) (*fhirmodel.Device, error) {
	return Read[fhirmodel.Device](ctx, c, "Device", id)
}

func (c *Client) UpdateDevice(ctx context.Context, id string, data *fhirmodel.Device) (*fhirmodel.Device, error) {
	return Update(ctx, c, "Device", id, data)
}

func (c *Client) CreateDevice(ctx context.Context, data *fhirmodel.Device) (*fhirmodel.Device, error) {
	return Create(ctx, c, "Device", data)
}

func (c *Client) SearchDevice(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Device, error) {
	return Search[fhirmodel.Device](ctx, c, "Device", params)
}

func (c *Client) ReadDeviceDefinition(ctx context.Context, id string) (*fhirmodel.DeviceDefinition, error) {
	return Read[fhirmodel.DeviceDefinition](ctx, c, "DeviceDefinition", id)
}

func (c *Client) UpdateDeviceDefinition(ctx context.Context, id string, data *fhirmodel.DeviceDefinition) (*fhirmodel.DeviceDefinition, error) {
	return Update(ctx, c, "DeviceDefinition", id, data)
}

func (c *Client) CreateDeviceDefinition(ctx context.Context, data *fhirmodel.DeviceDefinition) (*fhirmodel.DeviceDefinition, error) {
	return Create(ctx, c, "DeviceDefinition", data)
}

func (c *Client) SearchDeviceDefinition(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.DeviceDefinition, error) {
	return Search[fhirmodel.DeviceDefinition](ctx, c, "DeviceDefinition", params)
}

func (c *Client) ReadDeviceMetric(ctx context.Context, id string) (*fhirmodel.DeviceMetric, error) {
	return Read[fhirmodel.DeviceMetric](ctx, c, "DeviceMetric", id)
}

func (c *Client) UpdateDeviceMetric(ctx context.Context, id string, data *fhirmodel.DeviceMetric) (*fhirmodel.DeviceMetric, error) {
	return Update(ctx, c, "DeviceMetric", id, data)
}

func (c *Client) CreateDeviceMetric(ctx context.Context, data *fhirmodel.DeviceMetric) (*fhirmodel.DeviceMetric, error) {
	return Create(ctx, c, "DeviceMetric", data)
}

func (c *Client) SearchDeviceMetric(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.DeviceMetric, error) {
	return Search[fhirmodel.DeviceMetric](ctx, c, "DeviceMetric", params)
}

func (c *Client) ReadDeviceRequest(ctx context.Context, id string) (*fhirmodel.DeviceRequest, error) {
	return Read[fhirmodel.DeviceRequest](ctx, c, "DeviceRequest", id)
}

func (c *Client) UpdateDeviceRequest(ctx context.Context, id string, data *fhirmodel.DeviceRequest) (*fhirmodel.DeviceRequest, error) {
	return Update(ctx, c, "DeviceRequest", id, data)
}

func (c *Client) CreateDeviceRequest(ctx context.Context, data *fhirmodel.DeviceRequest) (*fhirmodel.DeviceRequest, error) {
	return Create(ctx, c, "DeviceRequest", data)
}

func (c *Client) SearchDeviceRequest(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.DeviceRequest, error) {
	return Search[fhirmodel.DeviceRequest](ctx, c, "DeviceRequest", params)
}

func (c *Client) ReadDeviceUseStatement(ctx context.Context, id string) (*fhirmodel.DeviceUseStatement, error) {
	return Read[fhirmodel.DeviceUseStatement](ctx, c, "DeviceUseStatement", id)
}

func (c *Client) UpdateDeviceUseStatement(ctx context.Context, id string, data *fhirmodel.DeviceUseStatement) (*fhirmodel.DeviceUseStatement, error) {
	return Update(ctx, c, "DeviceUseStatement", id, data)
}

func (c *Client) CreateDeviceUseStatement(ctx context.Context, data *fhirmodel.DeviceUseStatement) (*fhirmodel.DeviceUseStatement, error) {
	return Create(ctx, c, "DeviceUseStatement", data)
}

func (c *Client) SearchDeviceUseStatement(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.DeviceUseStatement, error) {
	return Search[fhirmodel.DeviceUseStatement](ctx, c, "DeviceUseStatement", params)
}

func (c *Client) ReadDiagnosticReport(ctx context.Context, id string) (*fhirmodel.DiagnosticReport, error) {
	return Read[fhirmodel.DiagnosticReport](ctx, c, "DiagnosticReport", id)
}

func (c *Client) UpdateDiagnosticReport(ctx context.Context, id string, data *fhirmodel.DiagnosticReport) (*fhirmodel.DiagnosticReport, error) {
	return Update(ctx, c, "DiagnosticReport", id, data)
}

func (c *Client) CreateDiagnosticReport(ctx context.Context, data *fhirmodel.DiagnosticReport) (*fhirmodel.DiagnosticReport, error) {
	return Create(ctx, c, "DiagnosticReport", data)
}

func (c *Client) SearchDiagnosticReport(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.DiagnosticReport, error) {
	return Search[fhirmodel.DiagnosticReport](ctx, c, "DiagnosticReport", params)
}

func (c *Client) ReadDocumentManifest(ctx context.Context, id string) (*fhirmodel.DocumentManifest, error) {
	return Read[fhirmodel.DocumentManifest](ctx, c, "DocumentManifest", id)
}

func (c *Client) UpdateDocumentManifest(ctx context.Context, id string, data *fhirmodel.DocumentManifest) (*fhirmodel.DocumentManifest, error) {
	return Update(ctx, c, "DocumentManifest", id, data)
}

func (c *Client) CreateDocumentManifest(ctx context.Context, data *fhirmodel.DocumentManifest) (*fhirmodel.DocumentManifest, error) {
	return Create(ctx, c, "DocumentManifest", data)
}

func (c *Client) SearchDocumentManifest(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.DocumentManifest, error) {
	return Search[fhirmodel.DocumentManifest](ctx, c, "DocumentManifest", params)
}

func (c *Client) ReadDocumentReference(ctx context.Context, id string) (*fhirmodel.DocumentReference, error) {
	return Read[fhirmodel.DocumentReference](ctx, c, "DocumentReference", id)
}

func (c *Client) UpdateDocumentReference(ctx context.Context, id string, data *fhirmodel.DocumentReference) (*fhirmodel.DocumentReference, error) {
	return Update(ctx, c, "DocumentReference", id, data)
}

func (c *Client) CreateDocumentReference(ctx context.Context, data *fhirmodel.DocumentReference) (*fhirmodel.DocumentReference, error) {
	return Create(ctx, c, "DocumentReference", data)
}

func (c *Client) SearchDocumentReference(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.DocumentReference, error) {
	return Search[fhirmodel.DocumentReference](ctx, c, "DocumentReference", params)
}

func (c *Client) ReadEffectEvidenceSynthesis(ctx context.Context, id string) (*fhirmodel.EffectEvidenceSynthesis, error) {
	return Read[fhirmodel.EffectEvidenceSynthesis](ctx, c, "EffectEvidenceSynthesis", id)
}

func (c *Client) UpdateEffectEvidenceSynthesis(ctx context.Context, id string, data *fhirmodel.EffectEvidenceSynthesis) (*fhirmodel.EffectEvidenceSynthesis, error) {
	return Update(ctx, c, "EffectEvidenceSynthesis", id, data)
}

func (c *Client) CreateEffectEvidenceSynthesis(ctx context.Context, data *fhirmodel.EffectEvidenceSynthesis) (*fhirmodel.EffectEvidenceSynthesis, error) {
	return Create(ctx, c, "EffectEvidenceSynthesis", data)
}

func (c *Client) SearchEffectEvidenceSynthesis(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.EffectEvidenceSynthesis, error) {
	return Search[fhirmodel.EffectEvidenceSynthesis](ctx, c, "EffectEvidenceSynthesis", params)
}

func (c *Client) ReadEncounter(ctx context.Context, id string) (*fhirmodel.Encounter, error) {
	return Read[fhirmodel.Encounter](ctx, c, "Encounter", id)
}

func (c *Client) UpdateEncounter(ctx context.Context, id string, data *fhirmodel.Encounter) (*fhirmodel.Encounter, error) {
	return Update(ctx, c, "Encounter", id, data)
}

func (c *Client) CreateEncounter(ctx context.Context, data *fhirmodel.Encounter) (*fhirmodel.Encounter, error) {
	return Create(ctx, c, "Encounter", data)
}

func (c *Client) SearchEncounter(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Encounter, error) {
	return Search[fhirmodel.Encounter](ctx, c, "Encounter", params)
}

func (c *Client) ReadEndpoint(ctx context.Context, id string) (*fhirmodel.Endpoint, error) {
	return Read[fhirmodel.Endpoint](ctx, c, "Endpoint", id)
}

func (c *Client) UpdateEndpoint(ctx context.Context, id string, data *fhirmodel.Endpoint) (*fhirmodel.Endpoint, error) {
	return Update(ctx, c, "Endpoint", id, data)
}

func (c *Client) CreateEndpoint(ctx context.Context, data *fhirmodel.Endpoint) (*fhirmodel.Endpoint, error) {
	return Create(ctx, c, "Endpoint", data)
}

func (c *Client) SearchEndpoint(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Endpoint, error) {
	return Search[fhirmodel.Endpoint](ctx, c, "Endpoint", params)
}

func (c *Client) ReadEnrollmentRequest(ctx context.Context, id string) (*fhirmodel.EnrollmentRequest, error) {
	return Read[fhirmodel.EnrollmentRequest](ctx, c, "EnrollmentRequest", id)
}

func (c *Client) UpdateEnrollmentRequest(ctx context.Context, id string, data *fhirmodel.EnrollmentRequest) (*fhirmodel.EnrollmentRequest, error) {
	return Update(ctx, c, "EnrollmentRequest", id, data)
}

func (c *Client) CreateEnrollmentRequest(ctx context.Context, data *fhirmodel.EnrollmentRequest) (*fhirmodel.EnrollmentRequest, error) {
	return Create(ctx, c, "EnrollmentRequest", data)
}

func (c *Client) SearchEnrollmentRequest(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.EnrollmentRequest, error) {
	return Search[fhirmodel.EnrollmentRequest](ctx, c, "EnrollmentRequest", params)
}

func (c *Client) ReadEnrollmentResponse(ctx context.Context, id string) (*fhirmodel.EnrollmentResponse, error) {
	return Read[fhirmodel.EnrollmentResponse](ctx, c, "EnrollmentResponse", id)
}

func (c *Client) UpdateEnrollmentResponse(ctx context.Context, id string, data *fhirmodel.EnrollmentResponse) (*fhirmodel.EnrollmentResponse, error) {
	return Update(ctx, c, "EnrollmentResponse", id, data)
}

func (c *Client) CreateEnrollmentResponse(ctx context.Context, data *fhirmodel.EnrollmentResponse) (*fhirmodel.EnrollmentResponse, error) {
	return Create(ctx, c, "EnrollmentResponse", data)
}

func (c *Client) SearchEnrollmentResponse(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.EnrollmentResponse, error) {
	return Search[fhirmodel.EnrollmentResponse](ctx, c, "EnrollmentResponse", params)
}

func (c *Client) ReadEpisodeOfCare(ctx context.Context, id string) (*fhirmodel.EpisodeOfCare, error) {
	return Read[fhirmodel.EpisodeOfCare](ctx, c, "EpisodeOfCare", id)
}

func (c *Client) UpdateEpisodeOfCare(ctx context.Context, id string, data *fhirmodel.EpisodeOfCare) (*fhirmodel.EpisodeOfCare, error) {
	return Update(ctx, c, "EpisodeOfCare", id, data)
}

func (c *Client) CreateEpisodeOfCare(ctx context.Context, data *fhirmodel.EpisodeOfCare) (*fhirmodel.EpisodeOfCare, error) {
	return Create(ctx, c, "EpisodeOfCare", data)
}

func (c *Client) SearchEpisodeOfCare(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.EpisodeOfCare, error) {
	return Search[fhirmodel.EpisodeOfCare](ctx, c, "EpisodeOfCare", params)
}

func (c *Client) ReadEventDefinition(ctx context.Context, id string) (*fhirmodel.EventDefinition, error) {
	return Read[fhirmodel.EventDefinition](ctx, c, "EventDefinition", id)
}

func (c *Client) UpdateEventDefinition(ctx context.Context, id string, data *fhirmodel.EventDefinition) (*fhirmodel.EventDefinition, error) {
	return Update(ctx, c, "EventDefinition", id, data)
}

func (c *Client) CreateEventDefinition(ctx context.Context, data *fhirmodel.EventDefinition) (*fhirmodel.EventDefinition, error) {
	return Create(ctx, c, "EventDefinition", data)
}

func (c *Client) SearchEventDefinition(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.EventDefinition, error) {
	return Search[fhirmodel.EventDefinition](ctx, c, "EventDefinition", params)
}

func (c *Client) ReadEvidence(ctx context.Context, id string) (*fhirmodel.Evidence, error) {
	return Read[fhirmodel.Evidence](ctx, c, "Evidence", id)
}

func (c *Client) UpdateEvidence(ctx context.Context, id string, data *fhirmodel.Evidence) (*fhirmodel.Evidence, error) {
	return Update(ctx, c, "Evidence", id, data)
}

func (c *Client) CreateEvidence(ctx context.Context, data *fhirmodel.Evidence) (*fhirmodel.Evidence, error) {
	return Create(ctx, c, "Evidence", data)
}

func (c *Client) SearchEvidence(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Evidence, error) {
	return Search[fhirmodel.Evidence](ctx, c, "Evidence", params)
}

func (c *Client) ReadEvidenceVariable(ctx context.Context, id string) (*fhirmodel.EvidenceVariable, error) {
	return Read[fhirmodel.EvidenceVariable](ctx, c, "EvidenceVariable", id)
}

func (c *Client) UpdateEvidenceVariable(ctx context.Context, id string, data *fhirmodel.EvidenceVariable) (*fhirmodel.EvidenceVariable, error) {
	return Update(ctx, c, "EvidenceVariable", id, data)
}

func (c *Client) CreateEvidenceVariable(ctx context.Context, data *fhirmodel.EvidenceVariable) (*fhirmodel.EvidenceVariable, error) {
	return Create(ctx, c, "EvidenceVariable", data)
}

func (c *Client) SearchEvidenceVariable(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.EvidenceVariable, error) {
	return Search[fhirmodel.EvidenceVariable](ctx, c, "EvidenceVariable", params)
}

func (c *Client) ReadExampleScenario(ctx context.Context, id string) (*fhirmodel.ExampleScenario, error) {
	return Read[fhirmodel.ExampleScenario](ctx, c, "ExampleScenario", id)
}

func (c *Client) UpdateExampleScenario(ctx context.Context, id string, data *fhirmodel.ExampleScenario) (*fhirmodel.ExampleScenario, error) {
	return Update(ctx, c, "ExampleScenario", id, data)
}

func (c *Client) CreateExampleScenario(ctx context.Context, data *fhirmodel.ExampleScenario) (*fhirmodel.ExampleScenario, error) {
	return Create(ctx, c, "ExampleScenario", data)
}

func (c *Client) SearchExampleScenario(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.ExampleScenario, error) {
	return Search[fhirmodel.ExampleScenario](ctx, c, "ExampleScenario", params)
}

func (c *Client) ReadExplanationOfBenefit(ctx context.Context, id string) (*fhirmodel.ExplanationOfBenefit, error) {
	return Read[fhirmodel.ExplanationOfBenefit](ctx, c, "ExplanationOfBenefit", id)
}

func (c *Client) UpdateExplanationOfBenefit(ctx context.Context, id string, data *fhirmodel.ExplanationOfBenefit) (*fhirmodel.ExplanationOfBenefit, error) {
	return Update(ctx, c, "ExplanationOfBenefit", id, data)
}

func (c *Client) CreateExplanationOfBenefit(ctx context.Context, data *fhirmodel.ExplanationOfBenefit) (*fhirmodel.ExplanationOfBenefit, error) {
	return Create(ctx, c, "ExplanationOfBenefit", data)
}

func (c *Client) SearchExplanationOfBenefit(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.ExplanationOfBenefit, error) {
	return Search[fhirmodel.ExplanationOfBenefit](ctx, c, "ExplanationOfBenefit", params)
}

func (c *Client) ReadFamilyMemberHistory(ctx context.Context, id string) (*fhirmodel.FamilyMemberHistory, error) {
	return Read[fhirmodel.FamilyMemberHistory](ctx, c, "FamilyMemberHistory", id)
}

func (c *Client) UpdateFamilyMemberHistory(ctx context.Context, id string, data *fhirmodel.FamilyMemberHistory) (*fhirmodel.FamilyMemberHistory, error) {
	return Update(ctx, c, "FamilyMemberHistory", id, data)
}

func (c *Client) CreateFamilyMemberHistory(ctx context.Context, data *fhirmodel.FamilyMemberHistory) (*fhirmodel.FamilyMemberHistory, error) {
	return Create(ctx, c, "FamilyMemberHistory", data)
}

func (c *Client) SearchFamilyMemberHistory(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.FamilyMemberHistory, error) {
	return Search[fhirmodel.FamilyMemberHistory](ctx, c, "FamilyMemberHistory", params)
}

func (c *Client) ReadFlag(ctx context.Context, id string) (*fhirmodel.Flag, error) {
	return Read[fhirmodel.Flag](ctx, c, "Flag", id)
}

func (c *Client) UpdateFlag(ctx context.Context, id string, data *fhirmodel.Flag) (*fhirmodel.Flag, error) {
	return Update(ctx, c, "Flag", id, data)
}

func (c *Client) CreateFlag(ctx context.Context, data *fhirmodel.Flag) (*fhirmodel.Flag, error) {
	return Create(ctx, c, "Flag", data)
}

func (c *Client) SearchFlag(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Flag, error) {
	return Search[fhirmodel.Flag](ctx, c, "Flag", params)
}

func (c *Client) ReadGoal(ctx context.Context, id string) (*fhirmodel.Goal, error) {
	return Read[fhirmodel.Goal](ctx, c, "Goal", id)
}

func (c *Client) UpdateGoal(ctx context.Context, id string, data *fhirmodel.Goal) (*fhirmodel.Goal, error) {
	return Update(ctx, c, "Goal", id, data)
}

func (c *Client) CreateGoal(ctx context.Context, data *fhirmodel.Goal) (*fhirmodel.Goal, error) {
	return Create(ctx, c, "Goal", data)
}

func (c *Client) SearchGoal(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Goal, error) {
	return Search[fhirmodel.Goal](ctx, c, "Goal", params)
}

func (c *Client) ReadGraphDefinition(ctx context.Context, id string) (*fhirmodel.GraphDefinition, error) {
	return Read[fhirmodel.GraphDefinition](ctx, c, "GraphDefinition", id)
}

func (c *Client) UpdateGraphDefinition(ctx context.Context, id string, data *fhirmodel.GraphDefinition) (*fhirmodel.GraphDefinition, error) {
	return Update(ctx, c, "GraphDefinition", id, data)
}

func (c *Client) CreateGraphDefinition(ctx context.Context, data *fhirmodel.GraphDefinition) (*fhirmodel.GraphDefinition, error) {
	return Create(ctx, c, "GraphDefinition", data)
}

func (c *Client) SearchGraphDefinition(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.GraphDefinition, error) {
	return Search[fhirmodel.GraphDefinition](ctx, c, "GraphDefinition", params)
}

func (c *Client) ReadGroup(ctx context.Context, id string) (*fhirmodel.Group, error) {
	return Read[fhirmodel.Group](ctx, c, "Group", id)
}

func (c *Client) UpdateGroup(ctx context.Context, id string, data *fhirmodel.Group) (*fhirmodel.Group, error) {
	return Update(ctx, c, "Group", id, data)
}

func (c *Client) CreateGroup(ctx context.Context, data *fhirmodel.Group) (*fhirmodel.Group, error) {
	return Create(ctx, c, "Group", data)
}

func (c *Client) SearchGroup(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Group, error) {
	return Search[fhirmodel.Group](ctx, c, "Group", params)
}

func (c *Client) ReadGuidanceResponse(ctx context.Context, id string) (*fhirmodel.GuidanceResponse, error) {
	return Read[fhirmodel.GuidanceResponse](ctx, c, "GuidanceResponse", id)
}

func (c *Client) UpdateGuidanceResponse(ctx context.Context, id string, data *fhirmodel.GuidanceResponse) (*fhirmodel.GuidanceResponse, error) {
	return Update(ctx, c, "GuidanceResponse", id, data)
}

func (c *Client) CreateGuidanceResponse(ctx context.Context, data *fhirmodel.GuidanceResponse) (*fhirmodel.GuidanceResponse, error) {
	return Create(ctx, c, "GuidanceResponse", data)
}

func (c *Client) SearchGuidanceResponse(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.GuidanceResponse, error) {
	return Search[fhirmodel.GuidanceResponse](ctx, c, "GuidanceResponse", params)
}

func (c *Client) ReadHealthcareService(ctx context.Context, id string) (*fhirmodel.HealthcareService, error) {
	return Read[fhirmodel.HealthcareService](ctx, c, "HealthcareService", id)
}

func (c *Client) UpdateHealthcareService(ctx context.Context, id string, data *fhirmodel.HealthcareService) (*fhirmodel.HealthcareService, error) {
	return Update(ctx, c, "HealthcareService", id, data)
}

func (c *Client) CreateHealthcareService(ctx context.Context, data *fhirmodel.HealthcareService) (*fhirmodel.HealthcareService, error) {
	return Create(ctx, c, "HealthcareService", data)
}

func (c *Client) SearchHealthcareService(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.HealthcareService, error) {
	return Search[fhirmodel.HealthcareService](ctx, c, "HealthcareService", params)
}

func (c *Client) ReadImagingStudy(ctx context.Context, id string) (*fhirmodel.ImagingStudy, error) {
	return Read[fhirmodel.ImagingStudy](ctx, c, "ImagingStudy", id)
}

func (c *Client) UpdateImagingStudy(ctx context.Context, id string, data *fhirmodel.ImagingStudy) (*fhirmodel.ImagingStudy, error) {
	return Update(ctx, c, "ImagingStudy", id, data)
}

func (c *Client) CreateImagingStudy(ctx context.Context, data *fhirmodel.ImagingStudy) (*fhirmodel.ImagingStudy, error) {
	return Create(ctx, c, "ImagingStudy", data)
}

func (c *Client) SearchImagingStudy(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.ImagingStudy, error) {
	return Search[fhirmodel.ImagingStudy](ctx, c, "ImagingStudy", params)
}

func (c *Client) ReadImmunization(ctx context.Context, id string) (*fhirmodel.Immunization, error) {
	return Read[fhirmodel.Immunization](ctx, c, "Immunization", id)
}

func (c *Client) UpdateImmunization(ctx context.Context, id string, data *fhirmodel.Immunization) (*fhirmodel.Immunization, error) {
	return Update(ctx, c, "Immunization", id, data)
}

func (c *Client) CreateImmunization(ctx context.Context, data *fhirmodel.Immunization) (*fhirmodel.Immunization, error) {
	return Create(ctx, c, "Immunization", data)
}

func (c *Client) SearchImmunization(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Immunization, error) {
	return Search[fhirmodel.Immunization](ctx, c, "Immunization", params)
}

func (c *Client) ReadImmunizationEvaluation(ctx context.Context, id string) (*fhirmodel.ImmunizationEvaluation, error) {
	return Read[fhirmodel.ImmunizationEvaluation](ctx, c, "ImmunizationEvaluation", id)
}

func (c *Client) UpdateImmunizationEvaluation(ctx context.Context, id string, data *fhirmodel.ImmunizationEvaluation) (*fhirmodel.ImmunizationEvaluation, error) {
	return Update(ctx, c, "ImmunizationEvaluation", id, data)
}

func (c *Client) CreateImmunizationEvaluation(ctx context.Context, data *fhirmodel.ImmunizationEvaluation) (*fhirmodel.ImmunizationEvaluation, error) {
	return Create(ctx, c, "ImmunizationEvaluation", data)
}

func (c *Client) SearchImmunizationEvaluation(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.ImmunizationEvaluation, error) {
	return Search[fhirmodel.ImmunizationEvaluation](ctx, c, "ImmunizationEvaluation", params)
}

func (c *Client) ReadImmunizationRecommendation(ctx context.Context, id string) (*fhirmodel.ImmunizationRecommendation, error) {
	return Read[fhirmodel.ImmunizationRecommendation](ctx, c, "ImmunizationRecommendation", id)
}

func (c *Client) UpdateImmunizationRecommendation(ctx context.Context, id string, data *fhirmodel.ImmunizationRecommendation) (*fhirmodel.ImmunizationRecommendation, error) {
	return Update(ctx, c, "ImmunizationRecommendation", id, data)
}

func (c *Client) CreateImmunizationRecommendation(ctx context.Context, data *fhirmodel.ImmunizationRecommendation) (*fhirmodel.ImmunizationRecommendation, error) {
	return Create(ctx, c, "ImmunizationRecommendation", data)
}

func (c *Client) SearchImmunizationRecommendation(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.ImmunizationRecommendation, error) {
	return Search[fhirmodel.ImmunizationRecommendation](ctx, c, "ImmunizationRecommendation", params)
}

func (c *Client) ReadImplementationGuide(ctx context.Context, id string) (*fhirmodel.ImplementationGuide, error) {
	return Read[fhirmodel.ImplementationGuide](ctx, c, "ImplementationGuide", id)
}

func (c *Client) UpdateImplementationGuide(ctx context.Context, id string, data *fhirmodel.ImplementationGuide) (*fhirmodel.ImplementationGuide, error) {
	return Update(ctx, c, "ImplementationGuide", id, data)
}

func (c *Client) CreateImplementationGuide(ctx context.Context, data *fhirmodel.ImplementationGuide) (*fhirmodel.ImplementationGuide, error) {
	return Create(ctx, c, "ImplementationGuide", data)
}

func (c *Client) SearchImplementationGuide(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.ImplementationGuide, error) {
	return Search[fhirmodel.ImplementationGuide](ctx, c, "ImplementationGuide", params)
}

func (c *Client) ReadInsurancePlan(ctx context.Context, id string) (*fhirmodel.InsurancePlan, error) {
	return Read[fhirmodel.InsurancePlan](ctx, c, "InsurancePlan", id)
}

func (c *Client) UpdateInsurancePlan(ctx context.Context, id string, data *fhirmodel.InsurancePlan) (*fhirmodel.InsurancePlan, error) {
	return Update(ctx, c, "InsurancePlan", id, data)
}

func (c *Client) CreateInsurancePlan(ctx context.Context, data *fhirmodel.InsurancePlan) (*fhirmodel.InsurancePlan, error) {
	return Create(ctx, c, "InsurancePlan", data)
}

func (c *Client) SearchInsurancePlan(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.InsurancePlan, error) {
	return Search[fhirmodel.InsurancePlan](ctx, c, "InsurancePlan", params)
}

func (c *Client) ReadInvoice(ctx context.Context, id string) (*fhirmodel.Invoice, error) {
	return Read[fhirmodel.Invoice](ctx, c, "Invoice", id)
}

func (c *Client) UpdateInvoice(ctx context.Context, id string, data *fhirmodel.Invoice) (*fhirmodel.Invoice, error) {
	return Update(ctx, c, "Invoice", id, data)
}

func (c *Client) CreateInvoice(ctx context.Context, data *fhirmodel.Invoice) (*fhirmodel.Invoice, error) {
	return Create(ctx, c, "Invoice", data)
}

func (c *Client) SearchInvoice(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Invoice, error) {
	return Search[fhirmodel.Invoice](ctx, c, "Invoice", params)
}

func (c *Client) ReadLibrary(ctx context.Context, id string) (*fhirmodel.Library, error) {
	return Read[fhirmodel.Library](ctx, c, "Library", id)
}

func (c *Client) UpdateLibrary(ctx context.Context, id string, data *fhirmodel.Library) (*fhirmodel.Library, error) {
	return Update(ctx, c, "Library", id, data)
}

func (c *Client) CreateLibrary(ctx context.Context, data *fhirmodel.Library) (*fhirmodel.Library, error) {
	return Create(ctx, c, "Library", data)
}

func (c *Client) SearchLibrary(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Library, error) {
	return Search[fhirmodel.Library](ctx, c, "Library", params)
}

func (c *Client) ReadLinkage(ctx context.Context, id string) (*fhirmodel.Linkage, error) {
	return Read[fhirmodel.Linkage](ctx, c, "Linkage", id)
}

func (c *Client) UpdateLinkage(ctx context.Context, id string, data *fhirmodel.Linkage) (*fhirmodel.Linkage, error) {
	return Update(ctx, c, "Linkage", id, data)
}

func (c *Client) CreateLinkage(ctx context.Context, data *fhirmodel.Linkage) (*fhirmodel.Linkage, error) {
	return Create(ctx, c, "Linkage", data)
}

func (c *Client) SearchLinkage(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Linkage, error) {
	return Search[fhirmodel.Linkage](ctx, c, "Linkage", params)
}

func (c *Client) ReadList(ctx context.Context, id string) (*fhirmodel.List, error) {
	return Read[fhirmodel.List](ctx, c, "List", id)
}

func (c *Client) UpdateList(ctx context.Context, id string, data *fhirmodel.List) (*fhirmodel.List, error) {
	return Update(ctx, c, "List", id, data)
}

func (c *Client) CreateList(ctx context.Context, data *fhirmodel.List) (*fhirmodel.List, error) {
	return Create(ctx, c, "List", data)
}

func (c *Client) SearchList(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.List, error) {
	return Search[fhirmodel.List](ctx, c, "List", params)
}

func (c *Client) ReadLocation(ctx context.Context, id string) (*fhirmodel.Location, error) {
	return Read[fhirmodel.Location](ctx, c, "Location", id)
}

func (c *Client) UpdateLocation(ctx context.Context, id string, data *fhirmodel.Location) (*fhirmodel.Location, error) {
	return Update(ctx, c, "Location", id, data)
}

func (c *Client) CreateLocation(ctx context.Context, data *fhirmodel.Location) (*fhirmodel.Location, error) {
	return Create(ctx, c, "Location", data)
}

func (c *Client) SearchLocation(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Location, error) {
	return Search[fhirmodel.Location](ctx, c, "Location", params)
}

func (c *Client) ReadMeasure(ctx context.Context, id string) (*fhirmodel.Measure, error) {
	return Read[fhirmodel.Measure](ctx, c, "Measure", id)
}

func (c *Client) UpdateMeasure(ctx context.Context, id string, data *fhirmodel.Measure) (*fhirmodel.Measure, error) {
	return Update(ctx, c, "Measure", id, data)
}

func (c *Client) CreateMeasure(ctx context.Context, data *fhirmodel.Measure) (*fhirmodel.Measure, error) {
	return Create(ctx, c, "Measure", data)
}

func (c *Client) SearchMeasure(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Measure, error) {
	return Search[fhirmodel.Measure](ctx, c, "Measure", params)
}

func (c *Client) ReadMeasureReport(ctx context.Context, id string) (*fhirmodel.MeasureReport, error) {
	return Read[fhirmodel.MeasureReport](ctx, c, "MeasureReport", id)
}

func (c *Client) UpdateMeasureReport(ctx context.Context, id string, data *fhirmodel.MeasureReport) (*fhirmodel.MeasureReport, error) {
	return Update(ctx, c, "MeasureReport", id, data)
}

func (c *Client) CreateMeasureReport(ctx context.Context, data *fhirmodel.MeasureReport) (*fhirmodel.MeasureReport, error) {
	return Create(ctx, c, "MeasureReport", data)
}

func (c *Client) SearchMeasureReport(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.MeasureReport, error) {
	return Search[fhirmodel.MeasureReport](ctx, c, "MeasureReport", params)
}

func (c *Client) ReadMedia(ctx context.Context, id string) (*fhirmodel.Media, error) {
	return Read[fhirmodel.Media](ctx, c, "Media", id)
}

func (c *Client) UpdateMedia(ctx context.Context, id string, data *fhirmodel.Media) (*fhirmodel.Media, error) {
	return Update(ctx, c, "Media", id, data)
}

func (c *Client) CreateMedia(ctx context.Context, data *fhirmodel.Media) (*fhirmodel.Media, error) {
	return Create(ctx, c, "Media", data)
}

func (c *Client) SearchMedia(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Media, error) {
	return Search[fhirmodel.Media](ctx, c, "Media", params)
}

func (c *Client) ReadMedication(ctx context.Context, id string) (*fhirmodel.Medication, error) {
	return Read[fhirmodel.Medication](ctx, c, "Medication", id)
}

func (c *Client) UpdateMedication(ctx context.Context, id string, data *fhirmodel.Medication) (*fhirmodel.Medication, error) {
	return Update(ctx, c, "Medication", id, data)
}

func (c *Client) CreateMedication(ctx context.Context, data *fhirmodel.Medication) (*fhirmodel.Medication, error) {
	return Create(ctx, c, "Medication", data)
}

func (c *Client) SearchMedication(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Medication, error) {
	return Search[fhirmodel.Medication](ctx, c, "Medication", params)
}

func (c *Client) ReadMedicationAdministration(ctx context.Context, id string) (*fhirmodel.MedicationAdministration, error) {
	return Read[fhirmodel.MedicationAdministration](ctx, c, "MedicationAdministration", id)
}

func (c *Client) UpdateMedicationAdministration(ctx context.Context, id string, data *fhirmodel.MedicationAdministration) (*fhirmodel.MedicationAdministration, error) {
	return Update(ctx, c, "MedicationAdministration", id, data)
}

func (c *Client) CreateMedicationAdministration(ctx context.Context, data *fhirmodel.MedicationAdministration) (*fhirmodel.MedicationAdministration, error) {
	return Create(ctx, c, "MedicationAdministration", data)
}

func (c *Client) SearchMedicationAdministration(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.MedicationAdministration, error) {
	return Search[fhirmodel.MedicationAdministration](ctx, c, "MedicationAdministration", params)
}

func (c *Client) ReadMedicationDispense(ctx context.Context, id string) (*fhirmodel.MedicationDispense, error) {
	return Read[fhirmodel.MedicationDispense](ctx, c, "MedicationDispense", id)
}

func (c *Client) UpdateMedicationDispense(ctx context.Context, id string, data *fhirmodel.MedicationDispense) (*fhirmodel.MedicationDispense, error) {
	return Update(ctx, c, "MedicationDispense", id, data)
}

func (c *Client) CreateMedicationDispense(ctx context.Context, data *fhirmodel.MedicationDispense) (*fhirmodel.MedicationDispense, error) {
	return Create(ctx, c, "MedicationDispense", data)
}

func (c *Client) SearchMedicationDispense(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.MedicationDispense, error) {
	return Search[fhirmodel.MedicationDispense](ctx, c, "MedicationDispense", params)
}

func (c *Client) ReadMedicationKnowledge(ctx context.Context, id string) (*fhirmodel.MedicationKnowledge, error) {
	return Read[fhirmodel.MedicationKnowledge](ctx, c, "MedicationKnowledge", id)
}

func (c *Client) UpdateMedicationKnowledge(ctx context.Context, id string, data *fhirmodel.MedicationKnowledge) (*fhirmodel.MedicationKnowledge, error) {
	return Update(ctx, c, "MedicationKnowledge", id, data)
}

func (c *Client) CreateMedicationKnowledge(ctx context.Context, data *fhirmodel.MedicationKnowledge) (*fhirmodel.MedicationKnowledge, error) {
	return Create(ctx, c, "MedicationKnowledge", data)
}

func (c *Client) SearchMedicationKnowledge(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.MedicationKnowledge, error) {
	return Search[fhirmodel.MedicationKnowledge](ctx, c, "MedicationKnowledge", params)
}

func (c *Client) ReadMedicationRequest(ctx context.Context, id string) (*fhirmodel.MedicationRequest, error) {
	return Read[fhirmodel.MedicationRequest](ctx, c, "MedicationRequest", id)
}

func (c *Client) UpdateMedicationRequest(ctx context.Context, id string, data *fhirmodel.MedicationRequest) (*fhirmodel.MedicationRequest, error) {
	return Update(ctx, c, "MedicationRequest", id, data)
}

func (c *Client) CreateMedicationRequest(ctx context.Context, data *fhirmodel.MedicationRequest) (*fhirmodel.MedicationRequest, error) {
	return Create(ctx, c, "MedicationRequest", data)
}

func (c *Client) SearchMedicationRequest(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.MedicationRequest, error) {
	return Search[fhirmodel.MedicationRequest](ctx, c, "MedicationRequest", params)
}

func (c *Client) ReadMedicationStatement(ctx context.Context, id string) (*fhirmodel.MedicationStatement, error) {
	return Read[fhirmodel.MedicationStatement](ctx, c, "MedicationStatement", id)
}

func (c *Client) UpdateMedicationStatement(ctx context.Context, id string, data *fhirmodel.MedicationStatement) (*fhirmodel.MedicationStatement, error) {
	return Update(ctx, c, "MedicationStatement", id, data)
}

func (c *Client) CreateMedicationStatement(ctx context.Context, data *fhirmodel.MedicationStatement) (*fhirmodel.MedicationStatement, error) {
	return Create(ctx, c, "MedicationStatement", data)
}

func (c *Client) SearchMedicationStatement(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.MedicationStatement, error) {
	return Search[fhirmodel.MedicationStatement](ctx, c, "MedicationStatement", params)
}

func (c *Client) ReadMedicinalProduct(ctx context.Context, id string) (*fhirmodel.MedicinalProduct, error) {
	return Read[fhirmodel.MedicinalProduct](ctx, c, "MedicinalProduct", id)
}

func (c *Client) UpdateMedicinalProduct(ctx context.Context, id string, data *fhirmodel.MedicinalProduct) (*fhirmodel.MedicinalProduct, error) {
	return Update(ctx, c, "MedicinalProduct", id, data)
}

func (c *Client) CreateMedicinalProduct(ctx context.Context, data *fhirmodel.MedicinalProduct) (*fhirmodel.MedicinalProduct, error) {
	return Create(ctx, c, "MedicinalProduct", data)
}

func (c *Client) SearchMedicinalProduct(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.MedicinalProduct, error) {
	return Search[fhirmodel.MedicinalProduct](ctx, c, "MedicinalProduct", params)
}

func (c *Client) ReadMedicinalProductAuthorization(ctx context.Context, id string) (*fhirmodel.MedicinalProductAuthorization, error) {
	return Read[fhirmodel.MedicinalProductAuthorization](ctx, c, "MedicinalProductAuthorization", id)
}

func (c *Client) UpdateMedicinalProductAuthorization(ctx context.Context, id string, data *fhirmodel.MedicinalProductAuthorization) (*fhirmodel.MedicinalProductAuthorization, error) {
	return Update(ctx, c, "MedicinalProductAuthorization", id, data)
}

func (c *Client) CreateMedicinalProductAuthorization(ctx context.Context, data *fhirmodel.MedicinalProductAuthorization) (*fhirmodel.MedicinalProductAuthorization, error) {
	return Create(ctx, c, "MedicinalProductAuthorization", data)
}

func (c *Client) SearchMedicinalProductAuthorization(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.MedicinalProductAuthorization, error) {
	return Search[fhirmodel.MedicinalProductAuthorization](ctx, c, "MedicinalProductAuthorization", params)
}

func (c *Client) ReadMedicinalProductContraindication(ctx context.Context, id string) (*fhirmodel.MedicinalProductContraindication, error) {
	return Read[fhirmodel.MedicinalProductContraindication](ctx, c, "MedicinalProductContraindication", id)
}

func (c *Client) UpdateMedicinalProductContraindication(ctx context.Context, id string, data *fhirmodel.MedicinalProductContraindication) (*fhirmodel.MedicinalProductContraindication, error) {
	return Update(ctx, c, "MedicinalProductContraindication", id, data)
}

func (c *Client) CreateMedicinalProductContraindication(ctx context.Context, data *fhirmodel.MedicinalProductContraindication) (*fhirmodel.MedicinalProductContraindication, error) {
	return Create(ctx, c, "MedicinalProductContraindication", data)
}

func (c *Client) SearchMedicinalProductContraindication(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.MedicinalProductContraindication, error) {
	return Search[fhirmodel.MedicinalProductContraindication](ctx, c, "MedicinalProductContraindication", params)
}

func (c *Client) ReadMedicinalProductIndication(ctx context.Context, id string) (*fhirmodel.MedicinalProductIndication, error) {
	return Read[fhirmodel.MedicinalProductIndication](ctx, c, "MedicinalProductIndication", id)
}

func (c *Client) UpdateMedicinalProductIndication(ctx context.Context, id string, data *fhirmodel.MedicinalProductIndication) (*fhirmodel.MedicinalProductIndication, error) {
	return Update(ctx, c, "MedicinalProductIndication", id, data)
}

func (c *Client) CreateMedicinalProductIndication(ctx context.Context, data *fhirmodel.MedicinalProductIndication) (*fhirmodel.MedicinalProductIndication, error) {
	return Create(ctx, c, "MedicinalProductIndication", data)
}

func (c *Client) SearchMedicinalProductIndication(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.MedicinalProductIndication, error) {
	return Search[fhirmodel.MedicinalProductIndication](ctx, c, "MedicinalProductIndication", params)
}

func (c *Client) ReadMedicinalProductIngredient(ctx context.Context, id string) (*fhirmodel.MedicinalProductIngredient, error) {
	return Read[fhirmodel.MedicinalProductIngredient](ctx, c, "MedicinalProductIngredient", id)
}

func (c *Client) UpdateMedicinalProductIngredient(ctx context.Context, id string, data *fhirmodel.MedicinalProductIngredient) (*fhirmodel.MedicinalProductIngredient, error) {
	return Update(ctx, c, "MedicinalProductIngredient", id, data)
}

func (c *Client) CreateMedicinalProductIngredient(ctx context.Context, data *fhirmodel.MedicinalProductIngredient) (*fhirmodel.MedicinalProductIngredient, error) {
	return Create(ctx, c, "MedicinalProductIngredient", data)
}

func (c *Client) SearchMedicinalProductIngredient(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.MedicinalProductIngredient, error) {
	return Search[fhirmodel.MedicinalProductIngredient](ctx, c, "MedicinalProductIngredient", params)
}

func (c *Client) ReadMedicinalProductInteraction(ctx context.Context, id string) (*fhirmodel.MedicinalProductInteraction, error) {
	return Read[fhirmodel.MedicinalProductInteraction](ctx, c, "MedicinalProductInteraction", id)
}

func (c *Client) UpdateMedicinalProductInteraction(ctx context.Context, id string, data *fhirmodel.MedicinalProductInteraction) (*fhirmodel.MedicinalProductInteraction, error) {
	return Update(ctx, c, "MedicinalProductInteraction", id, data)
}

func (c *Client) CreateMedicinalProductInteraction(ctx context.Context, data *fhirmodel.MedicinalProductInteraction) (*fhirmodel.MedicinalProductInteraction, error) {
	return Create(ctx, c, "MedicinalProductInteraction", data)
}

func (c *Client) SearchMedicinalProductInteraction(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.MedicinalProductInteraction, error) {
	return Search[fhirmodel.MedicinalProductInteraction](ctx, c, "MedicinalProductInteraction", params)
}

func (c *Client) ReadMedicinalProductManufactured(ctx context.Context, id string) (*fhirmodel.MedicinalProductManufactured, error) {
	return Read[fhirmodel.MedicinalProductManufactured](ctx, c, "MedicinalProductManufactured", id)
}

func (c *Client) UpdateMedicinalProductManufactured(ctx context.Context, id string, data *fhirmodel.MedicinalProductManufactured) (*fhirmodel.MedicinalProductManufactured, error) {
	return Update(ctx, c, "MedicinalProductManufactured", id, data)
}

func (c *Client) CreateMedicinalProductManufactured(ctx context.Context, data *fhirmodel.MedicinalProductManufactured) (*fhirmodel.MedicinalProductManufactured, error) {
	return Create(ctx, c, "MedicinalProductManufactured", data)
}

func (c *Client) SearchMedicinalProductManufactured(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.MedicinalProductManufactured, error) {
	return Search[fhirmodel.MedicinalProductManufactured](ctx, c, "MedicinalProductManufactured", params)
}

func (c *Client) ReadMedicinalProductPackaged(ctx context.Context, id string) (*fhirmodel.MedicinalProductPackaged, error) {
	return Read[fhirmodel.MedicinalProductPackaged](ctx, c, "MedicinalProductPackaged", id)
}

func (c *Client) UpdateMedicinalProductPackaged(ctx context.Context, id string, data *fhirmodel.MedicinalProductPackaged) (*fhirmodel.MedicinalProductPackaged, error) {
	return Update(ctx, c, "MedicinalProductPackaged", id, data)
}

func (c *Client) CreateMedicinalProductPackaged(ctx context.Context, data *fhirmodel.MedicinalProductPackaged) (*fhirmodel.MedicinalProductPackaged, error) {
	return Create(ctx, c, "MedicinalProductPackaged", data)
}

func (c *Client) SearchMedicinalProductPackaged(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.MedicinalProductPackaged, error) {
	return Search[fhirmodel.MedicinalProductPackaged](ctx, c, "MedicinalProductPackaged", params)
}

func (c *Client) ReadMedicinalProductPharmaceutical(ctx context.Context, id string) (*fhirmodel.MedicinalProductPharmaceutical, error) {
	return Read[fhirmodel.MedicinalProductPharmaceutical](ctx, c, "MedicinalProductPharmaceutical", id)
}

func (c *Client) UpdateMedicinalProductPharmaceutical(ctx context.Context, id string, data *fhirmodel.MedicinalProductPharmaceutical) (*fhirmodel.MedicinalProductPharmaceutical, error) {
	return Update(ctx, c, "MedicinalProductPharmaceutical", id, data)
}

func (c *Client) CreateMedicinalProductPharmaceutical(ctx context.Context, data *fhirmodel.MedicinalProductPharmaceutical) (*fhirmodel.MedicinalProductPharmaceutical, error) {
	return Create(ctx, c, "MedicinalProductPharmaceutical", data)
}

func (c *Client) SearchMedicinalProductPharmaceutical(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.MedicinalProductPharmaceutical, error) {
	return Search[fhirmodel.MedicinalProductPharmaceutical](ctx, c, "MedicinalProductPharmaceutical", params)
}

func (c *Client) ReadMedicinalProductUndesirableEffect(ctx context.Context, id string) (*fhirmodel.MedicinalProductUndesirableEffect, error) {
	return Read[fhirmodel.MedicinalProductUndesirableEffect](ctx, c, "MedicinalProductUndesirableEffect", id)
}

func (c *Client) UpdateMedicinalProductUndesirableEffect(ctx context.Context, id string, data *fhirmodel.MedicinalProductUndesirableEffect) (*fhirmodel.MedicinalProductUndesirableEffect, error) {
	return Update(ctx, c, "MedicinalProductUndesirableEffect", id, data)
}

func (c *Client) CreateMedicinalProductUndesirableEffect(ctx context.Context, data *fhirmodel.MedicinalProductUndesirableEffect) (*fhirmodel.MedicinalProductUndesirableEffect, error) {
	return Create(ctx, c, "MedicinalProductUndesirableEffect", data)
}

func (c *Client) SearchMedicinalProductUndesirableEffect(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.MedicinalProductUndesirableEffect, error) {
	return Search[fhirmodel.MedicinalProductUndesirableEffect](ctx, c, "MedicinalProductUndesirableEffect", params)
}

func (c *Client) ReadMessageDefinition(ctx context.Context, id string) (*fhirmodel.MessageDefinition, error) {
	return Read[fhirmodel.MessageDefinition](ctx, c, "MessageDefinition", id)
}

func (c *Client) UpdateMessageDefinition(ctx context.Context, id string, data *fhirmodel.MessageDefinition) (*fhirmodel.MessageDefinition, error) {
	return Update(ctx, c, "MessageDefinition", id, data)
}

func (c *Client) CreateMessageDefinition(ctx context.Context, data *fhirmodel.MessageDefinition) (*fhirmodel.MessageDefinition, error) {
	return Create(ctx, c, "MessageDefinition", data)
}

func (c *Client) SearchMessageDefinition(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.MessageDefinition, error) {
	return Search[fhirmodel.MessageDefinition](ctx, c, "MessageDefinition", params)
}

func (c *Client) ReadMessageHeader(ctx context.Context, id string) (*fhirmodel.MessageHeader, error) {
	return Read[fhirmodel.MessageHeader](ctx, c, "MessageHeader", id)
}

func (c *Client) UpdateMessageHeader(ctx context.Context, id string, data *fhirmodel.MessageHeader) (*fhirmodel.MessageHeader, error) {
	return Update(ctx, c, "MessageHeader", id, data)
}

func (c *Client) CreateMessageHeader(ctx context.Context, data *fhirmodel.MessageHeader) (*fhirmodel.MessageHeader, error) {
	return Create(ctx, c, "MessageHeader", data)
}

func (c *Client) SearchMessageHeader(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.MessageHeader, error) {
	return Search[fhirmodel.MessageHeader](ctx, c, "MessageHeader", params)
}

func (c *Client) ReadMolecularSequence(ctx context.Context, id string) (*fhirmodel.MolecularSequence, error) {
	return Read[fhirmodel.MolecularSequence](ctx, c, "MolecularSequence", id)
}

func (c *Client) UpdateMolecularSequence(ctx context.Context, id string, data *fhirmodel.MolecularSequence) (*fhirmodel.MolecularSequence, error) {
	return Update(ctx, c, "MolecularSequence", id, data)
}

func (c *Client) CreateMolecularSequence(ctx context.Context, data *fhirmodel.MolecularSequence) (*fhirmodel.MolecularSequence, error) {
	return Create(ctx, c, "MolecularSequence", data)
}

func (c *Client) SearchMolecularSequence(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.MolecularSequence, error) {
	return Search[fhirmodel.MolecularSequence](ctx, c, "MolecularSequence", params)
}

func (c *Client) ReadNamingSystem(ctx context.Context, id string) (*fhirmodel.NamingSystem, error) {
	return Read[fhirmodel.NamingSystem](ctx, c, "NamingSystem", id)
}

func (c *Client) UpdateNamingSystem(ctx context.Context, id string, data *fhirmodel.NamingSystem) (*fhirmodel.NamingSystem, error) {
	return Update(ctx, c, "NamingSystem", id, data)
}

func (c *Client) CreateNamingSystem(ctx context.Context, data *fhirmodel.NamingSystem) (*fhirmodel.NamingSystem, error) {
	return Create(ctx, c, "NamingSystem", data)
}

func (c *Client) SearchNamingSystem(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.NamingSystem, error) {
	return Search[fhirmodel.NamingSystem](ctx, c, "NamingSystem", params)
}

func (c *Client) ReadNutritionOrder(ctx context.Context, id string) (*fhirmodel.NutritionOrder, error) {
	return Read[fhirmodel.NutritionOrder](ctx, c, "NutritionOrder", id)
}

func (c *Client) UpdateNutritionOrder(ctx context.Context, id string, data *fhirmodel.NutritionOrder) (*fhirmodel.NutritionOrder, error) {
	return Update(ctx, c, "NutritionOrder", id, data)
}

func (c *Client) CreateNutritionOrder(ctx context.Context, data *fhirmodel.NutritionOrder) (*fhirmodel.NutritionOrder, error) {
	return Create(ctx, c, "NutritionOrder", data)
}

func (c *Client) SearchNutritionOrder(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.NutritionOrder, error) {
	return Search[fhirmodel.NutritionOrder](ctx, c, "NutritionOrder", params)
}

func (c *Client) ReadObservation(ctx context.Context, id string) (*fhirmodel.Observation, error) {
	return Read[fhirmodel.Observation](ctx, c, "Observation", id)
}

func (c *Client) UpdateObservation(ctx context.Context, id string, data *fhirmodel.Observation) (*fhirmodel.Observation, error) {
	return Update(ctx, c, "Observation", id, data)
}

func (c *Client) CreateObservation(ctx context.Context, data *fhirmodel.Observation) (*fhirmodel.Observation, error) {
	return Create(ctx, c, "Observation", data)
}

func (c *Client) SearchObservation(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Observation, error) {
	return Search[fhirmodel.Observation](ctx, c, "Observation", params)
}

func (c *Client) ReadObservationDefinition(ctx context.Context, id string) (*fhirmodel.ObservationDefinition, error) {
	return Read[fhirmodel.ObservationDefinition](ctx, c, "ObservationDefinition", id)
}

func (c *Client) UpdateObservationDefinition(ctx context.Context, id string, data *fhirmodel.ObservationDefinition) (*fhirmodel.ObservationDefinition, error) {
	return Update(ctx, c, "ObservationDefinition", id, data)
}

func (c *Client) CreateObservationDefinition(ctx context.Context, data *fhirmodel.ObservationDefinition) (*fhirmodel.ObservationDefinition, error) {
	return Create(ctx, c, "ObservationDefinition", data)
}

func (c *Client) SearchObservationDefinition(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.ObservationDefinition, error) {
	return Search[fhirmodel.ObservationDefinition](ctx, c, "ObservationDefinition", params)
}

func (c *Client) ReadOperationDefinition(ctx context.Context, id string) (*fhirmodel.OperationDefinition, error) {
	return Read[fhirmodel.OperationDefinition](ctx, c, "OperationDefinition", id)
}

func (c *Client) UpdateOperationDefinition(ctx context.Context, id string, data *fhirmodel.OperationDefinition) (*fhirmodel.OperationDefinition, error) {
	return Update(ctx, c, "OperationDefinition", id, data)
}

func (c *Client) CreateOperationDefinition(ctx context.Context, data *fhirmodel.OperationDefinition) (*fhirmodel.OperationDefinition, error) {
	return Create(ctx, c, "OperationDefinition", data)
}

func (c *Client) SearchOperationDefinition(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.OperationDefinition, error) {
	return Search[fhirmodel.OperationDefinition](ctx, c, "OperationDefinition", params)
}

func (c *Client) ReadOperationOutcome(ctx context.Context, id string) (*fhirmodel.OperationOutcome, error) {
	return Read[fhirmodel.OperationOutcome](ctx, c, "OperationOutcome", id)
}

func (c *Client) UpdateOperationOutcome(ctx context.Context, id string, data *fhirmodel.OperationOutcome) (*fhirmodel.OperationOutcome, error) {
	return Update(ctx, c, "OperationOutcome", id, data)
}

func (c *Client) CreateOperationOutcome(ctx context.Context, data *fhirmodel.OperationOutcome) (*fhirmodel.OperationOutcome, error) {
	return Create(ctx, c, "OperationOutcome", data)
}

func (c *Client) SearchOperationOutcome(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.OperationOutcome, error) {
	return Search[fhirmodel.OperationOutcome](ctx, c, "OperationOutcome", params)
}

func (c *Client) ReadOrganization(ctx context.Context, id string) (*fhirmodel.Organization, error) {
	return Read[fhirmodel.Organization](ctx, c, "Organization", id)
}

func (c *Client) UpdateOrganization(ctx context.Context, id string, data *fhirmodel.Organization) (*fhirmodel.Organization, error) {
	return Update(ctx, c, "Organization", id, data)
}

func (c *Client) CreateOrganization(ctx context.Context, data *fhirmodel.Organization) (*fhirmodel.Organization, error) {
	return Create(ctx, c, "Organization", data)
}

func (c *Client) SearchOrganization(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Organization, error) {
	return Search[fhirmodel.Organization](ctx, c, "Organization", params)
}

func (c *Client) ReadOrganizationAffiliation(ctx context.Context, id string) (*fhirmodel.OrganizationAffiliation, error) {
	return Read[fhirmodel.OrganizationAffiliation](ctx, c, "OrganizationAffiliation", id)
}

func (c *Client) UpdateOrganizationAffiliation(ctx context.Context, id string, data *fhirmodel.OrganizationAffiliation) (*fhirmodel.OrganizationAffiliation, error) {
	return Update(ctx, c, "OrganizationAffiliation", id, data)
}

func (c *Client) CreateOrganizationAffiliation(ctx context.Context, data *fhirmodel.OrganizationAffiliation) (*fhirmodel.OrganizationAffiliation, error) {
	return Create(ctx, c, "OrganizationAffiliation", data)
}

func (c *Client) SearchOrganizationAffiliation(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.OrganizationAffiliation, error) {
	return Search[fhirmodel.OrganizationAffiliation](ctx, c, "OrganizationAffiliation", params)
}

func (c *Client) ReadPatient(ctx context.Context, id string) (*fhirmodel.Patient, error) {
	return Read[fhirmodel.Patient](ctx, c, "Patient", id)
}

func (c *Client) UpdatePatient(ctx context.Context, id string, data *fhirmodel.Patient) (*fhirmodel.Patient, error) {
	return Update(ctx, c, "Patient", id, data)
}

func (c *Client) CreatePatient(ctx context.Context, data *fhirmodel.Patient) (*fhirmodel.Patient, error) {
	return Create(ctx, c, "Patient", data)
}

func (c *Client) SearchPatient(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Patient, error) {
	return Search[fhirmodel.Patient](ctx, c, "Patient", params)
}

func (c *Client) ReadPaymentNotice(ctx context.Context, id string) (*fhirmodel.PaymentNotice, error) {
	return Read[fhirmodel.PaymentNotice](ctx, c, "PaymentNotice", id)
}

func (c *Client) UpdatePaymentNotice(ctx context.Context, id string, data *fhirmodel.PaymentNotice) (*fhirmodel.PaymentNotice, error) {
	return Update(ctx, c, "PaymentNotice", id, data)
}

func (c *Client) CreatePaymentNotice(ctx context.Context, data *fhirmodel.PaymentNotice) (*fhirmodel.PaymentNotice, error) {
	return Create(ctx, c, "PaymentNotice", data)
}

func (c *Client) SearchPaymentNotice(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.PaymentNotice, error) {
	return Search[fhirmodel.PaymentNotice](ctx, c, "PaymentNotice", params)
}

func (c *Client) ReadPaymentReconciliation(ctx context.Context, id string) (*fhirmodel.PaymentReconciliation, error) {
	return Read[fhirmodel.PaymentReconciliation](ctx, c, "PaymentReconciliation", id)
}

func (c *Client) UpdatePaymentReconciliation(ctx context.Context, id string, data *fhirmodel.PaymentReconciliation) (*fhirmodel.PaymentReconciliation, error) {
	return Update(ctx, c, "PaymentReconciliation", id, data)
}

func (c *Client) CreatePaymentReconciliation(ctx context.Context, data *fhirmodel.PaymentReconciliation) (*fhirmodel.PaymentReconciliation, error) {
	return Create(ctx, c, "PaymentReconciliation", data)
}

func (c *Client) SearchPaymentReconciliation(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.PaymentReconciliation, error) {
	return Search[fhirmodel.PaymentReconciliation](ctx, c, "PaymentReconciliation", params)
}

func (c *Client) ReadPerson(ctx context.Context, id string) (*fhirmodel.Person, error) {
	return Read[fhirmodel.Person](ctx, c, "Person", id)
}

func (c *Client) UpdatePerson(ctx context.Context, id string, data *fhirmodel.Person) (*fhirmodel.Person, error) {
	return Update(ctx, c, "Person", id, data)
}

func (c *Client) CreatePerson(ctx context.Context, data *fhirmodel.Person) (*fhirmodel.Person, error) {
	return Create(ctx, c, "Person", data)
}

func (c *Client) SearchPerson(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Person, error) {
	return Search[fhirmodel.Person](ctx, c, "Person", params)
}

func (c *Client) ReadPlanDefinition(ctx context.Context, id string) (*fhirmodel.PlanDefinition, error) {
	return Read[fhirmodel.PlanDefinition](ctx, c, "PlanDefinition", id)
}

func (c *Client) UpdatePlanDefinition(ctx context.Context, id string, data *fhirmodel.PlanDefinition) (*fhirmodel.PlanDefinition, error) {
	return Update(ctx, c, "PlanDefinition", id, data)
}

func (c *Client) CreatePlanDefinition(ctx context.Context, data *fhirmodel.PlanDefinition) (*fhirmodel.PlanDefinition, error) {
	return Create(ctx, c, "PlanDefinition", data)
}

func (c *Client) SearchPlanDefinition(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.PlanDefinition, error) {
	return Search[fhirmodel.PlanDefinition](ctx, c, "PlanDefinition", params)
}

func (c *Client) ReadPractitioner(ctx context.Context, id string) (*fhirmodel.Practitioner, error) {
	return Read[fhirmodel.Practitioner](ctx, c, "Practitioner", id)
}

func (c *Client) UpdatePractitioner(ctx context.Context, id string, data *fhirmodel.Practitioner) (*fhirmodel.Practitioner, error) {
	return Update(ctx, c, "Practitioner", id, data)
}

func (c *Client) CreatePractitioner(ctx context.Context, data *fhirmodel.Practitioner) (*fhirmodel.Practitioner, error) {
	return Create(ctx, c, "Practitioner", data)
}

func (c *Client) SearchPractitioner(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Practitioner, error) {
	return Search[fhirmodel.Practitioner](ctx, c, "Practitioner", params)
}

func (c *Client) ReadPractitionerRole(ctx context.Context, id string) (*fhirmodel.PractitionerRole, error) {
	return Read[fhirmodel.PractitionerRole](ctx, c, "PractitionerRole", id)
}

func (c *Client) UpdatePractitionerRole(ctx context.Context, id string, data *fhirmodel.PractitionerRole) (*fhirmodel.PractitionerRole, error) {
	return Update(ctx, c, "PractitionerRole", id, data)
}

func (c *Client) CreatePractitionerRole(ctx context.Context, data *fhirmodel.PractitionerRole) (*fhirmodel.PractitionerRole, error) {
	return Create(ctx, c, "PractitionerRole", data)
}

func (c *Client) SearchPractitionerRole(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.PractitionerRole, error) {
	return Search[fhirmodel.PractitionerRole](ctx, c, "PractitionerRole", params)
}

func (c *Client) ReadProcedure(ctx context.Context, id string) (*fhirmodel.Procedure, error) {
	return Read[fhirmodel.Procedure](ctx, c, "Procedure", id)
}

func (c *Client) UpdateProcedure(ctx context.Context, id string, data *fhirmodel.Procedure) (*fhirmodel.Procedure, error) {
	return Update(ctx, c, "Procedure", id, data)
}

func (c *Client) CreateProcedure(ctx context.Context, data *fhirmodel.Procedure) (*fhirmodel.Procedure, error) {
	return Create(ctx, c, "Procedure", data)
}

func (c *Client) SearchProcedure(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Procedure, error) {
	return Search[fhirmodel.Procedure](ctx, c, "Procedure", params)
}

func (c *Client) ReadProvenance(ctx context.Context, id string) (*fhirmodel.Provenance, error) {
	return Read[fhirmodel.Provenance](ctx, c, "Provenance", id)
}

func (c *Client) UpdateProvenance(ctx context.Context, id string, data *fhirmodel.Provenance) (*fhirmodel.Provenance, error) {
	return Update(ctx, c, "Provenance", id, data)
}

func (c *Client) CreateProvenance(ctx context.Context, data *fhirmodel.Provenance) (*fhirmodel.Provenance, error) {
	return Create(ctx, c, "Provenance", data)
}

func (c *Client) SearchProvenance(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Provenance, error) {
	return Search[fhirmodel.Provenance](ctx, c, "Provenance", params)
}

func (c *Client) ReadQuestionnaire(ctx context.Context, id string) (*fhirmodel.Questionnaire, error) {
	return Read[fhirmodel.Questionnaire](ctx, c, "Questionnaire", id)
}

func (c *Client) UpdateQuestionnaire(ctx context.Context, id string, data *fhirmodel.Questionnaire) (*fhirmodel.Questionnaire, error) {
	return Update(ctx, c, "Questionnaire", id, data)
}

func (c *Client) CreateQuestionnaire(ctx context.Context, data *fhirmodel.Questionnaire) (*fhirmodel.Questionnaire, error) {
	return Create(ctx, c, "Questionnaire", data)
}

func (c *Client) SearchQuestionnaire(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Questionnaire, error) {
	return Search[fhirmodel.Questionnaire](ctx, c, "Questionnaire", params)
}

func (c *Client) ReadQuestionnaireResponse(ctx context.Context, id string) (*fhirmodel.QuestionnaireResponse, error) {
	return Read[fhirmodel.QuestionnaireResponse](ctx, c, "QuestionnaireResponse", id)
}

func (c *Client) UpdateQuestionnaireResponse(ctx context.Context, id string, data *fhirmodel.QuestionnaireResponse) (*fhirmodel.QuestionnaireResponse, error) {
	return Update(ctx, c, "QuestionnaireResponse", id, data)
}

func (c *Client) CreateQuestionnaireResponse(ctx context.Context, data *fhirmodel.QuestionnaireResponse) (*fhirmodel.QuestionnaireResponse, error) {
	return Create(ctx, c, "QuestionnaireResponse", data)
}

func (c *Client) SearchQuestionnaireResponse(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.QuestionnaireResponse, error) {
	return Search[fhirmodel.QuestionnaireResponse](ctx, c, "QuestionnaireResponse", params)
}

func (c *Client) ReadRelatedPerson(ctx context.Context, id string) (*fhirmodel.RelatedPerson, error) {
	return Read[fhirmodel.RelatedPerson](ctx, c, "RelatedPerson", id)
}

func (c *Client) UpdateRelatedPerson(ctx context.Context, id string, data *fhirmodel.RelatedPerson) (*fhirmodel.RelatedPerson, error) {
	return Update(ctx, c, "RelatedPerson", id, data)
}

func (c *Client) CreateRelatedPerson(ctx context.Context, data *fhirmodel.RelatedPerson) (*fhirmodel.RelatedPerson, error) {
	return Create(ctx, c, "RelatedPerson", data)
}

func (c *Client) SearchRelatedPerson(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.RelatedPerson, error) {
	return Search[fhirmodel.RelatedPerson](ctx, c, "RelatedPerson", params)
}

func (c *Client) ReadRequestGroup(ctx context.Context, id string) (*fhirmodel.RequestGroup, error) {
	return Read[fhirmodel.RequestGroup](ctx, c, "RequestGroup", id)
}

func (c *Client) UpdateRequestGroup(ctx context.Context, id string, data *fhirmodel.RequestGroup) (*fhirmodel.RequestGroup, error) {
	return Update(ctx, c, "RequestGroup", id, data)
}

func (c *Client) CreateRequestGroup(ctx context.Context, data *fhirmodel.RequestGroup) (*fhirmodel.RequestGroup, error) {
	return Create(ctx, c, "RequestGroup", data)
}

func (c *Client) SearchRequestGroup(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.RequestGroup, error) {
	return Search[fhirmodel.RequestGroup](ctx, c, "RequestGroup", params)
}

func (c *Client) ReadResearchDefinition(ctx context.Context, id string) (*fhirmodel.ResearchDefinition, error) {
	return Read[fhirmodel.ResearchDefinition](ctx, c, "ResearchDefinition", id)
}

func (c *Client) UpdateResearchDefinition(ctx context.Context, id string, data *fhirmodel.ResearchDefinition) (*fhirmodel.ResearchDefinition, error) {
	return Update(ctx, c, "ResearchDefinition", id, data)
}

func (c *Client) CreateResearchDefinition(ctx context.Context, data *fhirmodel.ResearchDefinition) (*fhirmodel.ResearchDefinition, error) {
	return Create(ctx, c, "ResearchDefinition", data)
}

func (c *Client) SearchResearchDefinition(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.ResearchDefinition, error) {
	return Search[fhirmodel.ResearchDefinition](ctx, c, "ResearchDefinition", params)
}

func (c *Client) ReadResearchElementDefinition(ctx context.Context, id string) (*fhirmodel.ResearchElementDefinition, error) {
	return Read[fhirmodel.ResearchElementDefinition](ctx, c, "ResearchElementDefinition", id)
}

func (c *Client) UpdateResearchElementDefinition(ctx context.Context, id string, data *fhirmodel.ResearchElementDefinition) (*fhirmodel.ResearchElementDefinition, error) {
	return Update(ctx, c, "ResearchElementDefinition", id, data)
}

func (c *Client) CreateResearchElementDefinition(ctx context.Context, data *fhirmodel.ResearchElementDefinition) (*fhirmodel.ResearchElementDefinition, error) {
	return Create(ctx, c, "ResearchElementDefinition", data)
}

func (c *Client) SearchResearchElementDefinition(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.ResearchElementDefinition, error) {
	return Search[fhirmodel.ResearchElementDefinition](ctx, c, "ResearchElementDefinition", params)
}

func (c *Client) ReadResearchStudy(ctx context.Context, id string) (*fhirmodel.ResearchStudy, error) {
	return Read[fhirmodel.ResearchStudy](ctx, c, "ResearchStudy", id)
}

func (c *Client) UpdateResearchStudy(ctx context.Context, id string, data *fhirmodel.ResearchStudy) (*fhirmodel.ResearchStudy, error) {
	return Update(ctx, c, "ResearchStudy", id, data)
}

func (c *Client) CreateResearchStudy(ctx context.Context, data *fhirmodel.ResearchStudy) (*fhirmodel.ResearchStudy, error) {
	return Create(ctx, c, "ResearchStudy", data)
}

func (c *Client) SearchResearchStudy(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.ResearchStudy, error) {
	return Search[fhirmodel.ResearchStudy](ctx, c, "ResearchStudy", params)
}

func (c *Client) ReadResearchSubject(ctx context.Context, id string) (*fhirmodel.ResearchSubject, error) {
	return Read[fhirmodel.ResearchSubject](ctx, c, "ResearchSubject", id)
}

func (c *Client) UpdateResearchSubject(ctx context.Context, id string, data *fhirmodel.ResearchSubject) (*fhirmodel.ResearchSubject, error) {
	return Update(ctx, c, "ResearchSubject", id, data)
}

func (c *Client) CreateResearchSubject(ctx context.Context, data *fhirmodel.ResearchSubject) (*fhirmodel.ResearchSubject, error) {
	return Create(ctx, c, "ResearchSubject", data)
}

func (c *Client) SearchResearchSubject(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.ResearchSubject, error) {
	return Search[fhirmodel.ResearchSubject](ctx, c, "ResearchSubject", params)
}

func (c *Client) ReadRiskAssessment(ctx context.Context, id string) (*fhirmodel.RiskAssessment, error) {
	return Read[fhirmodel.RiskAssessment](ctx, c, "RiskAssessment", id)
}

func (c *Client) UpdateRiskAssessment(ctx context.Context, id string, data *fhirmodel.RiskAssessment) (*fhirmodel.RiskAssessment, error) {
	return Update(ctx, c, "RiskAssessment", id, data)
}

func (c *Client) CreateRiskAssessment(ctx context.Context, data *fhirmodel.RiskAssessment) (*fhirmodel.RiskAssessment, error) {
	return Create(ctx, c, "RiskAssessment", data)
}

func (c *Client) SearchRiskAssessment(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.RiskAssessment, error) {
	return Search[fhirmodel.RiskAssessment](ctx, c, "RiskAssessment", params)
}

func (c *Client) ReadRiskEvidenceSynthesis(ctx context.Context, id string) (*fhirmodel.RiskEvidenceSynthesis, error) {
	return Read[fhirmodel.RiskEvidenceSynthesis](ctx, c, "RiskEvidenceSynthesis", id)
}

func (c *Client) UpdateRiskEvidenceSynthesis(ctx context.Context, id string, data *fhirmodel.RiskEvidenceSynthesis) (*fhirmodel.RiskEvidenceSynthesis, error) {
	return Update(ctx, c, "RiskEvidenceSynthesis", id, data)
}

func (c *Client) CreateRiskEvidenceSynthesis(ctx context.Context, data *fhirmodel.RiskEvidenceSynthesis) (*fhirmodel.RiskEvidenceSynthesis, error) {
	return Create(ctx, c, "RiskEvidenceSynthesis", data)
}

func (c *Client) SearchRiskEvidenceSynthesis(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.RiskEvidenceSynthesis, error) {
	return Search[fhirmodel.RiskEvidenceSynthesis](ctx, c, "RiskEvidenceSynthesis", params)
}

func (c *Client) ReadSchedule(ctx context.Context, id string) (*fhirmodel.Schedule, error) {
	return Read[fhirmodel.Schedule](ctx, c, "Schedule", id)
}

func (c *Client) UpdateSchedule(ctx context.Context, id string, data *fhirmodel.Schedule) (*fhirmodel.Schedule, error) {
	return Update(ctx, c, "Schedule", id, data)
}

func (c *Client) CreateSchedule(ctx context.Context, data *fhirmodel.Schedule) (*fhirmodel.Schedule, error) {
	return Create(ctx, c, "Schedule", data)
}

func (c *Client) SearchSchedule(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Schedule, error) {
	return Search[fhirmodel.Schedule](ctx, c, "Schedule", params)
}

func (c *Client) ReadSearchParameter(ctx context.Context, id string) (*fhirmodel.SearchParameter, error) {
	return Read[fhirmodel.SearchParameter](ctx, c, "SearchParameter", id)
}

func (c *Client) UpdateSearchParameter(ctx context.Context, id string, data *fhirmodel.SearchParameter) (*fhirmodel.SearchParameter, error) {
	return Update(ctx, c, "SearchParameter", id, data)
}

func (c *Client) CreateSearchParameter(ctx context.Context, data *fhirmodel.SearchParameter) (*fhirmodel.SearchParameter, error) {
	return Create(ctx, c, "SearchParameter", data)
}

func (c *Client) SearchSearchParameter(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.SearchParameter, error) {
	return Search[fhirmodel.SearchParameter](ctx, c, "SearchParameter", params)
}

func (c *Client) ReadServiceRequest(ctx context.Context, id string) (*fhirmodel.ServiceRequest, error) {
	return Read[fhirmodel.ServiceRequest](ctx, c, "ServiceRequest", id)
}

func (c *Client) UpdateServiceRequest(ctx context.Context, id string, data *fhirmodel.ServiceRequest) (*fhirmodel.ServiceRequest, error) {
	return Update(ctx, c, "ServiceRequest", id, data)
}

func (c *Client) CreateServiceRequest(ctx context.Context, data *fhirmodel.ServiceRequest) (*fhirmodel.ServiceRequest, error) {
	return Create(ctx, c, "ServiceRequest", data)
}

func (c *Client) SearchServiceRequest(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.ServiceRequest, error) {
	return Search[fhirmodel.ServiceRequest](ctx, c, "ServiceRequest", params)
}

func (c *Client) ReadSlot(ctx context.Context, id string) (*fhirmodel.Slot, error) {
	return Read[fhirmodel.Slot](ctx, c, "Slot", id)
}

func (c *Client) UpdateSlot(ctx context.Context, id string, data *fhirmodel.Slot) (*fhirmodel.Slot, error) {
	return Update(ctx, c, "Slot", id, data)
}

func (c *Client) CreateSlot(ctx context.Context, data *fhirmodel.Slot) (*fhirmodel.Slot, error) {
	return Create(ctx, c, "Slot", data)
}

func (c *Client) SearchSlot(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Slot, error) {
	return Search[fhirmodel.Slot](ctx, c, "Slot", params)
}

func (c *Client) ReadSpecimen(ctx context.Context, id string) (*fhirmodel.Specimen, error) {
	return Read[fhirmodel.Specimen](ctx, c, "Specimen", id)
}

func (c *Client) UpdateSpecimen(ctx context.Context, id string, data *fhirmodel.Specimen) (*fhirmodel.Specimen, error) {
	return Update(ctx, c, "Specimen", id, data)
}

func (c *Client) CreateSpecimen(ctx context.Context, data *fhirmodel.Specimen) (*fhirmodel.Specimen, error) {
	return Create(ctx, c, "Specimen", data)
}

func (c *Client) SearchSpecimen(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Specimen, error) {
	return Search[fhirmodel.Specimen](ctx, c, "Specimen", params)
}

func (c *Client) ReadSpecimenDefinition(ctx context.Context, id string) (*fhirmodel.SpecimenDefinition, error) {
	return Read[fhirmodel.SpecimenDefinition](ctx, c, "SpecimenDefinition", id)
}

func (c *Client) UpdateSpecimenDefinition(ctx context.Context, id string, data *fhirmodel.SpecimenDefinition) (*fhirmodel.SpecimenDefinition, error) {
	return Update(ctx, c, "SpecimenDefinition", id, data)
}

func (c *Client) CreateSpecimenDefinition(ctx context.Context, data *fhirmodel.SpecimenDefinition) (*fhirmodel.SpecimenDefinition, error) {
	return Create(ctx, c, "SpecimenDefinition", data)
}

func (c *Client) SearchSpecimenDefinition(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.SpecimenDefinition, error) {
	return Search[fhirmodel.SpecimenDefinition](ctx, c, "SpecimenDefinition", params)
}

func (c *Client) ReadStructureDefinition(ctx context.Context, id string) (*fhirmodel.StructureDefinition, error) {
	return Read[fhirmodel.StructureDefinition](ctx, c, "StructureDefinition", id)
}

func (c *Client) UpdateStructureDefinition(ctx context.Context, id string, data *fhirmodel.StructureDefinition) (*fhirmodel.StructureDefinition, error) {
	return Update(ctx, c, "StructureDefinition", id, data)
}

func (c *Client) CreateStructureDefinition(ctx context.Context, data *fhirmodel.StructureDefinition) (*fhirmodel.StructureDefinition, error) {
	return Create(ctx, c, "StructureDefinition", data)
}

func (c *Client) SearchStructureDefinition(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.StructureDefinition, error) {
	return Search[fhirmodel.StructureDefinition](ctx, c, "StructureDefinition", params)
}

func (c *Client) ReadStructureMap(ctx context.Context, id string) (*fhirmodel.StructureMap, error) {
	return Read[fhirmodel.StructureMap](ctx, c, "StructureMap", id)
}

func (c *Client) UpdateStructureMap(ctx context.Context, id string, data *fhirmodel.StructureMap) (*fhirmodel.StructureMap, error) {
	return Update(ctx, c, "StructureMap", id, data)
}

func (c *Client) CreateStructureMap(ctx context.Context, data *fhirmodel.StructureMap) (*fhirmodel.StructureMap, error) {
	return Create(ctx, c, "StructureMap", data)
}

func (c *Client) SearchStructureMap(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.StructureMap, error) {
	return Search[fhirmodel.StructureMap](ctx, c, "StructureMap", params)
}

func (c *Client) ReadSubscription(ctx context.Context, id string) (*fhirmodel.Subscription, error) {
	return Read[fhirmodel.Subscription](ctx, c, "Subscription", id)
}

func (c *Client) UpdateSubscription(ctx context.Context, id string, data *fhirmodel.Subscription) (*fhirmodel.Subscription, error) {
	return Update(ctx, c, "Subscription", id, data)
}

func (c *Client) CreateSubscription(ctx context.Context, data *fhirmodel.Subscription) (*fhirmodel.Subscription, error) {
	return Create(ctx, c, "Subscription", data)
}

func (c *Client) SearchSubscription(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Subscription, error) {
	return Search[fhirmodel.Subscription](ctx, c, "Subscription", params)
}

func (c *Client) ReadSubstance(ctx context.Context, id string) (*fhirmodel.Substance, error) {
	return Read[fhirmodel.Substance](ctx, c, "Substance", id)
}

func (c *Client) UpdateSubstance(ctx context.Context, id string, data *fhirmodel.Substance) (*fhirmodel.Substance, error) {
	return Update(ctx, c, "Substance", id, data)
}

func (c *Client) CreateSubstance(ctx context.Context, data *fhirmodel.Substance) (*fhirmodel.Substance, error) {
	return Create(ctx, c, "Substance", data)
}

func (c *Client) SearchSubstance(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Substance, error) {
	return Search[fhirmodel.Substance](ctx, c, "Substance", params)
}

func (c *Client) ReadSubstanceNucleicAcid(ctx context.Context, id string) (*fhirmodel.SubstanceNucleicAcid, error) {
	return Read[fhirmodel.SubstanceNucleicAcid](ctx, c, "SubstanceNucleicAcid", id)
}

func (c *Client) UpdateSubstanceNucleicAcid(ctx context.Context, id string, data *fhirmodel.SubstanceNucleicAcid) (*fhirmodel.SubstanceNucleicAcid, error) {
	return Update(ctx, c, "SubstanceNucleicAcid", id, data)
}

func (c *Client) CreateSubstanceNucleicAcid(ctx context.Context, data *fhirmodel.SubstanceNucleicAcid) (*fhirmodel.SubstanceNucleicAcid, error) {
	return Create(ctx, c, "SubstanceNucleicAcid", data)
}

func (c *Client) SearchSubstanceNucleicAcid(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.SubstanceNucleicAcid, error) {
	return Search[fhirmodel.SubstanceNucleicAcid](ctx, c, "SubstanceNucleicAcid", params)
}

func (c *Client) ReadSubstancePolymer(ctx context.Context, id string) (*fhirmodel.SubstancePolymer, error) {
	return Read[fhirmodel.SubstancePolymer](ctx, c, "SubstancePolymer", id)
}

func (c *Client) UpdateSubstancePolymer(ctx context.Context, id string, data *fhirmodel.SubstancePolymer) (*fhirmodel.SubstancePolymer, error) {
	return Update(ctx, c, "SubstancePolymer", id, data)
}

func (c *Client) CreateSubstancePolymer(ctx context.Context, data *fhirmodel.SubstancePolymer) (*fhirmodel.SubstancePolymer, error) {
	return Create(ctx, c, "SubstancePolymer", data)
}

func (c *Client) SearchSubstancePolymer(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.SubstancePolymer, error) {
	return Search[fhirmodel.SubstancePolymer](ctx, c, "SubstancePolymer", params)
}

func (c *Client) ReadSubstanceProtein(ctx context.Context, id string) (*fhirmodel.SubstanceProtein, error) {
	return Read[fhirmodel.SubstanceProtein](ctx, c, "SubstanceProtein", id)
}

func (c *Client) UpdateSubstanceProtein(ctx context.Context, id string, data *fhirmodel.SubstanceProtein) (*fhirmodel.SubstanceProtein, error) {
	return Update(ctx, c, "SubstanceProtein", id, data)
}

func (c *Client) CreateSubstanceProtein(ctx context.Context, data *fhirmodel.SubstanceProtein) (*fhirmodel.SubstanceProtein, error) {
	return Create(ctx, c, "SubstanceProtein", data)
}

func (c *Client) SearchSubstanceProtein(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.SubstanceProtein, error) {
	return Search[fhirmodel.SubstanceProtein](ctx, c, "SubstanceProtein", params)
}

func (c *Client) ReadSubstanceReferenceInformation(ctx context.Context, id string) (*fhirmodel.SubstanceReferenceInformation, error) {
	return Read[fhirmodel.SubstanceReferenceInformation](ctx, c, "SubstanceReferenceInformation", id)
}

func (c *Client) UpdateSubstanceReferenceInformation(ctx context.Context, id string, data *fhirmodel.SubstanceReferenceInformation) (*fhirmodel.SubstanceReferenceInformation, error) {
	return Update(ctx, c, "SubstanceReferenceInformation", id, data)
}

func (c *Client) CreateSubstanceReferenceInformation(ctx context.Context, data *fhirmodel.SubstanceReferenceInformation) (*fhirmodel.SubstanceReferenceInformation, error) {
	return Create(ctx, c, "SubstanceReferenceInformation", data)
}

func (c *Client) SearchSubstanceReferenceInformation(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.SubstanceReferenceInformation, error) {
	return Search[fhirmodel.SubstanceReferenceInformation](ctx, c, "SubstanceReferenceInformation", params)
}

func (c *Client) ReadSubstanceSourceMaterial(ctx context.Context, id string) (*fhirmodel.SubstanceSourceMaterial, error) {
	return Read[fhirmodel.SubstanceSourceMaterial](ctx, c, "SubstanceSourceMaterial", id)
}

func (c *Client) UpdateSubstanceSourceMaterial(ctx context.Context, id string, data *fhirmodel.SubstanceSourceMaterial) (*fhirmodel.SubstanceSourceMaterial, error) {
	return Update(ctx, c, "SubstanceSourceMaterial", id, data)
}

func (c *Client) CreateSubstanceSourceMaterial(ctx context.Context, data *fhirmodel.SubstanceSourceMaterial) (*fhirmodel.SubstanceSourceMaterial, error) {
	return Create(ctx, c, "SubstanceSourceMaterial", data)
}

func (c *Client) SearchSubstanceSourceMaterial(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.SubstanceSourceMaterial, error) {
	return Search[fhirmodel.SubstanceSourceMaterial](ctx, c, "SubstanceSourceMaterial", params)
}

func (c *Client) ReadSubstanceSpecification(ctx context.Context, id string) (*fhirmodel.SubstanceSpecification, error) {
	return Read[fhirmodel.SubstanceSpecification](ctx, c, "SubstanceSpecification", id)
}

func (c *Client) UpdateSubstanceSpecification(ctx context.Context, id string, data *fhirmodel.SubstanceSpecification) (*fhirmodel.SubstanceSpecification, error) {
	return Update(ctx, c, "SubstanceSpecification", id, data)
}

func (c *Client) CreateSubstanceSpecification(ctx context.Context, data *fhirmodel.SubstanceSpecification) (*fhirmodel.SubstanceSpecification, error) {
	return Create(ctx, c, "SubstanceSpecification", data)
}

func (c *Client) SearchSubstanceSpecification(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.SubstanceSpecification, error) {
	return Search[fhirmodel.SubstanceSpecification](ctx, c, "SubstanceSpecification", params)
}

func (c *Client) ReadSupplyDelivery(ctx context.Context, id string) (*fhirmodel.SupplyDelivery, error) {
	return Read[fhirmodel.SupplyDelivery](ctx, c, "SupplyDelivery", id)
}

func (c *Client) UpdateSupplyDelivery(ctx context.Context, id string, data *fhirmodel.SupplyDelivery) (*fhirmodel.SupplyDelivery, error) {
	return Update(ctx, c, "SupplyDelivery", id, data)
}

func (c *Client) CreateSupplyDelivery(ctx context.Context, data *fhirmodel.SupplyDelivery) (*fhirmodel.SupplyDelivery, error) {
	return Create(ctx, c, "SupplyDelivery", data)
}

func (c *Client) SearchSupplyDelivery(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.SupplyDelivery, error) {
	return Search[fhirmodel.SupplyDelivery](ctx, c, "SupplyDelivery", params)
}

func (c *Client) ReadSupplyRequest(ctx context.Context, id string) (*fhirmodel.SupplyRequest, error) {
	return Read[fhirmodel.SupplyRequest](ctx, c, "SupplyRequest", id)
}

func (c *Client) UpdateSupplyRequest(ctx context.Context, id string, data *fhirmodel.SupplyRequest) (*fhirmodel.SupplyRequest, error) {
	return Update(ctx, c, "SupplyRequest", id, data)
}

func (c *Client) CreateSupplyRequest(ctx context.Context, data *fhirmodel.SupplyRequest) (*fhirmodel.SupplyRequest, error) {
	return Create(ctx, c, "SupplyRequest", data)
}

func (c *Client) SearchSupplyRequest(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.SupplyRequest, error) {
	return Search[fhirmodel.SupplyRequest](ctx, c, "SupplyRequest", params)
}

func (c *Client) ReadTask(ctx context.Context, id string) (*fhirmodel.Task, error) {
	return Read[fhirmodel.Task](ctx, c, "Task", id)
}

func (c *Client) UpdateTask(ctx context.Context, id string, data *fhirmodel.Task) (*fhirmodel.Task, error) {
	return Update(ctx, c, "Task", id, data)
}

func (c *Client) CreateTask(ctx context.Context, data *fhirmodel.Task) (*fhirmodel.Task, error) {
	return Create(ctx, c, "Task", data)
}

func (c *Client) SearchTask(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.Task, error) {
	return Search[fhirmodel.Task](ctx, c, "Task", params)
}

func (c *Client) ReadTerminologyCapabilities(ctx context.Context, id string) (*fhirmodel.TerminologyCapabilities, error) {
	return Read[fhirmodel.TerminologyCapabilities](ctx, c, "TerminologyCapabilities", id)
}

func (c *Client) UpdateTerminologyCapabilities(ctx context.Context, id string, data *fhirmodel.TerminologyCapabilities) (*fhirmodel.TerminologyCapabilities, error) {
	return Update(ctx, c, "TerminologyCapabilities", id, data)
}

func (c *Client) CreateTerminologyCapabilities(ctx context.Context, data *fhirmodel.TerminologyCapabilities) (*fhirmodel.TerminologyCapabilities, error) {
	return Create(ctx, c, "TerminologyCapabilities", data)
}

func (c *Client) SearchTerminologyCapabilities(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.TerminologyCapabilities, error) {
	return Search[fhirmodel.TerminologyCapabilities](ctx, c, "TerminologyCapabilities", params)
}

func (c *Client) ReadTestReport(ctx context.Context, id string) (*fhirmodel.TestReport, error) {
	return Read[fhirmodel.TestReport](ctx, c, "TestReport", id)
}

func (c *Client) UpdateTestReport(ctx context.Context, id string, data *fhirmodel.TestReport) (*fhirmodel.TestReport, error) {
	return Update(ctx, c, "TestReport", id, data)
}

func (c *Client) CreateTestReport(ctx context.Context, data *fhirmodel.TestReport) (*fhirmodel.TestReport, error) {
	return Create(ctx, c, "TestReport", data)
}

func (c *Client) SearchTestReport(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.TestReport, error) {
	return Search[fhirmodel.TestReport](ctx, c, "TestReport", params)
}

func (c *Client) ReadTestScript(ctx context.Context, id string) (*fhirmodel.TestScript, error) {
	return Read[fhirmodel.TestScript](ctx, c, "TestScript", id)
}

func (c *Client) UpdateTestScript(ctx context.Context, id string, data *fhirmodel.TestScript) (*fhirmodel.TestScript, error) {
	return Update(ctx, c, "TestScript", id, data)
}

func (c *Client) CreateTestScript(ctx context.Context, data *fhirmodel.TestScript) (*fhirmodel.TestScript, error) {
	return Create(ctx, c, "TestScript", data)
}

func (c *Client) SearchTestScript(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.TestScript, error) {
	return Search[fhirmodel.TestScript](ctx, c, "TestScript", params)
}

func (c *Client) ReadValueSet(ctx context.Context, id string) (*fhirmodel.ValueSet, error) {
	return Read[fhirmodel.ValueSet](ctx, c, "ValueSet", id)
}

func (c *Client) UpdateValueSet(ctx context.Context, id string, data *fhirmodel.ValueSet) (*fhirmodel.ValueSet, error) {
	return Update(ctx, c, "ValueSet", id, data)
}

func (c *Client) CreateValueSet(ctx context.Context, data *fhirmodel.ValueSet) (*fhirmodel.ValueSet, error) {
	return Create(ctx, c, "ValueSet", data)
}

func (c *Client) SearchValueSet(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.ValueSet, error) {
	return Search[fhirmodel.ValueSet](ctx, c, "ValueSet", params)
}

func (c *Client) ReadVerificationResult(ctx context.Context, id string) (*fhirmodel.VerificationResult, error) {
	return Read[fhirmodel.VerificationResult](ctx, c, "VerificationResult", id)
}

func (c *Client) UpdateVerificationResult(ctx context.Context, id string, data *fhirmodel.VerificationResult) (*fhirmodel.VerificationResult, error) {
	return Update(ctx, c, "VerificationResult", id, data)
}

func (c *Client) CreateVerificationResult(ctx context.Context, data *fhirmodel.VerificationResult) (*fhirmodel.VerificationResult, error) {
	return Create(ctx, c, "VerificationResult", data)
}

func (c *Client) SearchVerificationResult(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.VerificationResult, error) {
	return Search[fhirmodel.VerificationResult](ctx, c, "VerificationResult", params)
}

func (c *Client) ReadVisionPrescription(ctx context.Context, id string) (*fhirmodel.VisionPrescription, error) {
	return Read[fhirmodel.VisionPrescription](ctx, c, "VisionPrescription", id)
}

func (c *Client) UpdateVisionPrescription(ctx context.Context, id string, data *fhirmodel.VisionPrescription) (*fhirmodel.VisionPrescription, error) {
	return Update(ctx, c, "VisionPrescription", id, data)
}

func (c *Client) CreateVisionPrescription(ctx context.Context, data *fhirmodel.VisionPrescription) (*fhirmodel.VisionPrescription, error) {
	return Create(ctx, c, "VisionPrescription", data)
}

func (c *Client) SearchVisionPrescription(ctx context.Context, params SearchParams) (*fhirmodel.Bundle, []*fhirmodel.VisionPrescription, error) {
	return Search[fhirmodel.VisionPrescription](ctx, c, "VisionPrescription", params)
}
