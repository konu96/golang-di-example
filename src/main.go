package main

import (
	"fmt"
	"go.uber.org/dig"
	"golang-di-example/src/domain"
	"golang-di-example/src/wire"
)

func main() {
	container := dig.New()

	container.Provide(domain.NewMessage)
	container.Provide(domain.NewGreeter)
	container.Provide(domain.NewEvent)

	container.Invoke(func(event domain.Event) {
		fmt.Println(event.Start())
	})

	event := wire.InitializeEvent()
	fmt.Println(event.Start())
}