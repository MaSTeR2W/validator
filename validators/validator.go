package validators

type Validator interface {
	Validate(v any, lang string) error
	GetField() string
}
