package convert

import (
	"math"
	"reflect"
)

func (v Value) Int8() int8 {
	val := reflect.ValueOf(v.r)
	k := val.Kind()

	if isInt(k) {
		return int8(val.Int())
	}
	if isUint(k) {
		return int8(val.Uint())
	}
	if isFloat(k) {
		return int8(int64(val.Float()))
	}
	if isComplex(k) {
		return int8(int64(real(val.Complex())))
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

func (v Value) Int16() int16 {
	val := reflect.ValueOf(v.r)
	k := val.Kind()

	if isInt(k) {
		return int16(val.Int())
	}
	if isUint(k) {
		return int16(val.Uint())
	}
	if isFloat(k) {
		return int16(int64(val.Float()))
	}
	if isComplex(k) {
		return int16(int64(real(val.Complex())))
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

func (v Value) Int32() int32 {
	val := reflect.ValueOf(v.r)
	k := val.Kind()

	if isInt(k) {
		return int32(val.Int() & math.MaxInt32)
	}
	if isUint(k) {
		return int32(val.Uint() & math.MaxInt32)
	}
	if isFloat(k) {
		return int32(int64(val.Float()) & math.MaxInt32)
	}
	if isComplex(k) {
		return int32(int64(real(val.Complex())) & math.MaxInt32)
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

func (v Value) Int64() int64 {
	val := reflect.ValueOf(v.r)
	k := val.Kind()

	if isInt(k) {
		return val.Int()
	}
	if isUint(k) {
		return int64(val.Uint())
	}
	if isFloat(k) {
		return int64(val.Float())
	}
	if isComplex(k) {
		return int64(real(val.Complex()))
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

func (v Value) Uint8() uint8 {
	val := reflect.ValueOf(v.r)
	k := val.Kind()

	if isInt(k) {
		return uint8(val.Int())
	}
	if isUint(k) {
		return uint8(val.Uint())
	}
	if isFloat(k) {
		return uint8(int64(val.Float()))
	}
	if isComplex(k) {
		return uint8(int64(real(val.Complex())))
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

func (v Value) Uint16() uint16 {
	val := reflect.ValueOf(v.r)
	k := val.Kind()

	if isInt(k) {
		return uint16(val.Int())
	}
	if isUint(k) {
		return uint16(val.Uint())
	}
	if isFloat(k) {
		return uint16(int64(val.Float()))
	}
	if isComplex(k) {
		return uint16(int64(real(val.Complex())))
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

func (v Value) Uint32() uint32 {
	val := reflect.ValueOf(v.r)
	k := val.Kind()

	if isInt(k) {
		return uint32(val.Int())
	}
	if isUint(k) {
		return uint32(val.Uint())
	}
	if isFloat(k) {
		return uint32(int64(val.Float()))
	}
	if isComplex(k) {
		return uint32(int64(real(val.Complex())))
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

func (v Value) Uint64() uint64 {
	val := reflect.ValueOf(v.r)
	k := val.Kind()

	if isInt(k) {
		return uint64(val.Int())
	}
	if isUint(k) {
		return val.Uint()
	}
	if isFloat(k) {
		return uint64(val.Float())
	}
	if isComplex(k) {
		return uint64(real(val.Complex()))
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
