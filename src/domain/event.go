package domain

type Event struct {
	Greeter GreeterInterface // <- adding a Greeter field
}

func (e Event) Start() string {
	msg := e.Greeter.Greet()
	return string(msg)
}

func NewEvent(g GreeterInterface) Event {
	return Event{Greeter: g}
}