package greet

import (
	"testing"
)

func TestGreetDominic(t *testing.T) {
	v := SayHello("Dominic")
	if v != "Bonjour" {
		t.Errorf("expected 'Bonjour', got '%s'", v)
	}
}

func TestGreetOthers(t *testing.T) {
	v := SayHello("Bob")
	if v != "Hello" {
		t.Errorf("expected 'Hello', got '%s'", v)
	}
}
