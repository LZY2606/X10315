package values

import (
	"reflect"
)

// IsEmpty returns a bool indicating whether the value is empty according to Liquid semantics.
func IsEmpty(value any) bool {
	value = ToLiquid(value)
	if value == nil {
		return false
	}

	r := reflect.ValueOf(value)
	switch r.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return r.Len() == 0
	case reflect.Bool:
		return !r.Bool()
	default:
		return false
	}
}

// IsTruthy returns a bool indicating whether the value is truthy according to
// Liquid semantics: only nil and false are falsy. In particular, zero, empty
// strings, and empty arrays are truthy.
//
// This is the single definition of truthiness for the engine. It backs
// Value.Test, and it is consulted directly by the {% if %}/{% unless %}
// branches, the default and where filters, and boolean conversion, so that
// every AST node judges the same value the same way.
func IsTruthy(value any) bool {
	return value != nil && value != false
}
