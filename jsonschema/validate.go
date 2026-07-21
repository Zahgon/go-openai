package jsonschema

func CollectDefs(def Definition) map[string]Definition { _ = "STUB: not implemented"; return nil }

func collectDefsRecursive(def Definition, result map[string]Definition, prefix string) {
	_ = "STUB: not implemented"
	return
}

func VerifySchemaAndUnmarshal(schema Definition, content []byte, v any) error {
	_ = "STUB: not implemented"
	return nil
}

type validateArgs struct {
	Defs map[string]Definition
}

type ValidateOption func(*validateArgs)

func WithDefs(defs map[string]Definition) ValidateOption {
	_ = "STUB: not implemented"
	return *new(ValidateOption)
}

func Validate(schema Definition, data any, opts ...ValidateOption) bool {
	_ = "STUB: not implemented"
	return false
}

func validateObject(schema Definition, data any, defs map[string]Definition) bool {
	_ = "STUB: not implemented"
	return false
}

func validateArray(schema Definition, data any, defs map[string]Definition) bool {
	_ = "STUB: not implemented"
	return false
}

func contains[S ~[]E, E comparable](s S, v E) bool { _ = "STUB: not implemented"; return false }
