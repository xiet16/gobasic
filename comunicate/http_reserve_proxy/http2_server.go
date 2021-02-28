package http_reserve_proxy

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
http2 必须是https
http2 多路复用
流:
帧
*/

/*
创建ca 私钥
openssl genrsa -out ca.key 2048
创建ca 证书
openssl req -x509 -new -nodes -key ca.key -subj "/CN=httpsexample.com" -days 5000 -out ca.crt
服务器私钥
openssl genrsa -out server.key 2048
//服务器证书签名请求
openssl req -new -key server.key -subj "/CN=httpsexample.com" -out server.csr
//用上面两个文件生成服务器证书(days是有效天数)
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 5000
*/

func (r *RealServer) Http2Run() {
	log.Println("start http2 server at:" + r.Addr)
	mux := http.NewServeMux()
	mux.HandleFunc("/", r.SayHelloHandler)
	mux.HandleFunc("/base/error", r.LogErrorHandler)
	server := &http.Server{
		Addr:         r.Addr,
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}
	go func ()  {
		http2.ConfigureServer(server,&http2.Server{})
		//log.Fatal(server.ListenAndServeTLS())
	}
}

func Start() {
	s1 := &RealServer{Addr: "http://localhost:3003"}
	s1.Http2Run()
	s2 := &RealServer{Addr: "http://localhost:3004"}
	s2.Http2Run()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
