package auth

import "net/http"

func authHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func NewAuthHandler() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return authHandler(h)
	}
}
