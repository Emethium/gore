package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "english")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello when an empty string is provided", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying in spanish", func(t *testing.T) {
		got := Hello("Santana", "spanish")
		want := "Hola, Santana"
		assertCorrectMessage(t, got, want)
	})
}
