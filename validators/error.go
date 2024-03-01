package validators

import (
	"reflect"
	"strings"
)

var null *int8

type ValidationErr struct {
	Field   string `json:"field,omitempty"`
	Value   any    `json:"value,omitempty"`
	Message string `json:"message"`
}

func (v *ValidationErr) Error() string {
	return v.Message
}

type ValidationErrs []error

func (v *ValidationErrs) Error() string {
	var str = "[\n\t"

	var s = []error(*v)
	var lastIndex = len(s) - 1

	for i := 0; i < lastIndex; i++ {

		str += strings.ReplaceAll(s[i].Error(), "\n\t", "\n\t\t") + "\n\t"
	}

	return str + s[lastIndex].Error() + "\n]"
}

func invalidDataType(exp string, got any, lang string) string {
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
