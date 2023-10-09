package main

import (
	sd "BigDemo/grpc/hello_server/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	//连接到server端,此处禁用安全传输，没有加密和验证
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	//建立连接
	client := sd.NewSayHelloClient(conn)

	// 执行rpc调用(这个方法在服务器端来实现并返回结果)
	resp, _ := client.SayHello(context.Background(), &sd.HelloRequest{RequestName: "lsr"})

	fmt.Println(resp.GetResponseMsg())
}
