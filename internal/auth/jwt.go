package auth

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ErrIssueToken struct {
	err error
}

func (e *ErrIssueToken) Error() string {
	return "failed to issue access token: " + e.err.Error()
}

type ErrRefreshToken struct {
	err error
}

func (e *ErrRefreshToken) Error() string {
	return "failed to issue refresh token: " + e.err.Error()
}

type ErrParseToken struct {
	err error
}

func (e *ErrParseToken) Error() string {
	return "failed to parse token: " + e.err.Error()
}

type ErrTokenExpired struct{}

func (e *ErrTokenExpired) Error() string {
	return "error token is expired"
}

type ErrSigningMethod struct {
	method any
}

func (e *ErrSigningMethod) Error() string {
	return fmt.Sprintf("unexpected signing method: %v", e.method)
}

type ErrInvalidToken struct{}

func (e *ErrInvalidToken) Error() string {
	return "error invalid token"
}

type JWT struct {
	accessKey  string
	refreshKey string
}

func NewJWT(accessKey string, refreshKey string) *JWT {
	return &JWT{
		accessKey:  accessKey,
		refreshKey: refreshKey,
	}
}

func (j *JWT) IssueToken(ctx context.Context, userId int, email string, expires time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.RegisteredClaims{
		Issuer:    "test",
		Subject:   strconv.Itoa(userId),
		ID:        email,
		ExpiresAt: jwt.NewNumericDate(expires),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
	})

	tok, err := token.SignedString([]byte(j.accessKey))
	if err != nil {
		return "", &ErrIssueToken{err}
	}

	return tok, nil
}

func (j *JWT) IssueRefreshToken(ctx context.Context, userId int, email string, expires time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.RegisteredClaims{
		Issuer:    "test",
		Subject:   strconv.Itoa(userId),
		ID:        email,
		ExpiresAt: jwt.NewNumericDate(expires),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
	})

	tok, err := token.SignedString([]byte(j.refreshKey))
	if err != nil {
		return "", &ErrRefreshToken{err}
	}

	return tok, nil
}

func (j *JWT) ValidAccessToken(ctx context.Context, tokenStr string) (bool, error) {
	token, err := parse(tokenStr, j.accessKey)
	if err != nil {
		return false, err
	}

	if err := validTime(token); err != nil {
		return false, err
	}

	return true, nil
}

// validTime returns an error if the token fails validations
func validTime(token *jwt.Token) error {
	claims, ok := token.Claims.(jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return &ErrInvalidToken{}
	}

	if time.Now().Unix() > claims.ExpiresAt.Unix() {
		return &ErrTokenExpired{}
	}

	return nil
}

func parse(tokenStr string, key string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &ErrSigningMethod{token.Header["alg"]}
		}

		return []byte(key), nil
	})
	if err != nil {
		return nil, &ErrParseToken{err}
	}

	return token, nil
}
