package errors

import "strings"

func RequiredFieldsErr(props []string, lang string) string {
	var strProps = strings.Join(props, ", ")
	if lang == "ar" {
		return "حقول مطلوبة: (" + strProps + ")"
	}
	return "Required fields: (" + strProps + ")"
}
