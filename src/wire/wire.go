// +build wireinject

package wire

import (
	"github.com/google/wire"
	"golang-di-example/src/domain"
)

func InitializeEvent() domain.Event {
	wire.Build(domain.NewEvent, domain.NewGreeter, domain.NewMessage)
	return domain.Event{}
}

// テスト用に DI するファイルを作った方が良い
func InitializeMockEvent() domain.Event {
	wire.Build(domain.NewEvent, domain.NewMockGreeter)
	return domain.Event{}
}