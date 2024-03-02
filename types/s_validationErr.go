package types

import (
	"strings"

	"github.com/MaSTeR2W/validator/hprFns"
)

var omit = 'i'

var Omit = &omit

type ValidationErr struct {
	Field   string
	Path    []any
	Value   any
	Message string
}

func (v *ValidationErr) Error() string {
	return v.Message
}

func (v ValidationErr) MarshalJSON() ([]byte, error) {
	var jsonStr = `{"field": "` + v.Field + `",`

	if v.Path != nil {
		var arrStr = ""
		for _, step := range v.Path {
			arrStr += "," + hprFns.ToJSONString(step)
		}

		jsonStr += `"path": ` + arrStr[1:] + "],"
	}

	if v.Value != Omit {
		jsonStr += `"value": ` + hprFns.ToJSONString(v.Value) + ","
	}

	return []byte(jsonStr + `"message": "` + strings.ReplaceAll(v.Message, `"`, `\"`) + `" }`), nil
}
