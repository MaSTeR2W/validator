package validators

import (
	"net"
	"strconv"
	"strings"

	"github.com/MaSTeR2W/validator/errors"
	"github.com/MaSTeR2W/validator/types"
)

type Email struct {
	Field   string
	CheckMx bool
}

func (e *Email) GetField() string {
	return e.Field
}

func (e *Email) Validate(v any, path []any, lang string) (string, error) {

	if v == nil {
		return "", &types.ValidationErr{
			Field:   e.Field,
			Path:    path,
			Value:   types.Omit,
			Message: errors.RequiredFieldErr(lang),
		}
	}

	var email, ok = v.(string)

	if !ok {
		return "", &types.ValidationErr{
			Field:   e.Field,
			Path:    path,
			Value:   types.Omit,
			Message: errors.InvalidDataType("string", v, lang),
		}
	}

	var emailLen = len(email)

	if emailLen < 5 {
		return "", &types.ValidationErr{
			Field:   e.Field,
			Path:    path,
			Value:   email,
			Message: shortEmailErr(emailLen, lang),
		}
	}

	if emailLen > 320 {
		return "", &types.ValidationErr{
			Field:   e.Field,
			Path:    path,
			Value:   email,
			Message: longEmailErr(emailLen, lang),
		}
	}

	var parts = strings.Split(email, "@")

	var partsLen = len(parts)

	if partsLen == 1 {
		return "", &types.ValidationErr{
			Field:   e.Field,
			Path:    path,
			Value:   email,
			Message: missingAtSignErr(lang),
		}
	}

	if partsLen > 2 {
		return "", &types.ValidationErr{
			Field:   e.Field,
			Path:    path,
			Value:   email,
			Message: tooManyAtSignErr(lang),
		}
	}

	var err = IsLocalPartValid(e.Field, email, parts[0], path, lang)

	if err != nil {
		return "", err
	}

	err = IsDomainNameValid(e.Field, email, parts[1], path, e.CheckMx, lang)
	if err != nil {
		return "", err
	}

	return email, nil

}

func shortEmailErr(emailLen int, lang string) string {
	var strLen = strconv.Itoa(emailLen)
	if lang == "ar" {
		return "يجب إطالة هذا البريد إلى 5 من الحروف أو أكثر (أنت تستخدم حاليا" + strLen + " من الحروف)"
	}

	return "Should lengthen this email to 5 characters or more (you are currently using " + strLen + " characters)"
}

func longEmailErr(emailLen int, lang string) string {
	var strLen = strconv.Itoa(emailLen)
	if lang == "ar" {
		return "يجب تقصير هذا البريد إلى 320 من الحروف أو أقل (أنت حاليا تستخدم " + strLen + " من الحروف)"
	}
	return "Should shorten this email to 320 characters or less (you are currently using " + strLen + " characters)"
}

func missingAtSignErr(lang string) string {
	if lang == "ar" {
		return "يجب أن يحتوي البريد الإلكتروني على العلامة @"
	}

	return "The email should contain the @ sign"
}

func tooManyAtSignErr(lang string) string {
	if lang == "ar" {
		return "يجب أن يحتوي البريد الإلكتروني على علامة @ واحدة فقط"
	}

	return "The email should contain only one @ sign"
}

var allowedChars_domainPart = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789."

var PERIOD byte = '.'

func IsDomainNameValid(field, email, domain string, path []any, checkMx bool, lang string) error {
	var l = len(domain)

	if l < 1 {
		return &types.ValidationErr{
			Field:   field,
			Path:    path,
			Value:   email,
			Message: domain_shortErr(l, lang),
		}
	}

	if l > 255 {
		return &types.ValidationErr{
			Field:   field,
			Path:    path,
			Value:   email,
			Message: domain_longErr(l, lang),
		}
	}

	if domain[0] == PERIOD {
		return &types.ValidationErr{
			Field:   field,
			Path:    path,
			Value:   email,
			Message: domain_startsWithPeriodErr(lang),
		}
	}

	if domain[l-1] == PERIOD {
		return &types.ValidationErr{
			Field:   field,
			Path:    path,
			Value:   email,
			Message: domain_endsWithPeriodErr(lang),
		}
	}

	for i, r := range domain {
		if r == rune(PERIOD) {
			if i > 1 && domain[i-1] == PERIOD {
				return &types.ValidationErr{
					Field:   field,
					Path:    path,
					Value:   email,
					Message: domain_twoAdjacentPeriodsErr(lang),
				}
			}
		}
		if !strings.ContainsRune(allowedChars_domainPart, r) {
			return &types.ValidationErr{
				Field:   field,
				Path:    path,
				Value:   email,
				Message: domain_invalidCharacterErr(r, lang),
			}
		}
	}

	for i, label := range strings.Split(domain, ".") {

		var l = len(label)

		if l > 63 {
			return &types.ValidationErr{
				Field:   field,
				Path:    path,
				Value:   email,
				Message: domain_longLabelErr(l, i+1, lang),
			}
		}

	}

	if checkMx {

		mxRcds, err := net.LookupMX(domain)

		if err != nil || len(mxRcds) == 0 {
			return &types.ValidationErr{
				Field:   field,
				Path:    path,
				Value:   email,
				Message: domain_invalidErr(lang),
			}
		}
	}

	return nil
}

func domain_shortErr(domainLen int, lang string) string {
	var strLen = strconv.Itoa(domainLen)

	if lang == "ar" {
		return "يجب إطالة الجزء الذي بعد العلامة @ إلى حرف أو أكثر (أنت تستخدم حاليا " + strLen + " من الحروف)"
	}

	return "Should lengthen the part after @ sign to 1 character or more (you are currently using " + strLen + " characters)"
}

