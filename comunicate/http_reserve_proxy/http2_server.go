package http_reserve_proxy

import (
	"log"
	"net/http"
	"time"
)

/*
http2 必须是https
http2 多路复用
流:
帧
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
