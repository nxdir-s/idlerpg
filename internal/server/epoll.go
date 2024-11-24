package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"syscall"

	"github.com/nxdir-s/idlerpg/internal/core/entity"
	"github.com/nxdir-s/idlerpg/internal/core/valobj"
	"golang.org/x/sys/unix"
)

type EpollError struct {
	err error
}

func (e *EpollError) Error() string {
	return "error creating epoll: " + e.err.Error()
}

type Epoll struct {
	fd   int
	pool *Pool

	Add    chan net.Conn
	Remove chan *Client
}

// NewEpoll creates an Epoll
func NewEpoll(pool *Pool) (*Epoll, error) {
	fd, err := unix.EpollCreate1(0)
	if err != nil {
		return nil, &EpollError{err}
	}

	return &Epoll{
		fd:     fd,
		pool:   pool,
		Add:    make(chan net.Conn),
		Remove: make(chan *Client),
	}, nil
}

func (e *Epoll) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case conn := <-e.Add:
			if err := e.add(conn); err != nil {
				fmt.Fprintf(os.Stdout, "failed to add connection: %+v\n", err)
			}
		case client := <-e.Remove:
			if err := e.remove(client); err != nil {
				fmt.Fprintf(os.Stdout, "failed to remove connection: %+v\n", err)
			}
		}
	}
}

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

func (e *Epoll) add(conn net.Conn) error {
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
		Conn:   conn,
		Fd:     fd,
		Msgs:   make(chan *valobj.Message),
		Player: entity.NewPlayer(),
	}

	return nil
}

func (e *Epoll) remove(client *Client) error {
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

	sfd := 0

	err = rawConn.Control(func(fd uintptr) {
		sfd = int(fd)
	})
	if err != nil {
		return 0, err
	}

	return sfd, nil
}
