package tcp_server

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

/* tcp 方式
开启服务
循环接收数据
协程处理数据包
*/

//相关url https://blog.csdn.net/cqims21/article/details/104740507
//常规无任何处理方式
func NormalPackageTest() {
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
			go DealNormalPackageHandler(conn)
		}
	}
}

func DealNormalPackageHandler(conn net.Conn) {
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

//数据包添加信息（长度和校验值）处理方式
func DefinePackageTest() {
	listener, err := net.Listen("tcp", "localhost:8500")
	if err != nil {
		fmt.Println("tcp server start error:", err.Error())
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("tcp receive error", err.Error())
		}
		go DealDefinePackageHandler(conn)
	}
}

/*
 使用4个字节存取表头表示 + 2个字节存取数据长度
*/
func packageSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if !atEOF && len(data) > 6 && binary.BigEndian.Uint32(data[0:4]) == 0x123456 {
		var datalen int16
		binary.Read(bytes.NewReader(data[4:6]), binary.BigEndian, &datalen)
		var packagelen = int(datalen) + 6

		//大于怎么办
		if packagelen < len(data) {
			return packagelen, data[:packagelen], nil
		}
	}
	return
}

func DealDefinePackageHandler(conn net.Conn) {
	defer conn.Close()
	defer fmt.Println("连接关闭")

	dataBuf := bytes.NewBuffer(nil)
	// 由于 标识数据包长度 的只有两个字节 故数据包最大为 2^16+4(魔数)+2(长度标识)
	var tmpBuf [65542]byte
	for {
		//为什么不直接写道dataBuf里呢
		dataLen, err := conn.Read(tmpBuf[0:])
		dataBuf.Write(tmpBuf[0:dataLen])
		if err != nil {
			if err == io.EOF {
				continue
			}
			fmt.Println("deal package error:", err.Error())
		}
		scanner := bufio.NewScanner(dataBuf)
		scanner.Split(packageSplitFunc)
		for scanner.Scan() {
			fmt.Println("receive msg :", string(scanner.Bytes()[6:]))
		}
	}
	dataBuf.Reset()
}
