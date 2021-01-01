// +build wireinject

package main

import (
	"github.com/google/wire"
	"golang-di-example/src/domain"
)

func InitializeEvent() domain.Event {
	wire.Build(domain.NewEvent, domain.NewGreeter, domain.NewMessage)
	return domain.Event{}
}