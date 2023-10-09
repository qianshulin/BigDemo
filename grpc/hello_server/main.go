package main

import (
	sd "BigDemo/grpc/hello_server/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	sd.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, rep *sd.HelloRequest) (*sd.HelloResponse, error) {
	msg := &sd.HelloResponse{
		ResponseMsg: "hello" + rep.RequestName,
	}
	return msg, nil
}

func main() {
	//开启端口
	listen, _ := net.Listen("tcp", ":9090")
	//创建grpc服务
	grpcServer := grpc.NewServer()
	//在grpc服务中去注册编写的服务
	sd.RegisterSayHelloServer(grpcServer, &server{})
	//启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
