package validators

import (
	"mime/multipart"
	"slices"
	"strconv"
	"strings"

	"github.com/MaSTeR2W/validator/types"
)

type File struct {
	Field      string
	Type       string
	Extensions []string
}

func (f *File) GetField() string {
	return f.Field
}

func (f *File) Validate(files []*multipart.FileHeader, path []any, lang string) (*multipart.FileHeader, error) {

	var l = len(files)

	if l == 0 || files[0] == nil {
		return nil, &types.ValidationErr{
			Field:   f.Field,
			Path:    path,
			Message: noFileErr(lang),
		}
	}

	if l > 1 {
		return nil, &types.ValidationErr{
			Field:   f.Field,
			Path:    path,
			Message: extraFilesErr(l, lang),
		}
	}
	var file = files[0]

	var mimeType = file.Header.Get("Content-type")
	var extension, found = strings.CutPrefix(mimeType, "image/")

	if !found {
		return nil, &types.ValidationErr{
			Field:   f.Field,
			Path:    path,
			Message: invalidFileType(strings.Split(mimeType, "/")[0], lang),
		}
	}

	if !slices.Contains(f.Extensions, extension) {
		return nil, &types.ValidationErr{
			Field:   f.Field,
			Path:    path,
			Message: invalidExtension(f.Extensions, extension, lang),
		}
	}

	return file, nil
}

func extraFilesErr(got int, lang string) string {
	var sGot = strconv.Itoa(got)
	if lang == "ar" {
		return "يجب إنقاص عدد الملفات إلى ملف واحد (عدد الملفات حاليا " + sGot + ")"
	}
	return "Should reduce the number of files to one file (the number of files currently " + sGot + ")"
}

func noFileErr(lang string) string {
	if lang == "ar" {
		return "ملف مطلوب"
	}
	return "Required file."
}

func invalidFileType(got string, lang string) string {
	if lang == "ar" {
		return "يجب أن يكون الملف من النوع (image) وليس (" + got + ")"
	}
	return "The file type should be (image), not (" + got + ")"
}

func invalidExtension(exp []string, got string, lang string) string {
	var strExp = strings.Join(exp, ",")
	if lang == "ar" {
		return "يجب أن يكون امتداد الملف واحد من: (" + strExp + ") وليس (" + got + ")"
	}
	return "The file extension should be one of: (" + strExp + "), not (" + got + ")"
}
