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
// Liquid semantics: every value except nil and false is truthy.
//
// This is the single truthiness judgment shared by the {% if %} family of
// tags, the "not" expression operator, and filters such as default and where,
// so that each AST node evaluates truthiness against the same boundary.
func IsTruthy(value any) bool {
	return value != nil && value != false
}
