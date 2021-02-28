package websocket_reverse_proxy

import (
	"log"
	"net/http"

	"xiet16.com/golearn/comunicate/load_balance"
	"xiet16.com/golearn/comunicate/middleware"
	"xiet16.com/golearn/comunicate/proxy"
)

var (
	wsporxyAddr = "127.0.0.1:2002"
)

func WebsocketProxyTest() {
	rb := load_balance.LoadBalanceFactory(load_balance.LbWeightRoundRobin)
	rb.Add("http://localhost:2003", "50")
	wsProxy := proxy.NewBalanceLoadBalanceReverseProxy(&middleware.SliceRouterContext{}, rb)
	log.Println("starting httpserver at" + wsporxyAddr)
	log.Fatal(http.ListenAndServe(wsporxyAddr, wsProxy))
}
