package errors

import "strings"

func OneOfIsRequired(props []string, lang string) string {
	var strProps = strings.Join(props, ", ")
	if lang == "ar" {
		return "مطلوب حقل واحد على الأقل مما يلي: (" + strProps + ")"
	}
	return "At least one of the following fields is required: (" + strProps + ")"
}
