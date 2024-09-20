package main

import (
	"fmt"
	"os"

	"github.com/dave/jennifer/jen"
)

func main() {
	generatedSource := jen.NewFile("fhir")
	fhirModelPackage := "github.com/samply/golang-fhir-models/fhir-models/fhir"
	generatedSource.ImportAlias(fhirModelPackage, "fhirmodel")

	fhirTypes := []string{
		"Account",
	}

	for _, fhirType := range fhirTypes {

		// create resource
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

		// read resource
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

		// update resource
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

		// search resource
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

	generatedFile, err := os.Create("generated.go")
	if err != nil {
		panic(err)
	}
	defer generatedFile.Close()

	if err := generatedSource.Render(generatedFile); err != nil {
		panic(err)
	}
}
