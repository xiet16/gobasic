package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"time"
)

func NormalClient() {

}

/*
  头部添加4个字节的首部校验(0x123456)
  然后添加2个字节的长度
*/
func DefinePackageClient() {
	conn, err := net.DialTimeout("tcp", "localhost:8503", time.Second*30)
	if err != nil {
		fmt.Println("tcp client start error:", err.Error)
	}

	for {
		packgeBuf := GetSendMsg()
		_, err := conn.Write(packgeBuf.Bytes())
		if err != nil {
			fmt.Println("tcp client send error:", err.Error())
		}
	}
}

func GetSendMsg() *bytes.Buffer {
	var str string
	num := rand.Intn(6) + 5
	for i := 1; i <= num; i++ {
		str += string(i)
	}
	content := []byte(str)
	cLen := len(content)
	headerBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(headerBuf, 0x123456)
	lengthBuf := make([]byte, 2)
	binary.BigEndian.PutUint16(lengthBuf, uint16(cLen))
	packgeBuf := bytes.NewBuffer(headerBuf)
	packgeBuf.Write(lengthBuf)
	packgeBuf.Write(content)
	return packgeBuf
}

func main() {

}
