package grpc_proxy

import (
	context "context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

/*
grpc
rpc (remote producer call) ：是一个远程调用协议，可以解决不同进程间的调用问题（个人理解主要是方便不同开发语言的接口对接）
rpc 使用http2 ，htt2 是http1.1 的扩展，不是替代，客户端会主动做尝试升级，如果失败，就使用http1.1 通信
http2 主要是为了解决延迟问题，使用了多路复用和流、帧的传输方式，速度大概是http1.1 的两倍
*/

type GrpcServer struct{}

func (g *GrpcServer) Count(context.Context, *CountRequest) (*CountResponse, error) {
	return &CountResponse{Message: "a + b response"}, nil
}

func GrpcServerRun() {
	listener, err := net.Listen("tcp", ":2013")
	if err != nil {
		log.Println("grpc server listen error :", err.Error())
		return
	}

	//创建grpc 服务器
	gs := grpc.NewServer()
	//注册服务
	RegisterGreeterServer(gs, &GrpcServer{})
	reflection.Register(gs)
	err = gs.Serve(listener)
	if err != nil {
		fmt.Printf("开启服务失败: %s", err)
		return
	}
}
