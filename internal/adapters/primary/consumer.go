package primary

import (
	"context"

	"github.com/nxdir-s/idlerpg/internal/ports"
	"github.com/nxdir-s/idlerpg/protobuf"
	"go.opentelemetry.io/otel/trace"
)

type ErrUserEvent struct {
	err error
}

func (e *ErrUserEvent) Error() string {
	return "failed to process user event: " + e.err.Error()
}

type ConsumerAdapter struct {
	events ports.Events
	tracer trace.Tracer
}

func NewConsumerAdapter(events ports.Events, tracer trace.Tracer) (*ConsumerAdapter, error) {
	return &ConsumerAdapter{
		events: events,
		tracer: tracer,
	}, nil
}

func (a *ConsumerAdapter) ProcessUserEvent(ctx context.Context, event *protobuf.UserEvent) error {
	ctx, span := a.tracer.Start(ctx, "adapter process.userevent")
	defer span.End()

	if err := a.events.HandleUserEvent(ctx, event); err != nil {
		return &ErrUserEvent{err}
	}

	return nil
}
