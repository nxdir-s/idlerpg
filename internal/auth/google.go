package auth

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleUserInfo struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Verified  bool   `json:"verified_email"`
	Name      string `json:"name"`
	FirstName string `json:"given_name"`
	LastName  string `json:"family_name"`
	Picture   string `json:"picture"`
}

func googleConfig(cId string, secret string) *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  RedirectURL,
		ClientID:     cId,
		ClientSecret: secret,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}
