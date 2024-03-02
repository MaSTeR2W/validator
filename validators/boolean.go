package validators

import (
	"github.com/MaSTeR2W/validator/errors"
	"github.com/MaSTeR2W/validator/types"
)

type Bool struct {
	Field   string
	NilAble bool
}

func (b *Bool) GetField() string {
	return b.Field
}

func (b *Bool) Validate(v any, path []any, lang string) (*bool, error) {

	if v == nil {
		if b.NilAble {
			return nil, &types.ValidationErr{
				Field:   b.Field,
				Path:    path,
				Value:   types.Omit,
				Message: errors.InvalidDataType("bolean", v, lang),
			}
		}
		return nil, nil
	}

	if vB, ok := v.(bool); ok {

		return &vB, nil
	}

	if strBool, ok := v.(string); ok {
		var bo bool = true
		if strBool == "true" {
			return &bo, nil
		}

		if strBool == "false" {
			bo = false
			return &bo, nil
		}
	}

	return nil, &types.ValidationErr{
		Field:   b.Field,
		Path:    path,
		Value:   v,
		Message: errors.InvalidDataType("boolean", v, lang),
	}

}
