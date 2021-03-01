package grpc_proxy

import (
	context "context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

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
