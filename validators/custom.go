package validators

var Strings = map[string]func(any, []any, string) (*string, error){}

var Intgers = map[string]func(any, []any, string) (*string, error){}
