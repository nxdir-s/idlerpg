package logs

import (
	"context"
	"log/slog"
)

type LogCtxKey struct{}

type ContextHandler struct {
	slog.Handler
}

func NewHandler(handler slog.Handler) *ContextHandler {
	return &ContextHandler{handler}
}

func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	if attrs, ok := ctx.Value(LogCtxKey{}).([]slog.Attr); ok {
		for _, v := range attrs {
			r.AddAttrs(v)
		}
	}

	return h.Handler.Handle(ctx, r)
}

func AddCtxAttr(ctx context.Context, attr slog.Attr) context.Context {
	if attrs, ok := ctx.Value(LogCtxKey{}).([]slog.Attr); ok {
		attrs = append(attrs, attr)

		return context.WithValue(ctx, LogCtxKey{}, attrs)
	}

	attrs := []slog.Attr{}
	attrs = append(attrs, attr)

	return context.WithValue(ctx, LogCtxKey{}, attrs)
}
