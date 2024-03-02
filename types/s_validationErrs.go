package types

import "strings"

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
