package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"reflect"
	"syscall"

	"github.com/nxdir-s/IdleRpg/internal/core/entity"
	"github.com/nxdir-s/IdleRpg/internal/core/valobj"
	"golang.org/x/sys/unix"
)

type EpollError struct {
	err error
}

func (e *EpollError) Error() string {
	return "failed to create epoll: " + e.err.Error()
}

type Epoll struct {
	fd   int
	pool *Pool

	Add    chan net.Conn
	Remove chan net.Conn
}

func NewEpoll(pool *Pool) (*Epoll, error) {
	fd, err := unix.EpollCreate1(0)
	if err != nil {
		return nil, &EpollError{err}
	}

	return &Epoll{
		fd:     fd,
		pool:   pool,
		Add:    make(chan net.Conn),
		Remove: make(chan net.Conn),
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
		case conn := <-e.Remove:
			if err := e.remove(conn); err != nil {
				fmt.Fprintf(os.Stdout, "failed to remove connection: %+v\n", err)
			}
		}
	}
}

func (e *Epoll) add(conn net.Conn) error {
	fd := e.getFileDescriptor(conn)

	err := unix.EpollCtl(e.fd, syscall.EPOLL_CTL_ADD, fd, &unix.EpollEvent{Events: unix.POLLIN | unix.POLLHUP, Fd: int32(fd)})
	if err != nil {
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

func (e *Epoll) remove(conn net.Conn) error {
	fd := e.getFileDescriptor(conn)

	err := unix.EpollCtl(e.fd, syscall.EPOLL_CTL_DEL, fd, nil)
	if err != nil {
		return err
	}

	e.pool.Unregister <- fd

	return nil
}

func (e *Epoll) Wait() ([]*Client, error) {
	events := make([]unix.EpollEvent, 0, 100)

	_, err := unix.EpollWait(e.fd, events, 100)
	if err != nil {
		return nil, err
	}

	event := &EpollEvent{
		Events: events,
		Resp:   make(chan []*Client),
	}

	e.pool.EpollEvents <- event

	return <-event.Resp, nil
}

func (e *Epoll) getFileDescriptor(conn net.Conn) int {
	tcpConn := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn")
	fdVal := tcpConn.FieldByName("fd")
	pfdVal := reflect.Indirect(fdVal).FieldByName("pfd")

	return int(pfdVal.FieldByName("Sysfd").Int())
}
