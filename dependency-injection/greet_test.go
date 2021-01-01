package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {

	t.Run("should test actual printing (fmt.Printf()", func(t *testing.T) {
		// buffer type from the bytes package implements the Writer interface
		buffer := bytes.Buffer{}

		// buffer is used as the Writer for function greet
		Greet(&buffer, "Banana")

		got := buffer.String()
		want := "Hello, Banana"

		if got != want {
			t.Errorf("got %q, but want %q", got, want)
		}
	})
}
