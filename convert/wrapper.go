package convert

import (
	"reflect"
	"strconv"
)

type Value struct {
	r interface{}
}

func Wrap(v interface{}) Value {
	return Value{v}
}

func isInt(k reflect.Kind) bool {
	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	default:
		return false
	}
}

func isUint(k reflect.Kind) bool {
	switch k {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return true
	default:
		return false
	}
}

func isFloat(k reflect.Kind) bool {
	switch k {
	case reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

func isComplex(k reflect.Kind) bool {
	switch k {
	case reflect.Complex64, reflect.Complex128:
		return true
	default:
		return false
	}
}

func (v Value) Interface() interface{} {
	return v.r
}

func (v Value) String() string {
	val := reflect.ValueOf(v.r)
	k := val.Kind()

	if isInt(k) {
		return strconv.FormatInt(val.Int(), 10)
	}
	if isUint(k) {
		return strconv.FormatUint(val.Uint(), 10)
	}
	if isFloat(k) {
		return strconv.FormatFloat(val.Float(), 'g', -1, 64)
	}
	if isComplex(k) {
		return strconv.FormatComplex(val.Complex(), 'g', -1, 128)
	}
	switch k {
	case reflect.Bool:
		if val.Bool() {
			return "true"
		}
		return "false"
	default:
		// TODO
		return val.String()
	}
}
