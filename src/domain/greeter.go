package domain

type Greeter struct {
	Message Message // <- adding a Message field
}

func (g Greeter) Greet() Message {
	return g.Message
}