package types

type StructValidator[T any] func(v any, field string, path []any, lang string) (*T, error)
