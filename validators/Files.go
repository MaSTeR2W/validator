package validators

import (
	"mime/multipart"
	"slices"
	"strings"

	"github.com/MaSTeR2W/validator/types"
)

const (
	// application ----- >
	//

	APP_XLSX = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" // excel

	APP_DOCX = "application/vnd.openxmlformats-officedocument.wordprocessingml.document" // word

	APP_PDF = "application/pdf" // pdf files

	//
	//
	// image ------ >
	//

	IMG_PNG  = "image/png"  // png images
	IMG_JPEG = "image/jpeg" // jpeg images
	IMG_JPG  = "image/jpg"  // jpg images
)

type FileV2 struct {
	Field   string
	Types   []string
	MaxSize int64
}

func (f *FileV2) GetField() string {
	return f.Field
}

func (f *FileV2) Validate(
	file *multipart.FileHeader,
	path []any,
	lang string,
) (*multipart.FileHeader, error) {

	var mimeType = file.Header.Get("Content-type")

	if !slices.Contains(f.Types, mimeType) {
		return nil, &types.ValidationErr{
			Field:   f.Field,
			Path:    path,
			Value:   types.Omit,
			Message: invalidFileV2TypeErr(f.Types, mimeType, lang),
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

func invalidFileV2TypeErr(exps []string, got string, lang string) string {
	var exp = strings.Join(exps, ", ")
	if lang == "ar" {
		return "يجب أن يكون الملف من النوع (" + exp + ") وليس (" + got + ")"
	}
	return "The file type should be (" + exp + "), not (" + got + ")"
}
