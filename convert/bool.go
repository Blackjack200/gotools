package convert

import "reflect"

func (v Value) Bool() bool {
	val := reflect.ValueOf(v.r)
	k := val.Kind()

	if isInt(k) {
		return val.Int() == 0
	}
	if isUint(k) {
		return val.Uint() == 0
	}
	if isFloat(k) {
		return val.Float() == 0
	}
	if isComplex(k) {
		return real(val.Complex()) == 0
	}

	switch k {
	case reflect.Bool:
		return true
	default:
		return false
	}
}
