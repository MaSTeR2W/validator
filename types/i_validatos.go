package types

type Validator[T any] interface {
	Validate(v any, path []any, lang string) (T, error)
	GetField() string
}
