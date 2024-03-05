package validators

import (
	"encoding/json"
	"slices"

	"github.com/MaSTeR2W/validator/errors"
	"github.com/MaSTeR2W/validator/types"
)

type Array_Str struct {
	Field     string
	NilAble   bool
	MinLength int
	MaxLength int
	Validator types.Validator[*string]
}

func (a *Array_Str) GetField() string {
	return a.Field
}

func (a *Array_Str) Validate(v any, path []any, lang string) ([]string, error) {

	var arrStr []string

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
		arrStr = make([]string, 0, len(arr))
		var ok bool
		for i, e := range arr {
			var strE string
			if strE, ok = e.(string); !ok {
				return nil, &types.ValidationErr{
					Field:   a.Field,
					Value:   v,
					Path:    append(path, i),
					Message: errors.InvalidDataType("string", e, lang),
				}
			}
			arrStr = append(arrStr, strE)
		}

	case string:
		var err = json.Unmarshal([]byte(arr), &arrStr)
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
			Message: errors.InvalidDataType("[]string", v, lang),
		}
	}

	var l = len(arrStr)

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

	if a.Validator != nil {
		var err error
		for i, e := range arrStr {
			var curPath = append(slices.Clone(path), i)
			_, err = a.Validator.Validate(e, curPath, lang)
			if err != nil {
				return nil, err
			}
		}
	}

	return arrStr, nil
}
