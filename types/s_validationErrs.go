package types

import (
	"strings"
)

type ValidationErrs []ValidationErr

func (v *ValidationErrs) Error() string {
	var str = "[\n\t"

	var lastIndex = len(*v) - 1

	for i := 0; i < lastIndex; i++ {

		str += strings.ReplaceAll((*v)[i].Error(), "\n\t", "\n\t\t") + "\n\t"
	}

	return str + (*v)[lastIndex].Error() + "\n]"
}

func (v ValidationErrs) MarshalJSON() ([]byte, error) {
	var es = []byte{91}

	for _, e := range v {
		js, _ := e.MarshalJSON()
		es = append(es, 44)
		es = append(es, js...)
	}

	return append(es, 93), nil
}
