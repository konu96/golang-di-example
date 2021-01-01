package domain

type Message string

type GreeterInterface interface {
	Greet() Message
}

func NewMessage() Message {
	return Message("Hi there!")
}

func NewGreeter(m Message) GreeterInterface {
	return Greeter{Message: m}
}