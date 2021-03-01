package tcp_proxy

import (
	"net"
)

/*
构建tcp 服务器
1 监听服务
2 获取构建新连接对象并设置超时时间以及keepalive
3 设置方法退出时连接关闭
4 调用回调接口
*/

type tcpKeepAliveListener struct {
	*net.TCPListener
}

// func (ln tcpKeepAliveListener) Accept() (net.Conn, error) {
// 	tc, err := ln.AcceptTCP()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return tc, nil
// }
