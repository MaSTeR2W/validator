package validators

import (
	"regexp"
	"strconv"
	"unicode"

	"github.com/MaSTeR2W/validator/errors"
	"github.com/MaSTeR2W/validator/types"
)

type Password struct {
	Field string
}

func (p *Password) GetField() string {
	return p.Field
}

var upRgp = regexp.MustCompile("[A-Z]")
var lwRgp = regexp.MustCompile("[a-z]")
var nmRgp = regexp.MustCompile("[0-9]")

func (p *Password) Validate(v any, path []any, lang string) (string, error) {
	var password string
	var ok bool
	if password, ok = v.(string); !ok {
		return "", &types.ValidationErr{
			Field:   p.Field,
			Path:    path,
			Value:   types.Omit,
			Message: errors.InvalidDataType("string", v, lang),
		}
	}

	var passLen = len(password)
	if passLen < 8 {
		return "", &types.ValidationErr{
			Field:   p.Field,
			Path:    path,
			Value:   types.Omit,
			Message: shortPasswordErr(passLen, lang),
		}
	}

	if passLen > 32 {
		return "", &types.ValidationErr{
			Field:   p.Field,
			Path:    path,
			Value:   types.Omit,
			Message: longPasswordErr(passLen, lang),
		}
	}

	if !upRgp.MatchString(password) {
		return "", &types.ValidationErr{
			Field:   p.Field,
			Path:    path,
			Value:   types.Omit,
			Message: missingCapitalLetterErr(lang),
		}
	}

	if !lwRgp.MatchString(password) {
		return "", &types.ValidationErr{
			Field:   p.Field,
			Path:    path,
			Value:   types.Omit,
			Message: missingLowercaseLetterErr(lang),
		}
	}

	if !nmRgp.MatchString(password) {
		return "", &types.ValidationErr{
			Field:   p.Field,
			Path:    path,
			Value:   types.Omit,
			Message: missingNumberErr(lang),
		}
	}

	var hasSymbol bool = false
	for _, r := range password {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			hasSymbol = true
			break
		}
	}
	if !hasSymbol {
		return "", &types.ValidationErr{
			Field:   p.Field,
			Path:    path,
			Value:   types.Omit,
			Message: missingSymbolErr(lang),
		}
	}
	return password, nil
}

func shortPasswordErr(got int, lang string) string {
	var sGot = strconv.Itoa(got)
	if lang == "ar" {
		return "يجب إطالة كلمة المرور هذه إلى 8 حروف أو أكثر (أنت حاليا تستخدم " + sGot + " من الحروف)"
	}
	return "Should lengthen this password to 8 characters or more (you are currently using " + sGot + " characters)"
}

func longPasswordErr(got int, lang string) string {
	var sGot = strconv.Itoa(got)
	if lang == "ar" {
		return "يجب تقصير كلمة المرور هذه إلى 32 حرفاً أو أقل (أنت حاليا تستخدم " + sGot + " من الحروف)"
	}
	return "Should shorten this password to 8 characters or less (you are currently using " + sGot + " characters)"
}

func missingCapitalLetterErr(lang string) string {
	if lang == "ar" {
		return "يجب أن تحتوي كلمة المرور على حرف كبير واحد على الأقل"
	}

	return "The password should contain at least one capital letter"
}

func missingLowercaseLetterErr(lang string) string {
	if lang == "ar" {
		return "يجب أن تحتوي كلمة المرور على حرف صغير واحد على الأقل"
	}

	return "The password should contain at least one lowercase letter"
}

func missingNumberErr(lang string) string {
	if lang == "ar" {
		return "يجب أن تحتوي كلمة المرور على رقم واحد على الأقل"
	}

	return "The password should contain at least one number"
}

func missingSymbolErr(lang string) string {
	if lang == "ar" {
		return "يجب أن تحتوي كلمة المرور على رمز واحد على الأقل"
	}

	return "The password should contain at least one symbol"
}
