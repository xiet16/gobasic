package http_reserve_proxy

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

/*
反向代理
多级反向代理
特点：
客户端不需要知道目标服务器的信息
代理服务器做转发
代理服务器可实现：
*/

/*
简单实现：
1 解析客户端地址，更改请求的协议和主机
2 向下游服务器请求
3 将返回信息返回客户端
httprequest:请求行+ 请求头+ 空行+ 请求体
httpresponse :
*/

var (
	proxy_Addr = "127.0.0.1:2003"
	port       = "2002"
)

func NormalReserveProxyHandler(w http.ResponseWriter, r *http.Request) {
	//解析目标url
	proxy, err := url.Parse(proxy_Addr)
	if err != nil {
		fmt.Println("proxy url analysis error:", err.Error())
		return
	}
	r.URL.Scheme = proxy.Scheme
	r.URL.Host = proxy.Host

	//请求下游
	transport := http.DefaultTransport
	respone, err := transport.RoundTrip(r)
	log.Print(r.RequestURI)
	if err != nil {
		fmt.Print(err)
		return
	}

	//返回请求内容
	for k, vs := range respone.Header {
		for _, v := range vs {
			w.Header().Add(k, v)
		}
	}

	defer respone.Body.Close()
	bufio.NewReader(respone.Body).WriteTo(w)
}

func ServerStart() {
	http.HandleFunc("/", NormalReserveProxyHandler)
	err := http.ListenAndServe("127.0.0.1:"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
