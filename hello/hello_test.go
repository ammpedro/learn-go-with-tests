package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, but want %q", got, want)
		}
	}

	t.Run("should say hello to you", func(t *testing.T) {
		got := Hello("you", "")
		want := "Hello, you"

		assertCorrectMessage(t, got, want)
	})

	t.Run("should say \"Hello, World\" when passing empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("should say hello in Spanish", func(t *testing.T) {
		got := Hello("Pepito", "es")
		want := "Hola, Pepito"

		assertCorrectMessage(t, got, want)
	})

	t.Run("should say hello in French", func(t *testing.T) {
		got := Hello("Pierre", "fr")
		want := "Bonjour, Pierre"

		assertCorrectMessage(t, got, want)
	})
}
