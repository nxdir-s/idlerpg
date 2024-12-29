package auth

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	LoginURL        string = "http://localhost:4000/login"
	RedirectURL     string = "http://localhost:8080/auth/google/callback"
	GoogleOAuthURL  string = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
	StateCookieName string = "oauthstate"
)

type ErrReadCookie struct {
	err error
}

func (e *ErrReadCookie) Error() string {
	return "error reading cookie: " + e.err.Error()
}

type ErrInvalidState struct {
	err error
}

func (e *ErrInvalidState) Error() string {
	return "invalid oauth state: " + e.err.Error()
}

type ServerHandler func(http.ResponseWriter, *http.Request) error

func httpHandler(fn ServerHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

type Server struct {
	mux       http.ServeMux
	googleCfg *oauth2.Config
}

func NewServer(ctx context.Context) (*Server, error) {
	s := &Server{
		googleCfg: googleConfig(),
	}

	s.mux.HandleFunc("/auth/google/login", httpHandler(s.googleLogin))
	s.mux.HandleFunc("/auth/google/callback", httpHandler(s.googleCallback))

	return s, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) googleLogin(w http.ResponseWriter, r *http.Request) error {
	verifier := generateVerifier(w)
	url := s.googleCfg.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

	return nil
}

func (s *Server) googleCallback(w http.ResponseWriter, r *http.Request) error {
	oauthState, err := r.Cookie(StateCookieName)
	if err != nil {
		return &ErrReadCookie{err}
	}

	if r.FormValue("state") != oauthState.Value {
		fmt.Fprint(os.Stdout, "invalid oauth state\n")
		http.Redirect(w, r, LoginURL, http.StatusPermanentRedirect)

		return nil
	}

	return nil
}

func generateVerifier(w http.ResponseWriter) string {
	verifier := oauth2.GenerateVerifier()

	http.SetCookie(w, &http.Cookie{
		Name:    StateCookieName,
		Value:   verifier,
		Expires: time.Now().Add(20 * time.Minute),
	})

	return verifier
}

func googleConfig() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  RedirectURL,
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}
