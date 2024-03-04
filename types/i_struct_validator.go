package types

type StructValidator[T any] interface {
	Validate(v any, path []any, lang string) error
}
