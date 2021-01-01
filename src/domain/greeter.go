package domain

type GreeterInterface interface {
	Greet() Message
}

type Greeter struct {
	Message Message // <- adding a Message field
}

func (g Greeter) Greet() Message {
	return g.Message
}

func NewGreeter(m Message) GreeterInterface {
	return Greeter{Message: m}
}

func NewMockGreeter() GreeterInterface {
	return Greeter{
		Message: "Hi Test",
	}
}