package tcp_server

import (
	"bytes"
	"fmt"
	"io"
	"net"
)

//相关url https://blog.csdn.net/cqims21/article/details/104740507
func ErrorPackageTest() {
	fmt.Println("tcp server start")
	listener, err := net.Listen("", "localhost:8500")
	if err != nil {
		fmt.Println("listener err", err.Error())
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("data receive error", err.Error())
		} else {
			go DealErrorPackageHandler(conn)
		}
	}
}

func DealErrorPackageHandler(conn net.Conn) {
	defer conn.Close()
	defer fmt.Println("连接关闭")
	fmt.Println("新连入地址:", conn.RemoteAddr)
	tmpdata := bytes.NewBuffer(nil)
	var buf [1024]byte
	for {
		len, err := conn.Read(buf[0:])
		tmpdata.Write(buf[0:len])
		if err != nil {
			if err == io.EOF {
				continue
			} else {
				fmt.Println("数据读取错误:", err.Error())
				break
			}
		} else {
			fmt.Println("接收到的数据:", tmpdata.String)
		}
	}
	tmpdata.Reset()
}
