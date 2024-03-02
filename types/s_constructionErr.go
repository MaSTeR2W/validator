package types

type ConstructionErr struct {
	Message string
}

func (c *ConstructionErr) Error() string {
	return c.Message
}
