package validators

import (
	"github.com/MaSTeR2W/validator/errors"
	"github.com/MaSTeR2W/validator/types"
)

type Bool struct {
	Field   string
	NilAble bool
	BeTrue  bool
	BeFalse bool
}

func (b *Bool) GetField() string {
	return b.Field
}

func (b *Bool) Validate(v any, path []any, lang string) (bool, error) {

	var vB bool
	var ok bool
	if vB, ok = v.(bool); ok {

	} else if strBool, ok := v.(string); ok {

		if strBool == "true" {
			vB = true
		}

		if strBool == "false" {
			vB = false
		}
	} else {
		return false, &types.ValidationErr{
			Field:   b.Field,
			Path:    path,
			Value:   v,
			Message: errors.InvalidDataType("boolean", v, lang),
		}
	}

	if b.BeTrue && !vB {
		return false, &types.ValidationErr{
			Field:   b.Field,
			Path:    path,
			Value:   v,
			Message: notTrueErr(lang),
		}
	}

	if b.BeFalse && vB {
		return false, &types.ValidationErr{
			Field:   b.Field,
			Path:    path,
			Value:   v,
			Message: notFalseErr(lang),
		}
	}

	return vB, nil
}

type NilBool struct {
	Field   string
	BeTrue  bool
	BeFalse bool
}

func (b *NilBool) GetField() string {
	return b.Field
}

func (b *NilBool) Validate(v any, path []any, lang string) (*bool, error) {

	if v == nil {
		return nil, nil
	}

	var vB bool
	var ok bool
	if vB, ok = v.(bool); ok {

	} else if strBool, ok := v.(string); ok {

		if strBool == "true" {
			vB = true
		}

		if strBool == "false" {
			vB = false
		}
	} else {
		return nil, &types.ValidationErr{
			Field:   b.Field,
			Path:    path,
			Value:   v,
			Message: errors.InvalidDataType("boolean", v, lang),
		}
	}

	if b.BeTrue && !vB {
		return nil, &types.ValidationErr{
			Field:   b.Field,
			Path:    path,
			Value:   v,
			Message: notTrueErr(lang),
		}
	}

	if b.BeFalse && vB {
		return nil, &types.ValidationErr{
			Field:   b.Field,
			Path:    path,
			Value:   v,
			Message: notFalseErr(lang),
		}
	}

	return &vB, nil
}

func notTrueErr(lang string) string {
	if lang == "ar" {
		return "يجب أن تكون القيمة (true) في حين أنها ليست كذلك"
	}
	return "Value should be (true) while it is not"

}

func notFalseErr(lang string) string {
	if lang == "ar" {
		return "يجب أن تكون القيمة (false) في حين أنها ليست كذلك"
	}
	return "Value should be (false) while it is not"
}
