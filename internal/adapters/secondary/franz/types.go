package franz

import (
	"context"

	"github.com/twmb/franz-go/pkg/kgo"
)

type Consumer interface {
	Process(ctx context.Context, record *kgo.Record) error
}
