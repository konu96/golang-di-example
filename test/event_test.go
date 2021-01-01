package test

import (
	"golang-di-example/src/wire"
	"testing"
)

func TestEvent(t *testing.T) {
	event := wire.InitializeMockEvent()
	actual := event.Start()
	expected := "Hi Test"

	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}