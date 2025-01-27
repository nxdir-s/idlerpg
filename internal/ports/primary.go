package ports

import "context"

type AuthPort interface {
	ValidateToken(ctx context.Context, token string, tokenType string) error
}