func domain_longErr(domainLen int, lang string) string {
	var strLen = strconv.Itoa(domainLen)

	if lang == "ar" {
		return "يجب تقصير الجزء الذي بعد العلامة @ إلى 255 من الحروف أو أقل (أنت تستخدم حاليا " + strLen + " من الحروف)"
	}

	return "Should shorten the part after @ sign to 255 characters or less (you are currently using " + strLen + " characters)"
}

func domain_invalidCharacterErr(r rune, lang string) string {
	if lang == "ar" {
		return "يجب ألَّا يحتوي الجزء الذي بعد العلامة @ على الرمز " + string(r)
	}
	return "The part after @ sign should not contain the " + string(r) + " symbol"
}

func domain_startsWithPeriodErr(lang string) string {
	if lang == "ar" {
		return "يجب ألَّا يبدأ الجزء بعد العلامة @ بنقطة (.)"
	}

	return "The part after @ sign should not begin with a period (.)"
}

func domain_endsWithPeriodErr(lang string) string {
	if lang == "ar" {
		return "يجب ألَّا ينتهي الجزء بعد العلامة @ بنقطة (.)"
	}

	return "The part after @ sign should not end with a period (.)"
}

func domain_twoAdjacentPeriodsErr(lang string) string {
	if lang == "ar" {
		return "يجب ألَّا يحتوي الجزء بعد العلامة @ على نقطتين متجاورتين"
	}

	return "The part after the @ sign should not contain two adjacent periods"
}

func domain_longLabelErr(labelLen, sec int, lang string) string {
	var strLen, strSec = strconv.Itoa(labelLen), strconv.Itoa(sec)

	if lang == "ar" {
		return "يجب ألا يتجاوز طول المقطع الواحد من الجزء بعد العلامة @ 63 حرفًا (أنت تستخدم حاليا في المقطع ال" + strSec + " " + strLen + " حرفاً)"
	}

	return "The length of one section of the part after the @ tag must not exceed 63 characters (you are currently using " + strLen + " characters in section " + strSec + ")"
}

func domain_invalidErr(lang string) string {
	if lang == "ar" {
		return "الجزء بعد العلامة @ غير صالح."
	}
	return "The part after @ sign is invalid"
}

// var allowedLocalPartChars = domainAllowedChars + "-+!#$%&'*/=?`^{}[]|~"

var allowedChars_localPart = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.+"

func IsLocalPartValid(field, email, localPart string, path []any, lang string) error {
	var l = len(localPart)

	if l == 0 {
		return &types.ValidationErr{
			Field:   field,
			Path:    path,
			Value:   email,
			Message: local_shortErr(l, lang),
		}
	}

	if l > 64 {
		return &types.ValidationErr{
			Field:   field,
			Path:    path,
			Value:   email,
			Message: local_longErr(l, lang),
		}
	}

	if localPart[0] == PERIOD {
		return &types.ValidationErr{
			Field:   field,
			Path:    path,
			Value:   email,
			Message: local_startsWithPeriodErr(lang),
		}
	}

	if localPart[l-1] == PERIOD {
		return &types.ValidationErr{
			Field:   field,
			Path:    path,
			Value:   email,
			Message: local_endsWithPeriodErr(lang),
		}
	}

	for i, r := range localPart {
		if r == rune(PERIOD) {
			if i > 1 && localPart[i-1] == PERIOD {
				return &types.ValidationErr{
					Field:   field,
					Path:    path,
					Value:   email,
					Message: local_twoAdjacentPeriodsErr(lang),
				}
			}
		}
		if !strings.ContainsRune(allowedChars_localPart, r) {
			return &types.ValidationErr{
				Field:   field,
				Path:    path,
				Value:   email,
				Message: local_invalidCharacterErr(r, lang),
			}
		}
	}

	return nil
}

func local_shortErr(localLen int, lang string) string {
	var strLen = strconv.Itoa(localLen)

	if lang == "ar" {
		return "يجب إطالة الجزء الذي قبل العلامة @ إلى حرف أو أكثر (أنت تستخدم حاليا " + strLen + " من الحروف)"
	}

	return "Should lengthen the part before @ sign to 1 character or more (you are currently using " + strLen + " characters)"
}

func local_longErr(localLen int, lang string) string {
	var strLen = strconv.Itoa(localLen)

	if lang == "ar" {
		return "يجب تقصير الجزء الذي قبل العلامة @ إلى 255 من الحروف أو أقل (أنت تستخدم حاليا " + strLen + " من الحروف)"
	}

	return "Should shorten the part before @ sign to 255 characters or less (you are currently using " + strLen + " characters)"
}

func local_invalidCharacterErr(r rune, lang string) string {
	if lang == "ar" {
		return "يجب ألَّا يحتوي الجزء الذي قبل العلامة @ على الرمز " + string(r)
	}
	return "The part before @ sign should not contain the " + string(r) + " symbol"
}

func local_startsWithPeriodErr(lang string) string {
	if lang == "ar" {
		return "يجب ألَّا يبدأ الجزء قبل العلامة @ بنقطة (.)"
	}

	return "The part before @ sign should not begin with a period (.)"
}

func local_endsWithPeriodErr(lang string) string {
	if lang == "ar" {
		return "يجب ألَّا ينتهي الجزء قبل العلامة @ بنقطة (.)"
	}

	return "The part before @ sign should not end with a period (.)"
}

func local_twoAdjacentPeriodsErr(lang string) string {
	if lang == "ar" {
		return "يجب ألَّا يحتوي الجزء قبل العلامة @ على نقطتين متجاورتين"
	}

	return "The part before the @ sign should not contain two adjacent periods"
}
