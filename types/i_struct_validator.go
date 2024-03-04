package types

type StructValidator[T any] interface {
	Validate(v any, field string, path []any, lang string) error
}
