package jsonschema

import (
	"reflect"
)

type DataType string

const (
	Object  DataType = "object"
	Number  DataType = "number"
	Integer DataType = "integer"
	String  DataType = "string"
	Array   DataType = "array"
	Null    DataType = "null"
	Boolean DataType = "boolean"
)

type Definition struct {
	Type DataType `json:"type,omitempty"`

	Description string `json:"description,omitempty"`

	Enum []string `json:"enum,omitempty"`

	Properties map[string]Definition `json:"properties,omitempty"`

	Required []string `json:"required,omitempty"`

	Items *Definition `json:"items,omitempty"`

	AdditionalProperties any `json:"additionalProperties,omitempty"`

	Nullable bool `json:"nullable,omitempty"`

	Ref string `json:"$ref,omitempty"`

	Defs map[string]Definition `json:"$defs,omitempty"`
}

func (d *Definition) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *Definition) Unmarshal(content string, v any) error { _ = "STUB: not implemented"; return nil }

func GenerateSchemaForType(v any) (*Definition, error) { _ = "STUB: not implemented"; return nil, nil }

func reflectSchema(t reflect.Type, defs map[string]Definition) (*Definition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func reflectSchemaObject(t reflect.Type, defs map[string]Definition) (*Definition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func containsRef(def Definition, targetRef string) bool { _ = "STUB: not implemented"; return false }
