package go_channel

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func ChannelTimerTest() {
	var h Handler
	h = "timer test"
	timer := &ChannelTimer{
		Interval: 2,
		Handler:  h.CallBack,
	}

	timer.Start()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

type Handler string

func (h *Handler) CallBack(data interface{}) {
	fmt.Println("timer callback")
}
