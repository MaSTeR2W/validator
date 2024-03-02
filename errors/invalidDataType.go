package errors

import (
	"reflect"
)

func InvalidDataType(exp string, got any, lang string) string {
	var t string
	if got == nil {
		t = "null"
	} else {
		t = reflect.TypeOf(got).String()
	}

	if lang == "ar" {
		return "يجب أن تكون البيانات من النوع (" + exp + "), وليس (" + t + ")"
	}
	return "Data should be of type (" + exp + "), not (" + t + ")"
}
