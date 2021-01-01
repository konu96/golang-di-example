package domain

import "fmt"

type Event struct {
	Greeter GreeterInterface // <- adding a Greeter field
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func NewEvent(g GreeterInterface) Event {
	return Event{Greeter: g}
}