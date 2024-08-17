package models

// SecretString is a custom type that represents a secret string.
type SecretString string

type FunctionInputs struct {
	StructuralObjectTypes []string     `json:"structural_object_types" validate:"required"`
	MechanicalObjectTypes []string     `json:"mechanical_object_types" validate:"required"`
	SecretMessage         SecretString `json:"secret_message" validate:"required"`
}

type JSONSchemaGenerator interface {
	Generate(v interface{}, mode string) (map[string]interface{}, error)
}
