package context

import (
	"context"
	"fmt"
	"net/http"
)

// Store is the store
type Store interface {
	Fetch(ctx context.Context) (string, error)
	//Cancel()
}

// Server accepts a store
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return //stuff
		}

		fmt.Fprint(w, data)
	}
}
