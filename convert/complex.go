package convert

import (
	"reflect"
)

func (v Value) Complex64() complex64 {
	val := reflect.ValueOf(v.r)
	k := val.Kind()

	if isInt(k) {
		return complex(float32(val.Int()), 0)
	}
	if isUint(k) {
		return complex(float32(val.Uint()), 0)
	}
	if isFloat(k) {
		return complex(float32(val.Float()), 0)
	}
	if isComplex(k) {
		return complex64(val.Complex())
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

func (v Value) Complex128() complex128 {
	return complex128(v.Complex64())
}
