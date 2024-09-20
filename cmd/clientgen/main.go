package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/dave/jennifer/jen"
	fhirmodel "github.com/samply/golang-fhir-models/fhir-models/fhir"
)

func main() {
	fhirDefinitionZip, err := os.CreateTemp("", "fhir-definitions.zip")
	if err != nil {
		panic(err)
	}
	defer os.Remove(fhirDefinitionZip.Name())

	fhirDefinitionResp, err := http.Get("https://www.hl7.org/fhir/R4/definitions.json.zip")
	if err != nil {
		panic(err)
	}
	defer fhirDefinitionResp.Body.Close()
	if _, err := io.Copy(fhirDefinitionZip, fhirDefinitionResp.Body); err != nil {
		panic(err)
	}

	fhirDefinitionZipReader, err := zip.OpenReader(fhirDefinitionZip.Name())
	if err != nil {
		panic(err)
	}
	defer fhirDefinitionZipReader.Close()

	fhirResourceDefinition, err := fhirDefinitionZipReader.Open("profiles-resources.json")
	if err != nil {
		panic(err)
	}
	defer fhirResourceDefinition.Close()

	var fhirResourceDefinitionBundle fhirmodel.Bundle
	if err := json.NewDecoder(fhirResourceDefinition).Decode(&fhirResourceDefinitionBundle); err != nil {
		panic(err)
	}

	var capabilityStatement fhirmodel.CapabilityStatement
	if err := json.Unmarshal(fhirResourceDefinitionBundle.Entry[0].Resource, &capabilityStatement); err != nil {
		panic(err)
	}

	generatedSource := jen.NewFile("fhir")
	fhirModelPackage := "github.com/samply/golang-fhir-models/fhir-models/fhir"
	generatedSource.ImportAlias(fhirModelPackage, "fhirmodel")

	for _, resource := range capabilityStatement.Rest[0].Resource {
		fhirType := resource.Type.Code()

		for _, interaction := range resource.Interaction {
			switch interaction.Code {
			case fhirmodel.TypeRestfulInteractionCreate:
				generatedSource.Func().Params(jen.Id("c").Op("*").Id("Client")).Id(fmt.Sprintf("Create%s", fhirType)).Params(
					jen.Id("ctx").Qual("context", "Context"),
					jen.Id("data").Op("*").Qual(fhirModelPackage, fhirType),
				).Params(
					jen.Op("*").Qual(fhirModelPackage, fhirType),
					jen.Error(),
				).Block(
					jen.Return(
						jen.Id("Create").Call(
							jen.Id("ctx"),
							jen.Id("c"),
							jen.Lit(fhirType),
							jen.Id("data"),
						),
					),
				).Line()

			case fhirmodel.TypeRestfulInteractionRead:
				generatedSource.Func().Params(jen.Id("c").Op("*").Id("Client")).Id(fmt.Sprintf("Read%s", fhirType)).Params(
					jen.Id("ctx").Qual("context", "Context"),
					jen.Id("id").String(),
				).Params(
					jen.Op("*").Qual(fhirModelPackage, fhirType),
					jen.Error(),
				).Block(
					jen.Return(
						jen.Id("Read").
							Types(
								jen.Qual(fhirModelPackage, fhirType),
							).
							Call(
								jen.Id("ctx"),
								jen.Id("c"),
								jen.Lit(fhirType),
								jen.Id("id"),
							),
					),
				).Line()
			case fhirmodel.TypeRestfulInteractionUpdate:
				generatedSource.Func().Params(jen.Id("c").Op("*").Id("Client")).Id(fmt.Sprintf("Update%s", fhirType)).Params(
					jen.Id("ctx").Qual("context", "Context"),
					jen.Id("id").String(),
					jen.Id("data").Op("*").Qual(fhirModelPackage, fhirType),
				).Params(
					jen.Op("*").Qual(fhirModelPackage, fhirType),
					jen.Error(),
				).Block(
					jen.Return(
						jen.Id("Update").Call(
							jen.Id("ctx"),
							jen.Id("c"),
							jen.Lit(fhirType),
							jen.Id("id"),
							jen.Id("data"),
						),
					),
				).Line()
			case fhirmodel.TypeRestfulInteractionSearchType:
				generatedSource.Func().Params(jen.Id("c").Op("*").Id("Client")).Id(fmt.Sprintf("Search%s", fhirType)).Params(
					jen.Id("ctx").Qual("context", "Context"),
					jen.Id("params").Id("SearchParams"),
				).Params(
					jen.Op("*").Qual(fhirModelPackage, "Bundle"),
					jen.Op("[]").Op("*").Qual(fhirModelPackage, fhirType),
					jen.Error(),
				).Block(
					jen.Return(
						jen.Id("Search").
							Types(
								jen.Qual(fhirModelPackage, fhirType),
							).
							Call(
								jen.Id("ctx"),
								jen.Id("c"),
								jen.Lit(fhirType),
								jen.Id("params"),
							),
					),
				).Line()
			}
		}
	}

	generatedFile, err := os.Create("generated.go")
	if err != nil {
		panic(err)
	}
	defer generatedFile.Close()

	if err := generatedSource.Render(generatedFile); err != nil {
		panic(err)
	}
}
