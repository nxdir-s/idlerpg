package server

import "net/http"

type ServerHandler func(http.ResponseWriter, *http.Request) error

func httpHandler(fn ServerHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
