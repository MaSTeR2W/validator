package hprFns

import (
	"strconv"

	"github.com/MaSTeR2W/validator/constraint"
)

func ParseFloatFromString[T constraint.Floats](s string) (T, error) {
	var v T

	switch any(v).(type) {
	case float32:
		var num, err = strconv.ParseFloat(s, 32)
		return T(num), err
	case float64:
		var num, err = strconv.ParseFloat(s, 64)
		return T(num), err
	// impossible case
	default:
		panic("unexpected type")
	}

}
