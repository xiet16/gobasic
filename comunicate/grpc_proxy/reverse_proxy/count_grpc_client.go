package grpc_proxy

import (
	"context"
	"fmt"
	"log"

	grpc "google.golang.org/grpc"
)

func GrpcClientRun() {
	conn, err := grpc.Dial(":2013", grpc.WithInsecure())
	if err != nil {
		log.Println("get grpc connection error: ", err.Error())
		return
	}
	defer conn.Close()

	//创建客户端
	client := NewGreeterClient(conn)
	response, err := client.Count(context.Background(), &CountRequest{Name: "xiet grpc"})
	if err != nil {
		log.Println("get count result error: ", err.Error())
	}

	fmt.Printf("get message: %v", response.Message)
}
