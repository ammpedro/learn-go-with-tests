package context

import (
	"fmt"
	"net/http"
)

// Store is the store
type Store interface {
	Fetch() string
	Cancel()
}

// Server accepts a store
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data := make(chan string, 1)

		go func() {
			//write result into data channel
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			//happens if fetch was first
			fmt.Fprint(w, d)
		case <-ctx.Done():
			//context's Done() method returns a channel signal when the context is done or cancelled
			//if done/cancelled signal is received, cancel the store
			store.Cancel()
		}
	}
}
