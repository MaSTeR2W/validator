package validators

import (
	"mime/multipart"
	"slices"
	"strconv"
	"strings"

	"github.com/MaSTeR2W/validator/types"
)

type Size = int64

const (
	BYTE Size = 1
	KB        = BYTE * 1024
	MG        = KB * 1024
	GB        = MG * 1024
)

type File struct {
	Field      string
	Type       string
	NilAble    bool
	Extensions []string
	MaxSize    int64
}

func (f *File) GetField() string {
	return f.Field
}

func (f *File) Validate(files []*multipart.FileHeader, path []any, lang string) (*multipart.FileHeader, error) {

	var l = len(files)

	if l == 0 || files[0] == nil {
		if f.NilAble {
			return nil, nil
		}
		return nil, &types.ValidationErr{
			Field:   f.Field,
			Value:   types.Omit,
			Path:    path,
			Message: noFileErr(lang),
		}
	}

	if l > 1 {
		return nil, &types.ValidationErr{
			Field:   f.Field,
			Value:   types.Omit,
			Path:    path,
			Message: extraFilesErr(l, lang),
		}
	}
	var file = files[0]

	var mimeType = file.Header.Get("Content-type")
	var extension, found = strings.CutPrefix(mimeType, f.Type+"/")

	if !found {
		return nil, &types.ValidationErr{
			Field:   f.Field,
			Path:    path,
			Value:   types.Omit,
			Message: invalidFileTypeErr(f.Type, strings.Split(mimeType, "/")[0], lang),
		}
	}

	if !slices.Contains(f.Extensions, extension) {
		return nil, &types.ValidationErr{
			Field:   f.Field,
			Path:    path,
			Value:   types.Omit,
			Message: invalidExtensionErr(f.Extensions, extension, lang),
		}
	}

	if f.MaxSize != 0 && file.Size > f.MaxSize {
		return nil, &types.ValidationErr{
			Field:   f.Field,
			Path:    path,
			Value:   types.Omit,
			Message: bigSizeFileErr(f.MaxSize, file.Size, lang),
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

func invalidFileTypeErr(exp string, got string, lang string) string {
	if lang == "ar" {
		return "يجب أن يكون الملف من النوع (" + exp + ") وليس (" + got + ")"
	}
	return "The file type should be (" + exp + "), not (" + got + ")"
}

func invalidExtensionErr(exp []string, got string, lang string) string {
	var strExp = strings.Join(exp, ",")
	if lang == "ar" {
		return "يجب أن يكون امتداد الملف واحد من: (" + strExp + ") وليس (" + got + ")"
	}
	return "The file extension should be one of: (" + strExp + "), not (" + got + ")"
}

func bigSizeFileErr(exp int64, got int64, lang string) string {
	var sExp, sGot = strconv.FormatInt(exp, 10), strconv.FormatInt(got, 10)

	if lang == "ar" {
		return "يجب ألَّا يزيد حجم الملف على " + sExp + "B (حجم الملف الحالي " + sGot + "B)"
	}
	return "The file size should not exceed " + sExp + "B (The current file size is " + sGot + "B)"
}
