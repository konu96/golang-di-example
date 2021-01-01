package domain

import "fmt"

type Event struct {
	Greeter Greeter // <- adding a Greeter field
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}