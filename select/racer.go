package racer

import (
	"fmt"
	"net/http"
	"time"
)

//Racer accepts two urls and returns the faster one
func Racer(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

//RacerSelect accepts two urls and returns the faster one using a Goroutine
func RacerSelect(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, 10*time.Second)
}

//ConfigurableRacer is a the original RacerSelect but with configurable timeouts
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	//wait on mutiple channels
	//receive from channel operations are blocking, waits for a value
	//the url of the first ping to send a value is returned first
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

//Ping accepts a url and returns a channel
//Use channel to signal that http.Get is complete
//struct{} is the smallest data type availble, no need to allocate what we dont need
func ping(url string) chan struct{} {
	//always make channels
	//using var will save a zero value
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
