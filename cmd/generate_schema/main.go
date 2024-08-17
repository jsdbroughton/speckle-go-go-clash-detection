package main

import (
	"encoding/json"
	"fmt"
	"github.com/jsdbroughton/speckle-go-go-clash-detection/internal/automate"
	"github.com/jsdbroughton/speckle-go-go-clash-detection/internal/models"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: generate_schema <output file path>")
		os.Exit(1)
	}

	outputFilePath := os.Args[1]

	ags := &automate.AutomateGenerateJsonSchema{
		SchemaDialect: "http://json-schema.org/draft-07/schema#",
	}

	schema, err := ags.Generate(models.FunctionInputs{}, "validation")
	if err != nil {
		fmt.Println("Error generating schema:", err)
		os.Exit(1)
	}

	data, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling schema to JSON:", err)
		os.Exit(1)
	}

	err = os.WriteFile(outputFilePath, data, 0644)
	if err != nil {
		fmt.Println("Error writing schema to file:", err)
		os.Exit(1)
	}

	fmt.Println("Schema generated successfully:", outputFilePath)
}
