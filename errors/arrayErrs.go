package errors

import "strconv"

func ShortArrErr(exp int, got int, lang string) string {
	var sExp, sGot = strconv.Itoa(exp), strconv.Itoa(got)

	if lang == "ar" {
		return "يجب زيادة عدد عناصر المصفوفة إلى " + sExp + " أو أكثر (عدد عناصر المصفوفة حاليا " + sGot + ")"
	}
	return "Should increase the number of elements of array to " + sExp + " or more (the current number of element is " + sGot + ")"
}

func LongArrErr(exp int, got int, lang string) string {
	var sExp, sGot = strconv.Itoa(exp), strconv.Itoa(got)

	if lang == "ar" {
		return "يجب إنقاص عدد عناصر المصفوفة إلى " + sExp + " أو أقل (عدد عناصر المصفوفة حاليا " + sGot + ")"
	}
	return "Should decrease the number of elements of array to " + sExp + " or less (the current number of element is " + sGot + ")"
}

func InvalidJSONArr(lang string) string {
	if lang == "ar" {
		return "لا يمكن تحويل القيمة إلى ([]string)"
	}
	return "Cannot convert value to ([]string)"
}
