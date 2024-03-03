package validators

import (
	"math"
	"strconv"

	"github.com/MaSTeR2W/validator/constraint"
	"github.com/MaSTeR2W/validator/errors"
	"github.com/MaSTeR2W/validator/types"
)

type Integer[T constraint.Ints | constraint.Uints] struct {
	Field   string
	NilAble bool
	Min     int64
	Max     int64
}

func (i *Integer[T]) GetField() string {
	return i.Field
}

func (i *Integer[T]) Validate(v any, path []any, lang string) (T, error) {

	// fV: float value
	// json.marshal convert any number to float64
	var iV int64
	var fV, isFloat = v.(float64)

	if !isFloat {
		// consider multipart request
		strV, isString := v.(string)

		if !isString {
			return 0, &types.ValidationErr{
				Field:   i.Field,
				Path:    path,
				Value:   types.Omit,
				Message: errors.InvalidDataType("integer", v, lang),
			}
		}

		var err error
		if iV, err = strconv.ParseInt(strV, 10, 64); err != nil {
			return 0, &types.ValidationErr{
				Field:   i.Field,
				Path:    path,
				Value:   types.Omit,
				Message: errors.InvalidDataType("integer", v, lang),
			}
		}
	} else {
		if fV != math.Trunc(fV) {
			return 0, &types.ValidationErr{
				Field:   i.Field,
				Path:    path,
				Value:   types.Omit,
				Message: errors.InvalidDataType("integer", v, lang),
			}
		}
		iV = int64(fV)
	}

	// iV: integer value

	if iV < i.Min {

		return 0, &types.ValidationErr{
			Field:   i.Field,
			Path:    path,
			Value:   iV,
			Message: smallIntErr(i.Min, lang),
		}
	}

	if iV > i.Max {
		return 0, &types.ValidationErr{
			Field:   i.Field,
			Path:    path,
			Value:   iV,
			Message: bigIntErr(i.Max, lang),
		}
	}

	return T(iV), nil
}

type NilInteger[T constraint.Ints | constraint.Uints] struct {
	Field string
	Min   int64
	Max   int64
}

func (i *NilInteger[T]) GetField() string {
	return i.Field
}

func (i *NilInteger[T]) Validate(v any, path []any, lang string) (*T, error) {

	if v == nil {
		return nil, nil
	}
	// fV: float value
	// json.marshal convert any number to float64
	var iV int64
	var fV, isFloat = v.(float64)

	if !isFloat {
		// consider multipart request
		strV, isString := v.(string)

		if !isString {
			return nil, &types.ValidationErr{
				Field:   i.Field,
				Path:    path,
				Value:   types.Omit,
				Message: errors.InvalidDataType("integer", v, lang),
			}
		}

		var err error
		if iV, err = strconv.ParseInt(strV, 10, 64); err != nil {
			return nil, &types.ValidationErr{
				Field:   i.Field,
				Path:    path,
				Value:   types.Omit,
				Message: errors.InvalidDataType("integer", v, lang),
			}
		}
	} else {
		if fV != math.Trunc(fV) {
			return nil, &types.ValidationErr{
				Field:   i.Field,
				Path:    path,
				Value:   types.Omit,
				Message: errors.InvalidDataType("integer", v, lang),
			}
		}
		iV = int64(fV)
	}

	// iV: integer value

	if iV < i.Min {

		return nil, &types.ValidationErr{
			Field:   i.Field,
			Path:    path,
			Value:   iV,
			Message: smallIntErr(i.Min, lang),
		}
	}

	if iV > i.Max {
		return nil, &types.ValidationErr{
			Field:   i.Field,
			Path:    path,
			Value:   iV,
			Message: bigIntErr(i.Max, lang),
		}
	}
	var tV = T(iV)
	return &tV, nil
}

func smallIntErr(exp int64, lang string) string {
	var sExp = strconv.FormatInt(exp, 10)

	if lang == "ar" {
		return "يجب أن تكون القيمة أكبر من أو تساوي " + sExp
	}

	return "Value should greater than or equal " + sExp
}

func bigIntErr(exp int64, lang string) string {
	var sExp = strconv.FormatInt(exp, 10)

	if lang == "ar" {
		return "يجب أن تكون القيمة أصغر من أو تساوي " + sExp
	}

	return "Value should be less than or equal " + sExp
}
