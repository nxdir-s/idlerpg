package primary

import (
	"context"

	"github.com/nxdir-s/idlerpg/internal/core/valobj"
	"github.com/nxdir-s/idlerpg/internal/ports"
)

type ErrInvalidToken struct {
	token string
}

func (e *ErrInvalidToken) Error() string {
	return "recieved invalid token type " + e.token
}

type AuthAdapter struct {
	jwt ports.JWTPort
}

func NewAuthAdapter(ctx context.Context, jwt ports.JWTPort) (*AuthAdapter, error) {
	return &AuthAdapter{
		jwt: jwt,
	}, nil
}

// ValidateToken returns an error if the supplied token is not valid
func (a *AuthAdapter) ValidateToken(ctx context.Context, token string, tokenType string) error {
	switch tokenType {
	case valobj.AccessToken:
		if err := a.jwt.ValidAccessToken(ctx, token); err != nil {
			return err
		}
	case valobj.RefreshToken:
		if err := a.jwt.ValidRefreshToken(ctx, token); err != nil {
			return err
		}
	default:
		return &ErrInvalidToken{tokenType}
	}

	return nil
}
