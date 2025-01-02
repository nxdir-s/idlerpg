package auth

import (
	"net/http"
	"strings"

	"github.com/nxdir-s/idlerpg/internal/core/valobj"
	"github.com/nxdir-s/idlerpg/internal/ports"
)

const (
	ErrAuthToken    string = "missing or malformed token"
	ErrInvalidToken string = "invalid or expired token"
)

type ServerHandler func(http.ResponseWriter, *http.Request) error

func httpHandler(fn ServerHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func NewAuthHandler(adapter ports.AuthPort) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return authHandler(adapter, h)
	}
}

func authHandler(adapter ports.AuthPort, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			http.Error(w, ErrAuthToken, http.StatusUnauthorized)
			return
		}

		token := strings.TrimSpace(strings.Replace(auth, "Bearer", "", 1))
		if token == "" {
			http.Error(w, ErrAuthToken, http.StatusForbidden)
			return
		}

		if err := adapter.ValidateToken(r.Context(), token, valobj.AccessToken); err != nil {
			http.Error(w, ErrInvalidToken, http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
