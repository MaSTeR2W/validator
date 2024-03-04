package validators

import (
	"time"

	"github.com/MaSTeR2W/validator/errors"
	"github.com/MaSTeR2W/validator/types"
)

type Date struct {
	Field     string
	BeforeNow bool
	Before    time.Time
	AfterNow  bool
	After     time.Time
}

func (d *Date) GetField() string {
	return d.Field
}

func (d *Date) Validate(v any, path []any, lang string) (string, error) {
	var dateStr, ok = v.(string)

	if !ok {
		return "", &types.ValidationErr{
			Field:   d.Field,
			Path:    path,
			Value:   v,
			Message: errors.InvalidDataType("string", v, lang),
		}
	}
	var dateStruct, err = time.Parse("2006-01-02", dateStr)

	if d.BeforeNow || d.AfterNow {
		var now, _ = time.Parse("2006-01-02", time.Now().Format("2006-01-02"))

		if d.AfterNow && !dateStruct.After(now) {
			return "", &types.ValidationErr{
				Field:   d.Field,
				Path:    path,
				Value:   v,
				Message: shouldAfterErr(now, dateStr, lang),
			}
		}

		if d.BeforeNow && !dateStruct.Before(now) {
			return "", &types.ValidationErr{
				Field:   d.Field,
				Path:    path,
				Value:   v,
				Message: shouldBeforeErr(now, dateStr, lang),
			}
		}
	} else {
		if !d.After.IsZero() && !dateStruct.After(d.After) {
			return "", &types.ValidationErr{
				Field:   d.Field,
				Path:    path,
				Value:   v,
				Message: shouldAfterErr(d.After, dateStr, lang),
			}
		}

		if !d.Before.IsZero() && !dateStruct.Before(d.Before) {
			return "", &types.ValidationErr{
				Field:   d.Field,
				Path:    path,
				Value:   v,
				Message: shouldBeforeErr(d.Before, dateStr, lang),
			}
		}

	}

	if err != nil {
		return "", &types.ValidationErr{
			Field:   d.Field,
			Path:    path,
			Value:   dateStr,
			Message: invalidDateErr(lang),
		}
	}

	return dateStr, nil
}

func invalidDateErr(lang string) string {
	if lang == "ar" {
		return "تاريخ غير صالح"
	}
	return "invalid date"
}

func shouldAfterErr(aft time.Time, input string, lang string) string {
	var sAft = aft.Format("2006-01-02")
	if lang == "ar" {
		return "يجب أن يكون التاريخ بعد " + sAft + " (التاريخ المدخل: " + input + ")"
	}
	return "Date should be after " + sAft + " (the entered date: " + input + ")"
}

func shouldBeforeErr(bef time.Time, input string, lang string) string {
	var sBef = bef.Format("2006-01-02")
	if lang == "ar" {
		return "يجب أن يكون التاريخ قبل: " + sBef + " (التاريخ المدخل: " + input + ")"
	}

	return "Date should be before " + sBef + " (the entered date: " + input + ")"
}
