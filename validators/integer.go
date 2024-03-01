package validators

import (
	"math"
	"strconv"
)

type Integer struct {
	Field  string
	NotNil bool
	Min    int64
	Max    int64
}

func (i *Integer) GetField() string {
	return i.Field
}

func (i *Integer) Validate(v any, lang string) error {

	if v == nil {
		if i.NotNil {
			return &ValidationErr{
				Field:   i.Field,
				Value:   null,
				Message: invalidDataType("string", v, lang),
			}
		}
		return nil
	}
	// fV: float value
	// json.marshal convert any number to float64
	var iV int64
	var fV, isFloat = v.(float64)

	if !isFloat {
		// consider multipart request
		strV, isString := v.(string)

		if !isString {
			return &ValidationErr{
				Field:   i.Field,
				Value:   v,
				Message: invalidDataType("integer", v, lang),
			}
		}

		var err error
		if iV, err = strconv.ParseInt(strV, 10, 64); err != nil {
			return &ValidationErr{
				Field:   i.Field,
				Value:   v,
				Message: invalidDataType("integer", v, lang),
			}
		}
	} else {
		if fV != math.Trunc(fV) {
			return &ValidationErr{
				Field:   i.Field,
				Value:   v,
				Message: invalidDataType("integer", v, lang),
			}
		}
		iV = int64(fV)
	}

	// iV: integer value

	if iV < i.Min {
		return &ValidationErr{
			Field:   i.Field,
			Value:   iV,
			Message: smallIntErr(i.Min, lang),
		}
	}

	if iV > i.Max {
		return &ValidationErr{
			Field:   i.Field,
			Value:   iV,
			Message: bigIntErr(i.Max, lang),
		}
	}

	return nil
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
