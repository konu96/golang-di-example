package main

import (
	"go.uber.org/dig"
	"golang-di-example/src/domain"
)

func main() {
	container := dig.New()

	container.Provide(domain.NewMessage)
	container.Provide(domain.NewGreeter)
	container.Provide(domain.NewEvent)

	container.Invoke(func(event domain.Event) {
		event.Start()
	})

	event := InitializeEvent()
	event.Start()
}