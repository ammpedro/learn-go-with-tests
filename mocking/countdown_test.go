package main

import (
	"bytes"
	"reflect"
	"testing"
)

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const sleep = "sleep"
const write = "write"

func TestCountdown(t *testing.T) {

	t.Run("should print count down", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		// send to bytes.Buffer so tests can capture data being generated
		Countdown(buffer, &CountdownOperationsSpy{})

		got := buffer.String()
		want := "3\n2\n1\nGo!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should sleep before every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write, //3
			sleep,
			write, //2
			sleep,
			write, //1
			sleep,
			write, //Go!
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted %v got %v", want, spySleepPrinter.Calls)
		}

	})
}
