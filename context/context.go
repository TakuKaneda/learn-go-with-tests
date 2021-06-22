package context

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// context
		ctx := r.Context()
		// channel for store fetching data
		data := make(chan string, 1)
		// fetch data by another goroutine
		go func() {
			data <- store.Fetch()
		}()
		// fetching data or cancel call
		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
