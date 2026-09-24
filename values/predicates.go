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

// IsTruthy returns a bool indicating whether the value is true according to Liquid semantics:
// every value except nil and false is truthy.
//
// This is the single truthiness judgment shared by the expression language
// (and/or/not) and the control-flow tags (if/unless), so that every AST node
// that branches on a value agrees on the outcome.
func IsTruthy(value any) bool {
	return ValueOf(value).Test()
}
