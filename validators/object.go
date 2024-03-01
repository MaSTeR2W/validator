package validators

import "strings"

type Object struct {
	Field    string
	Required []string
	OneOf    []string
	NotNil   bool
	Props    []Validator
}

func (o *Object) GetField() string {
	return o.Field
}

func (o *Object) Validate(v any, lang string) error {
	var errs = ValidationErrs{}
	if v == nil {
		if o.NotNil {
			errs = append(errs, &ValidationErr{
				Field:   o.Field,
				Value:   null,
				Message: invalidDataType("map[string]any", v, lang),
			})

			return &errs
		}

		return nil
	}

	var m, ok = v.(map[string]any)

	if !ok {
		errs = append(errs, &ValidationErr{
			Field:   o.Field,
			Value:   v,
			Message: invalidDataType("map[string]any", v, lang),
		})
		return &errs
	}

	if o.Required != nil {
		for _, key := range o.Required {
			if _, ok := m[key]; !ok {
				errs = append(errs, &ValidationErr{
					Field:   key,
					Message: missingProp(lang),
				})
			}
		}
	}

	if o.OneOf != nil {
		var found = false
		for _, key := range o.OneOf {
			if _, ok := m[key]; ok {
				found = true
				break
			}
		}
		if !found {
			errs = append(errs, &ValidationErr{
				Message: requiredOneOf(o.OneOf, lang),
			})

			// there is no need to check for other validation
			return &errs
		}
	}

	if len(o.Props) > 0 {
		for _, prop := range o.Props {
			if val, ok := m[prop.GetField()]; ok {
				if err := prop.Validate(val, lang); err != nil {
					errs = append(errs, err)
				}
			}
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return &errs
}

func missingProp(lang string) string {
	if lang == "ar" {
		return "خاصية مطلوبة"
	}

	return "Required property"
}

func requiredOneOf(oneOf []string, lang string) string {
	var joinedOneOf = strings.Join(oneOf, ", ")
	if lang == "ar" {
		return "مطلوب واحدة من الخاصيات التالية: (" + joinedOneOf + ")"
	}
	return "One of the following properties is required: (" + joinedOneOf + ")"
}

/* func mapKeys(m map[string]any) []string {
	var keys = make([]string, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
} */
