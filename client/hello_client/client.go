package main

import (
	"context"
	"fmt"
	"gRPC/proto/hello"
	"log"

	"google.golang.org/grpc"
)

const PORT = "8888"

func main() {
	// 建立链接
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	// 一定要记得关闭链接
	defer conn.Close()

	// 实例化客户端
	client := hello.NewUserServiceClient(conn)
	// 发起请求
	response, err := client.SayHi(context.Background(), &hello.Request{Name: "caius"})
	if err != nil {
		log.Fatalf("client.SayHi err: %v", err)
	}
	fmt.Printf("resp: %s", response.GetResult())
}
