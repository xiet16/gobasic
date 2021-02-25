package websocket_reverse_proxy

import "xiet16.com/golearn/comunicate/load_balance"

func WebsocketProxyTest() {
	rb := load_balance.LoadBalanceFactory(load_balance.LbWeightRoundRobin)
}
