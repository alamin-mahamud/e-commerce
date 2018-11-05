package cmd

import "testing"

func TestHelloWorld(t *testing.T) {
	a := "Hello"
	b := "World"

	if a == b {
		t.Error("How on earth is this possible?")
	}
}
