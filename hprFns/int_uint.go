package hprFns

import "github.com/MaSTeR2W/validator/constraint"

func ParseIntOrUint[T constraint.Ints | constraint.Uints](f float64) (T, error) {
	var v T

	switch any(v).(type) {
	case int:

	case int8:
	case int16:
	case int32:
	case int64:
	case uint:
	case uint8:
	case uint16:
	case uint32:
	case uint64:
	}
}
