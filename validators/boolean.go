package validators

type Bool struct {
	Field  string
	NotNil bool
}

func (b *Bool) GetField() string {
	return b.Field
}

func (b *Bool) Validate(v any, lang string) error {

	if v == nil {
		if b.NotNil {
			return &ValidationErr{
				Field:   b.Field,
				Value:   null,
				Message: invalidDataType("bolean", v, lang),
			}
		}
		return nil
	}

	var _, ok = v.(bool)

	if !ok {
		return &ValidationErr{
			Field:   b.Field,
			Value:   v,
			Message: invalidDataType("boolean", v, lang),
		}
	}
	return nil
}
