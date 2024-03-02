package validators

import (
	"github.com/MaSTeR2W/validator/errors"
	"github.com/MaSTeR2W/validator/types"
)

type Bool struct {
	Field  string
	NotNil bool
}

func (b *Bool) GetField() string {
	return b.Field
}

func (b *Bool) Validate(v any, path []any, lang string) (bool, error) {

	if v == nil {
		if b.NotNil {
			return false, &types.ValidationErr{
				Field:   b.Field,
				Path:    path,
				Value:   types.Omit,
				Message: errors.InvalidDataType("bolean", v, lang),
			}
		}
		return false, nil
	}

	var _, ok = v.(bool)

	if ok {
		return false, nil
	}

	var strBool string
	strBool, ok = v.(string)

	if ok {
		if strBool == "true" {
			return true, nil
		}

		if strBool == "false" {
			return false, nil
		}
	}
	return false, &types.ValidationErr{
		Field:   b.Field,
		Path:    path,
		Value:   v,
		Message: errors.InvalidDataType("boolean", v, lang),
	}

}
