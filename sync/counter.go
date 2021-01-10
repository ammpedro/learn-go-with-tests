package sync

import "sync"

// Counter counts
type Counter struct {
	mu    sync.Mutex
	value int
}

// Inc increments the counter
func (c *Counter) Inc() {
	//any goroutine calling Inc will acquire the lock on Counter if they are first.
	//All the other goroutines will have to wait for it to be Unlocked before getting access
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++
}

// Value returns the counter's current value
func (c *Counter) Value() int {
	return c.value
}
