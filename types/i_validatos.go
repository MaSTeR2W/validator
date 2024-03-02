package types

type Validator[T any] interface {
	Validate(v any, lang string) (T, error)
	GetField() string
}
