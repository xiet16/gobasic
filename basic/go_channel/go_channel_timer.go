package go_channel

import (
	"fmt"
	"time"
)

/*
用channel 实现一个定时器

*/

type StatuType uint8

const (
	Stoped StatuType = iota
	Running
)

type TaskDetail struct {
	Message string
}

type ChannelTimer struct {
	Interval  int
	StatuType StatuType
	Handler   func(object interface{})
}

func NewChannelTimer() *ChannelTimer {
	return &ChannelTimer{
		StatuType: Stoped,
	}
}

func (t *ChannelTimer) Start() {
	if t.StatuType == Running {
		return
	}

	t.StatuType = Running
	c := make(chan bool)

	go t.check(c)

	for {
		select {
		case ok := <-c:
			if ok {
				if t.Handler != nil {
					fmt.Println("通道收到信息")
					t.Handler(&TaskDetail{Message: "定时时间"})
				}
			} else {
				t.Handler(&TaskDetail{Message: "定时时间"})
			}
		}
	}
}

func (t *ChannelTimer) check(c chan bool) {
	if t.Interval <= 0 {
		close(c)
	}

	//每次睡眠100毫秒
	count := t.Interval * 1000 / 200
	total := 0
	for {
		fmt.Println("定时循环中，当前计数：", total)

		if t.StatuType == Stoped {
			fmt.Println("定时器停止")
			return
		}

		time.Sleep(time.Millisecond * 200)
		total++

		if total%count == 0 {
			fmt.Println("通道发送信号")
			c <- true
		}

		if total == 100 {
			total = 1
		}
	}
}

func (t *ChannelTimer) Stop() {
	t.StatuType = Stoped
}
