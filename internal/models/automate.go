package models

type FunctionInputs struct {
	StructuralObjectTypes []string `json:"structural_object_types" validate:"required"`
	MechanicalObjectTypes []string `json:"mechanical_object_types" validate:"required"`
	SecretMessage         string   `json:"whisper_message" validate:"required" secret:"true"`
}

type JSONSchemaGenerator interface {
	Generate(v interface{}, mode string) (map[string]interface{}, error)
}
