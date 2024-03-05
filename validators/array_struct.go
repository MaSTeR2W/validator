package validators

import (
	"encoding/json"
	"slices"

	"github.com/MaSTeR2W/validator/errors"
	"github.com/MaSTeR2W/validator/types"
)

type Array_Struct[T any] struct {
	Field     string
	NilAble   bool
	MinLength int
	MaxLength int
	Validator types.StructValidator[T]
}

func (a *Array_Struct[T]) GetField() string {
	return a.Field
}

func (a *Array_Struct[T]) Validate(v any, path []any, lang string) ([]*T, error) {
	var arrAny []any

	switch arr := v.(type) {
	case nil:
		if a.NilAble {
			return nil, nil
		}

		return nil, &types.ValidationErr{
			Field:   a.Field,
			Value:   types.Omit,
			Path:    path,
			Message: errors.RequiredFieldErr(lang),
		}

	case []any:
		arrAny = arr
	case string:
		var err = json.Unmarshal([]byte(arr), &arrAny)
		if err != nil {
			return nil, &types.ValidationErr{
				Field:   a.Field,
				Value:   arr,
				Path:    path,
				Message: errors.InvalidJSONArr(lang),
			}
		}

	default:
		return nil, &types.ValidationErr{
			Field:   a.Field,
			Value:   v,
			Path:    path,
			Message: errors.InvalidDataType("[]map[string]any", v, lang),
		}
	}

	var l = len(arrAny)

	if l < a.MinLength {
		return nil, &types.ValidationErr{
			Field:   a.Field,
			Value:   v,
			Path:    path,
			Message: errors.ShortArrErr(a.MinLength, l, lang),
		}
	}

	if a.MaxLength > -1 && l > a.MaxLength {
		return nil, &types.ValidationErr{
			Field:   a.Field,
			Value:   v,
			Path:    path,
			Message: errors.LongArrErr(a.MinLength, l, lang),
		}
	}

	var arrOfT = make([]*T, 0, l)

	for i, e := range arrAny {

		var t, err = a.Validator(e, a.Field, append(slices.Clone(path), i), lang)

		if err != nil {
			return nil, err
		}

		arrOfT = append(arrOfT, t)
	}

	return arrOfT, nil
}
