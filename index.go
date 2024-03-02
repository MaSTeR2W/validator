package validators

import (
	"github.com/MaSTeR2W/validator/validators"
)

func Bool(field string, nilAble bool) validators.Bool {
	return validators.Bool{
		Field:   field,
		NilAble: nilAble,
	}
}
