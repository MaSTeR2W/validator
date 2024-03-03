package validators

import (
	"slices"
	"strconv"

	"github.com/MaSTeR2W/validator/errors"
	"github.com/MaSTeR2W/validator/types"
)

type Array_Str struct {
	Field     string
	NilAble   bool
	MinLength int
	MaxLength int
	Validator *types.Validator[*string]
}

func (a *Array_Str) GetField() string {
	return a.Field
}

func (a *Array_Str) Validate(v any, path []any, lang string) ([]string, error) {
	if v == nil {
		if a.NilAble {
			return nil, nil
		}

		return nil, &types.ValidationErr{
			Field:   a.Field,
			Value:   v,
			Path:    path,
			Message: errors.InvalidDataType("[]string", v, lang),
		}
	}

	var arrAny []any
	var ok bool
	if arrAny, ok = v.([]any); !ok {
		return nil, &types.ValidationErr{
			Field:   a.Field,
			Value:   v,
			Path:    path,
			Message: errors.InvalidDataType("[]string", v, lang),
		}
	}

	var l = len(arrAny)

	if l < a.MinLength {
		return nil, &types.ValidationErr{
			Field:   a.Field,
			Value:   v,
			Path:    path,
			Message: shortArrErr(a.MinLength, l, lang),
		}

	}

	if a.MaxLength > -1 && l > a.MaxLength {
		return nil, &types.ValidationErr{
			Field:   a.Field,
			Value:   v,
			Path:    path,
			Message: longArrErr(a.MinLength, l, lang),
		}
	}

	var arrStr = make([]string, l)
	for i, e := range arrAny {
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

	if a.Validator != nil {
		var err error
		for i, e := range arrStr {
			var curPath = append(slices.Clone(path), i)
			_, err = (*a.Validator).Validate(e, curPath, lang)
			if err != nil {
				return nil, err
			}
		}
	}

	return arrStr, nil
}

func shortArrErr(exp int, got int, lang string) string {
	var sExp, sGot = strconv.Itoa(exp), strconv.Itoa(got)

	if lang == "ar" {
		return "يجب زيادة عدد عناصر المصفوفة إلى " + sExp + " أو أكثر (عدد عناصر المصفوفة حاليا " + sGot + ")"
	}
	return "Should increase the number of elements of array to " + sExp + " or more (the current number of element is " + sGot + ")"
}

func longArrErr(exp int, got int, lang string) string {
	var sExp, sGot = strconv.Itoa(exp), strconv.Itoa(got)

	if lang == "ar" {
		return "يجب إنقاص عدد عناصر المصفوفة إلى " + sExp + " أو أقل (عدد عناصر المصفوفة حاليا " + sGot + ")"
	}
	return "Should decrease the number of elements of array to " + sExp + " or less (the current number of element is " + sGot + ")"
}
