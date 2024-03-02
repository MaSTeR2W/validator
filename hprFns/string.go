package hprFns

import (
	"strconv"

	"github.com/MaSTeR2W/validator/constraint"
)

func FormatInt[T constraint.Ints](n T) string {

	switch n := any(n).(type) {
	case int:
		return strconv.FormatInt(int64(n), 10)
	case int8:
		return strconv.FormatInt(int64(n), 10)
	case int16:
		return strconv.FormatInt(int64(n), 10)
	case int32:
		return strconv.FormatInt(int64(n), 10)
	case int64:
		return strconv.FormatInt(n, 10)
	default:
		panic("unexpected integer type")
	}
}
