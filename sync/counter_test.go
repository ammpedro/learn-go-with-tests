package sync

import (
	"sync"
	"testing"
)

func NewCounter() *Counter {
	return &Counter{}
}

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3x leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)

	})

	t.Run("it runs safely cocurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		//A WaitGroup waits for a collection of goroutines to finish.
		var wg sync.WaitGroup
		// main goroutine calls Add to set the number of goroutines to wait for
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				// Then each of the goroutines runs and calls Done when finished.
				w.Done()
			}(&wg)
		}

		// Wait can be used to block until all goroutines have finished.
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t *testing.T, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
