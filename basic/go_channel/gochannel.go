package go_channel

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
channel : 通道(管道)
特点：
同一时间点，只能被一个协程访问,也就是说它是同步的
管道分为有缓冲和无缓冲
无缓冲情况下：在接收者准备好之前，发送者是阻塞的也就是说，首先得有接收者开始接收，才会开始发送
有缓冲情况下：在buf 满之后，直到有接收者接收，发送者才可以发送
*/

func GoroutinueChannel() {
	c := make(chan string, 5)

	go func() {
		send1(c)
	}()

	go func() {
		send2(c)
	}()
	for {
		select {
		case rev := <-c:
			fmt.Println("receive data: ", rev)
			time.Sleep(time.Second * 2)
		}
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func send1(c chan string) {
	for i := 0; i < 100; i++ {
		//fmt.Println("send1 send data before:", i)
		c <- fmt.Sprint("send1_", i)
		fmt.Println("send1 send data success:", i)
	}
}

func send2(c chan string) {
	for i := 0; i < 100; i++ {
		//fmt.Println("send2 send data before:", i)
		c <- fmt.Sprint("send2_", i)
		fmt.Println("send2 send data success:", i)
	}
}
