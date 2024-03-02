package hprFns

import (
	"strconv"

	"github.com/MaSTeR2W/validator/constraint"
)

func ParseIntOrUintFromStr[T constraint.Ints | constraint.Uints](s string) (T, error) {
	var v T

	switch any(v).(type) {
	case int, int8, int16, int32, int64:
		var num, err = strconv.ParseInt(s, 10, GetBitSize(v))
		return T(num), err
	case uint, uint8, uint16, uint32, uint64:
		var num, err = strconv.ParseUint(s, 10, GetBitSize(v))
		return T(num), err
	default:
		panic("unexpected integer type")
	}
}

func ParseIntFromStr[T constraint.Ints](s string) (T, error) {
	var v T
	var num, err = strconv.ParseInt(s, 10, GetBitSize(v))
	return T(num), err
}

func ParseUnitFromStr[T constraint.Uints](s string) (T, error) {
	var v T
	var num, err = strconv.ParseUint(s, 10, GetBitSize(v))
	return T(num), err
}

func GetBitSize[T constraint.Ints | constraint.Uints](n T) int {
	switch any(n).(type) {
	case int, uint:
		return strconv.IntSize
	case int8, uint8:
		return 8
	case int16, uint16:
		return 16
	case int32, uint32:
		return 32
	case int64, uint64:
		return 64
	default:
		panic("unexpected integer type")
	}
}
