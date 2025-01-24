package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/nxdir-s/idlerpg/internal/ports"
	"golang.org/x/oauth2"
)

const (
	StateCookieName string = "oauthstate"
	LoginURL        string = "http://localhost:4000/login"
	RedirectURL     string = "http://localhost:8080/auth/google/callback"
	GoogleOAuthURL  string = "https://www.googleapis.com/oauth2/v2/userinfo"
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

type Server struct {
	auth   ports.AuthPort
	mux    http.ServeMux
	google *oauth2.Config
}

func NewServer(ctx context.Context, adapter ports.AuthPort) (*Server, error) {
	s := &Server{
		auth:   adapter,
		google: googleConfig(),
	}

	s.mux.HandleFunc("GET /auth/google/login", enableCORS(httpHandler(s.googleLogin)))
	s.mux.HandleFunc("GET /auth/google/callback", enableCORS(httpHandler(s.googleCallback)))

	return s, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) googleLogin(w http.ResponseWriter, r *http.Request) error {
	verifier := generateVerifier(w)

	url := s.google.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

	return nil
}

func (s *Server) googleCallback(w http.ResponseWriter, r *http.Request) error {
	oauthState, err := r.Cookie(StateCookieName)
	if err != nil {
		return &ErrReadCookie{err}
	}

	token, err := s.google.Exchange(r.Context(), r.FormValue("code"), oauth2.VerifierOption(oauthState.Value))
	if err != nil {
		return err
	}

	resp, err := s.google.Client(r.Context(), token).Get(GoogleOAuthURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var userInfo GoogleUserInfo
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		return err
	}

	fmt.Fprintf(w, "UserInfo: %+v\n", userInfo)

	// Get/Create User
	// Respond with JWT

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
