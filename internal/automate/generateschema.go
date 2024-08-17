package automate

import (
	"reflect"
	"strings"
)

type AutomateGenerateJsonSchema struct {
	SchemaDialect string
}

func (ags *AutomateGenerateJsonSchema) Generate(v interface{}, mode string) (map[string]interface{}, error) {
	schema := map[string]interface{}{
		"$schema":    ags.SchemaDialect,
		"title":      reflect.TypeOf(v).Name(),
		"type":       "object",
		"properties": map[string]interface{}{},
		"required":   []string{},
	}

	t := reflect.TypeOf(v)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldName := field.Tag.Get("json")
		if fieldName == "" {
			fieldName = strings.ToLower(field.Name)
		}

		fieldSchema := map[string]interface{}{
			"type": "string", // Default type; could be extended based on kind
		}

		// Handle array/slice types
		if field.Type.Kind() == reflect.Slice {
			fieldSchema["type"] = "array"
			fieldSchema["items"] = map[string]interface{}{
				"type": "string",
			}
		}

		// Check if the field is marked as a secret
		if _, ok := field.Tag.Lookup("secret"); ok {
			fieldSchema["writeOnly"] = true
		}

		schema["properties"].(map[string]interface{})[fieldName] = fieldSchema

		// Add to required fields if validate:"required" tag is present
		if strings.Contains(string(field.Tag), `validate:"required"`) {
			schema["required"] = append(schema["required"].([]string), fieldName)
		}
	}

	return schema, nil
}
