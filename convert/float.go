package convert

import (
	"reflect"
)

func (v Value) Float32() float32 {
	return float32(v.Float64())
}

func (v Value) Float64() float64 {
	val := reflect.ValueOf(v.r)
	k := val.Kind()

	if isInt(k) {
		return float64(val.Int())
	}
	if isUint(k) {
		return float64(val.Uint())
	}
	if isFloat(k) {
		return val.Float()
	}
	if isComplex(k) {
		return real(val.Complex())
	}

	switch k {
	case reflect.Bool:
		if val.Bool() {
			return 1
		}
		return 0
	default:
		return 0
	}
}
