package tcp_proxy

import (
	"net"
	"sync"
)

// type Listener interface {
// 	// Accept waits for and returns the next connection to the listener.
// 	Accept() (Conn, error)

// 	// Close closes the listener.
// 	// Any blocked Accept operations will be unblocked and return errors.
// 	Close() error

// 	// Addr returns the listener's network address.
// 	Addr() Addr
// }

//要实现接口的三个方法
type onceCloseListener struct {
	Listener net.Listener
	once     sync.Once
	CloseErr error
}

func (oc *onceCloseListener) Close() error {
	oc.once.Do(oc.close)
	return oc.CloseErr
}

func (oc *onceCloseListener) close() {
	oc.CloseErr = oc.Listener.Close()
}

//这个接口为什么它不实现
// func (l *onceCloseListener) Addr() net.Addr {

// }
