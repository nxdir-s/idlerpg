package server

import (
	"context"
	"log/slog"
	"net"
	"syscall"

	"github.com/nxdir-s/idlerpg/internal/core/entity"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sys/unix"
)

const (
	ClientBuffer int = 100000
)

type ErrEpoll struct {
	err error
}

func (e *ErrEpoll) Error() string {
	return "error creating epoll: " + e.err.Error()
}

type Epoll struct {
	fd     int
	pool   *Pool
	tracer trace.Tracer
	logger *slog.Logger

	Add    chan net.Conn
	Remove chan *Client
}

// NewEpoll creates an Epoll
func NewEpoll(pool *Pool, logger *slog.Logger, tracer trace.Tracer) (*Epoll, error) {
	fd, err := unix.EpollCreate1(0)
	if err != nil {
		return nil, &ErrEpoll{err}
	}

	return &Epoll{
		fd:     fd,
		pool:   pool,
		tracer: tracer,
		logger: logger,
		Add:    make(chan net.Conn, ClientBuffer),
		Remove: make(chan *Client, ClientBuffer),
	}, nil
}

// Start adds and removes connections from Epoll
func (e *Epoll) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case conn := <-e.Add:
			if err := e.add(ctx, conn); err != nil {
				e.logger.Error("failed to add connection", slog.Any("err", err))
			}
		case client := <-e.Remove:
			if err := e.remove(ctx, client); err != nil {
				e.logger.Error("failed to remove connection", slog.Any("err", err))
			}
		}
	}
}

// Wait listens for new epoll events
func (e *Epoll) Wait() ([]*Client, error) {
	events := make([]unix.EpollEvent, 100)

	_, err := unix.EpollWait(e.fd, events, 100)
	if err != nil {
		if err == unix.EINTR {
			return nil, nil
		}

		return nil, err
	}

	event := &EpollEvent{
		Events: events,
		Resp:   make(chan []*Client),
	}

	e.pool.EpollEvents <- event

	return <-event.Resp, nil
}

func (e *Epoll) add(ctx context.Context, conn net.Conn) error {
	ctx, span := e.tracer.Start(ctx, "add.client")
	defer span.End()

	fd, err := e.getFileDescriptor(conn)
	if err != nil {
		return err
	}

	if err := unix.SetNonblock(fd, true); err != nil {
		return err
	}

	event := unix.EpollEvent{
		Events: unix.POLLIN | unix.POLLHUP,
		Fd:     int32(fd),
	}

	if err := unix.EpollCtl(e.fd, unix.EPOLL_CTL_ADD, fd, &event); err != nil {
		return err
	}

	e.pool.Register <- &Client{
		Conn: conn,
		Fd:   int32(fd),
		User: entity.NewUser(),
	}

	return nil
}

func (e *Epoll) remove(ctx context.Context, client *Client) error {
	ctx, span := e.tracer.Start(ctx, "remove.client")
	defer span.End()

	fd, err := e.getFileDescriptor(client.Conn)
	if err != nil {
		return err
	}

	e.pool.Remove <- client.Fd

	if err := unix.EpollCtl(e.fd, unix.EPOLL_CTL_DEL, fd, nil); err != nil {
		return err
	}

	return nil
}

func (e *Epoll) getFileDescriptor(conn net.Conn) (int, error) {
	rawConn, err := conn.(syscall.Conn).SyscallConn()
	if err != nil {
		return 0, err
	}

	var sfd int
	err = rawConn.Control(func(fd uintptr) {
		sfd = int(fd)
	})
	if err != nil {
		return 0, err
	}

	return sfd, nil
}
