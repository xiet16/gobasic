package http_reserve_proxy

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
创建一个两个http 服务器，分别使用两个端口，当做代理服务器的下游
这里继承、封装一下
没有什么是加一层解决不了的，如果有，就再加一层
*/

type RealServer struct {
	Addr string
}

//不好的写法,这样SayHelloHandler和SayHelloHandler 没有变量
// func (server *RealServer) Run_Bad() {
// 	http.HandleFunc("/", SayHelloHandler)
// 	http.HandleFunc("/base/error", SayHelloHandler)
// 	http.ListenAndServe(server.Addr, nil)
// }

func RealStart() {
	s1 := &RealServer{Addr: "127.0.0.1:2003"}
	s1.Run()
	s2 := &RealServer{Addr: "127.0.0.1:2004"}
	s2.Run()
	//监听关闭信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func (r *RealServer) Run() {
	log.Println(r.Addr + " start")
	mux := http.NewServeMux()
	mux.HandleFunc("/", r.SayHelloHandler)
	mux.HandleFunc("/base/error", r.LogErrorHandler)
	server := &http.Server{
		Addr:         r.Addr,
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()
}

func (server *RealServer) SayHelloHandler(w http.ResponseWriter, r *http.Request) {
	upath := fmt.Sprintf("http://%s%s\n", server.Addr, r.URL.Path)
	io.WriteString(w, upath)
}

func (server *RealServer) LogErrorHandler(w http.ResponseWriter, r *http.Request) {
	upath := "errorlog path"
	w.WriteHeader(500)
	io.WriteString(w, upath)
}
