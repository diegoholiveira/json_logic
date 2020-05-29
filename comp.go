package jsonlogic

import (
	"reflect"
)

func less(a, b interface{}) bool {
	// comparison to a nil value is falsy
	if a == nil || b == nil {
		return false
	}

	if isNumber(a) || isNumber(b) {
		return toNumber(b) > toNumber(a)
	}

	return toString(b) > toString(a)
}

func hardEquals(a, b interface{}) bool {
	ra := reflect.ValueOf(a).Kind()
	rb := reflect.ValueOf(b).Kind()

	if ra != rb {
		return false
	}

	return equals(a, b)
}

func equals(a, b interface{}) bool {
	// comparison to a nil value is falsy
	if a == nil || b == nil {
		return false
	}

	if isNumber(a) {
		return toNumber(a) == toNumber(b)
	}

	if isBool(a) {
		return isTrue(a) == isTrue(b)
	}

	return toString(a) == toString(b)
}
