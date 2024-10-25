package epoll

import (
	"context"
	"fmt"
	"net"
	"os"
	"reflect"
	"syscall"

	"golang.org/x/sys/unix"
)

type EpollError struct {
	err error
}

func (e *EpollError) Error() string {
	return "failed to create epoll: " + e.err.Error()
}

type Epoll struct {
	fd          int
	connections map[int]net.Conn

	Add    chan net.Conn
	Remove chan net.Conn
}

func NewEpoll() (*Epoll, error) {
	fd, err := unix.EpollCreate1(0)
	if err != nil {
		return nil, &EpollError{err}
	}

	return &Epoll{
		fd:          fd,
		connections: make(map[int]net.Conn),
		Add:         make(chan net.Conn),
		Remove:      make(chan net.Conn),
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

	e.connections[fd] = conn

	if len(e.connections)%100 == 0 {
		fmt.Fprintf(os.Stdout, "total number of connections: %d\n", len(e.connections))
	}

	return nil
}

func (e *Epoll) remove(conn net.Conn) error {
	fd := e.getFileDescriptor(conn)

	err := unix.EpollCtl(e.fd, syscall.EPOLL_CTL_DEL, fd, nil)
	if err != nil {
		return err
	}

	delete(e.connections, fd)

	if len(e.connections)%100 == 0 {
		fmt.Fprintf(os.Stdout, "total number of connections: %d", len(e.connections))
	}

	return nil
}

func (e *Epoll) Wait() ([]net.Conn, error) {
	events := make([]unix.EpollEvent, 100)
	n, err := unix.EpollWait(e.fd, events, 100)
	if err != nil {
		return nil, err
	}

	// e.lock.RLock()
	// defer e.lock.RUnlock()

	var connections []net.Conn
	for i := 0; i < n; i++ {
		conn := e.connections[int(events[i].Fd)]
		connections = append(connections, conn)
	}

	return connections, nil
}

func (e *Epoll) getFileDescriptor(conn net.Conn) int {
	tcpConn := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn")
	fdVal := tcpConn.FieldByName("fd")
	pfdVal := reflect.Indirect(fdVal).FieldByName("pfd")

	return int(pfdVal.FieldByName("Sysfd").Int())
}
