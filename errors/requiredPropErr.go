package errors

func RequiredFieldErr(lang string) string {
	if lang == "ar" {
		return "حقل ملطوب"
	}
	return "Required field"
}
