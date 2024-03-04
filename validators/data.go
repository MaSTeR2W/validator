package validators

import (
	"time"

	"github.com/MaSTeR2W/validator/errors"
	"github.com/MaSTeR2W/validator/types"
)

type Date struct {
	Field     string
	BeforeNow bool
	Before    int64 // unix
	AfterNow  bool
	After     int64 // unix
}

func (d *Date) GetField() string {
	return d.Field
}

func (d *Date) Validate(v any, path []any, lang string) (string, error) {
	var date, ok = v.(string)

	if !ok {
		return "", &types.ValidationErr{
			Field:   d.Field,
			Path:    path,
			Value:   v,
			Message: errors.InvalidDataType("string", v, lang),
		}
	}
	var t, err = time.Parse("2006-01-02", date)

	if err != nil {
		return "", &types.ValidationErr{
			Field:   d.Field,
			Path:    path,
			Value:   date,
			Message: invalidDateErr(lang),
		}
	}

	return t.Format("2006-01-02"), nil
}

func invalidDateErr(lang string) string {
	if lang == "ar" {
		return "تاريخ غير صالح"
	}
	return "invalid date"
}
