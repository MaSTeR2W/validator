package validators

import (
	"slices"
	"strconv"
	"strings"
)

type String struct {
	Field     string
	NotNil    bool
	MinLength int
	MaxLength int
	Enum      []string
}

func (s *String) GetField() string {
	return s.Field
}

func (s *String) Validate(v any, lang string) error {

	if v == nil {
		if s.NotNil {
			return &ValidationErr{
				Field:   s.Field,
				Value:   null,
				Message: invalidDataType("string", v, lang),
			}
		}
		return nil
	}

	var vStr, ok = v.(string)

	if !ok {
		return &ValidationErr{
			Field:   s.Field,
			Value:   v,
			Message: invalidDataType("string", v, lang),
		}
	}

	if s.Enum != nil {
		if !slices.Contains(s.Enum, vStr) {
			return &ValidationErr{
				Field:   s.Field,
				Value:   vStr,
				Message: unexpectedValue(s.Enum, vStr, lang),
			}
		}
		return nil
	}

	var l = len(vStr)
	if l < s.MinLength {
		return &ValidationErr{
			Field:   s.Field,
			Value:   vStr,
			Message: shortStringErr(s.MinLength, l, lang),
		}
	}

	if s.MaxLength > 0 && l > s.MaxLength {
		return &ValidationErr{
			Field:   s.Field,
			Value:   vStr,
			Message: lognStringErr(s.MaxLength, l, lang),
		}
	}

	return nil
}

func shortStringErr(exp, got int, lang string) string {
	var sExp, sGot = strconv.Itoa(exp), strconv.Itoa(got)

	if lang == "ar" {
		return "يجب إطالة هذا النص إلى " + sExp + "من الحروف أو أكثر (أنت تستخدم حاليا " + sGot + " من الحروف)"
	}

	return "Should lengthen this text to " + sExp + " characters or more (you are currently using " + sGot + " characters)"
}

func lognStringErr(exp, got int, lang string) string {
	var sExp, sGot = strconv.Itoa(exp), strconv.Itoa(got)

	if lang == "ar" {
		return "يجب تقصير هذا النص إلى " + sExp + " من الحروف أو أقل (أنت حاليا تستخدم " + sGot + " من الحروف)"
	}

	return "Should shorten this text to " + sExp + " characters (you are currently using " + sGot + " characters)"
}

func unexpectedValue(exp []string, got string, lang string) string {
	var enum = strings.Join(exp, ", ")

	if lang == "ar" {
		return "يجب أن تكون القيمة واحدة من: (" + enum + ") وليس (" + got + ")"
	}
	return "Value should be one of: (" + enum + "), not (" + got + ")"
}
