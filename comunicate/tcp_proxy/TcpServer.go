package tcp_proxy

import (
	"context"
	"errors"
	"net"
	"sync/atomic"
)

var (
	ErrServerClosed     = errors.New("tcp: Server Closed")
	ErrAbortHandler     = errors.New("tcp: abort TCPHandler")
	ServerContextKey    = &contextKey{name: "tcp-server"}
	LocalAddrContextKey = &contextKey{"local-addr"}
)

type TcpServer struct {
	Addr       string
	BaseCtx    context.Context
	inShutdown int32
	doneChan   chan struct{}
	l          *onceCloseListener
}

func (serv *TcpServer) shuttingDown() bool {
	return atomic.LoadInt32(&serv.inShutdown) != 0
}

func (serv *TcpServer) Close() error {
	atomic.StoreInt32(&serv.inShutdown, 1)
	close(serv.doneChan)
	serv.l.Close()
	return nil
}

func (serv *TcpServer) Serve(listener net.Listener) error {
	serv.l = &onceCloseListener{Listener: listener}
	defer serv.Close()
	if serv.BaseCtx == nil {
		serv.BaseCtx = context.Background()
	}
	baseContext := serv.BaseCtx
	ctx := context.WithValue(baseContext, ServerContextKey, serv)
	for {
		rw, err := listener.Accept()
		if err != nil {

		}
		if ctx == nil || rw == nil {

		}
	}

	return nil
}

func (serv *TcpServer) ListenAndServe() error {
	if serv.shuttingDown() {
		return ErrServerClosed
	}

	addr := serv.Addr
	if addr == "" {
		return errors.New("tcp server can not empty")
	}

	// listener, err := net.Listen("tcp", addr)
	// if err != nil {
	// 	return err
	// }
	return nil
}
