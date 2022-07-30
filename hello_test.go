package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, actual, expected string) {
		t.Helper()

		if actual != expected {
			t.Errorf("got %q want %q", actual, expected)
		}
	}

	t.Run("saying go-tdd-hello-world to people", func(t *testing.T) {
		actual := Hello("Burak", "")
		expected := "Hello, Burak"

		assertCorrectMessage(t, actual, expected)
	})

	t.Run("say 'Hello, World' when an empty string is supplied'", func(t *testing.T) {
		actual := Hello("", "")
		expected := "Hello, World"

		assertCorrectMessage(t, actual, expected)
	})

	t.Run("in spanish", func(t *testing.T) {
		actual := Hello("Elodie", "spanish")
		expected := "Hola, Elodie"

		assertCorrectMessage(t, actual, expected)
	})

	t.Run("in french", func(t *testing.T) {
		actual := Hello("Elodie", "french")
		expected := "Bonjour, Elodie"

		assertCorrectMessage(t, actual, expected)
	})
}
