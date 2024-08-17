package automate

import (
	"reflect"
	"strings"
)

type GenerateAutomateJsonSchema struct {
	SchemaDialect string
}

func (ags *GenerateAutomateJsonSchema) Generate(v interface{}, mode string) (map[string]interface{}, error) {
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

		fieldSchema := map[string]interface{}{}
		switch field.Type.Kind() {
		case reflect.String:
			fieldSchema["type"] = "string"
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fieldSchema["type"] = "integer"
		case reflect.Float32, reflect.Float64:
			fieldSchema["type"] = "number"
		case reflect.Bool:
			fieldSchema["type"] = "boolean"
		case reflect.Slice:
			fieldSchema["type"] = "array"
			fieldSchema["items"] = map[string]interface{}{
				"type": determineElementType(field.Type.Elem()),
			}
		default:
			fieldSchema["type"] = "string" // Fallback for unsupported types
		}

		// Check if the field is marked as a secret
		if _, ok := field.Tag.Lookup("secret"); ok {
			fieldSchema["writeOnly"] = true
		}

		schema["properties"].(map[string]interface{})[fieldName] = fieldSchema

		// Add to required fields if the "validate:'required'" tag is present
		if strings.Contains(string(field.Tag), `validate:"required"`) {
			schema["required"] = append(schema["required"].([]string), fieldName)
		}
	}

	return schema, nil
}

// Helper function to determine the JSON type for slice elements
func determineElementType(t reflect.Type) string {
	switch t.Kind() {
	case reflect.String:
		return "string"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return "integer"
	case reflect.Float32, reflect.Float64:
		return "number"
	case reflect.Bool:
		return "boolean"
	default:
		return "string" // Default fallback
	}
}
