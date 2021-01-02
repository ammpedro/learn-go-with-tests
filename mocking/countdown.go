package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// Sleeper is ...
type Sleeper interface {
	Sleep()
}

// ConfigurableSleeper ...
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep is configurable
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// Countdown displays a countdown
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{2 * time.Second, time.Sleep}

	// send data to os.Stdout for users to see the countdown on the terminal
	Countdown(os.Stdout, sleeper)
}
