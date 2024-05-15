package validators

import (
	"strconv"

	"github.com/MaSTeR2W/validator/constraint"
	"github.com/MaSTeR2W/validator/errors"
	"github.com/MaSTeR2W/validator/types"
)

type Float[T constraint.Floats] struct {
	Field string
	Min   float64
	Max   float64
}

func (i *Float[T]) GetField() string {
	return i.Field
}

func (i *Float[T]) Validate(v any, path []any, lang string) (T, error) {

	if v == nil {
		return 0, &types.ValidationErr{
			Field:   i.Field,
			Value:   types.Omit,
			Path:    path,
			Message: errors.RequiredFieldErr(lang),
		}
	}

	// fV: float value
	// json.marshal convert any number to float64
	var fV, isFloat = v.(float64)

	if !isFloat {
		// consider multipart request
		strV, isString := v.(string)

		if !isString {
			return 0, &types.ValidationErr{
				Field:   i.Field,
				Path:    path,
				Value:   types.Omit,
				Message: errors.InvalidDataType("float", v, lang),
			}
		}

		var err error
		if fV, err = strconv.ParseFloat(strV, 64); err != nil {
			return 0, &types.ValidationErr{
				Field:   i.Field,
				Path:    path,
				Value:   types.Omit,
				Message: errors.InvalidDataType("integer", v, lang),
			}
		}
	}

	// fV: integer value

	if fV < i.Min {

		return 0, &types.ValidationErr{
			Field:   i.Field,
			Path:    path,
			Value:   fV,
			Message: smallFloatErr(i.Min, lang),
		}
	}

	if fV > i.Max {
		return 0, &types.ValidationErr{
			Field:   i.Field,
			Path:    path,
			Value:   fV,
			Message: bigFloatErr(i.Max, lang),
		}
	}

	return T(fV), nil
}

type NilFloat[T constraint.Floats] struct {
	Field string
	Min   float64
	Max   float64
}

func (i *NilFloat[T]) GetField() string {
	return i.Field
}

func (i *NilFloat[T]) Validate(v any, path []any, lang string) (*T, error) {
	if v == nil {
		return nil, &types.ValidationErr{
			Field:   i.Field,
			Value:   types.Omit,
			Path:    path,
			Message: errors.RequiredFieldErr(lang),
		}
	}

	// fV: float value
	// json.marshal convert any number to float64
	var fV, isFloat = v.(float64)

	if !isFloat {
		// consider multipart request
		strV, isString := v.(string)

		if !isString {
			return nil, &types.ValidationErr{
				Field:   i.Field,
				Path:    path,
				Value:   types.Omit,
				Message: errors.InvalidDataType("float", v, lang),
			}
		}

		var err error
		if fV, err = strconv.ParseFloat(strV, 64); err != nil {
			return nil, &types.ValidationErr{
				Field:   i.Field,
				Path:    path,
				Value:   types.Omit,
				Message: errors.InvalidDataType("integer", v, lang),
			}
		}
	}

	// fV: integer value

	if fV < i.Min {

		return nil, &types.ValidationErr{
			Field:   i.Field,
			Path:    path,
			Value:   fV,
			Message: smallFloatErr(i.Min, lang),
		}
	}

	if fV > i.Max {
		return nil, &types.ValidationErr{
			Field:   i.Field,
			Path:    path,
			Value:   fV,
			Message: bigFloatErr(i.Max, lang),
		}
	}

	var pFV = T(fV)

	return &pFV, nil
}

func smallFloatErr(exp float64, lang string) string {
	var sExp = strconv.FormatFloat(exp, 'f', -1, 64)

	if lang == "ar" {
		return "يجب أن تكون القيمة أكبر من أو تساوي " + sExp
	}

	return "Value should greater than or equal " + sExp
}

func bigFloatErr(exp float64, lang string) string {
	var sExp = strconv.FormatFloat(exp, 'f', -1, 64)

	if lang == "ar" {
		return "يجب أن تكون القيمة أصغر من أو تساوي " + sExp
	}

	return "Value should be less than or equal " + sExp
}
