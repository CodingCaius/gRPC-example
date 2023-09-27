package main

import (
	"context"
	"gRPC/proto/search"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type service struct {
	search.UnimplementedSearchServiceServer
}

func (s *service) Search(ctx context.Context, req *search.SearchRequest) (res *search.SearchResponse, err error) {
	// 这个函数是 gRPC 服务的实际实现，当客户端调用 Search 服务时，将执行这个函数。
    // 它接收一个上下文（context）对象，用于管理请求的生命周期，以及客户端发送的请求对象 req。
    // 在这个示例中，它简单地将客户端请求的内容加上 " Server" 后返回给客户端。

	//fmt.Println(req.GetRequest())
	return &search.SearchResponse{Response: req.GetRequest() + " Server"}, nil
}

const PORT = "8888"

func main() {
	// 创建了 gRPC 服务器的 TLS 凭证（credentials）对象，这个凭证需要服务器的证书和私钥
	// 根据服务端输入的证书文件和密钥构造 TLS 凭证
	c, err := credentials.NewServerTLSFromFile("D:\\GitSpace\\gRPC-example\\conf\\server_side_TLS\\server.pem", "D:\\GitSpace\\gRPC-example\\conf\\server_side_TLS\\server.key")
	if err != nil {
		log.Fatalf("credentials.NewServerTLSFromFile err: %v", err)
	}
	// 返回一个 ServerOption，用于设置服务器连接的凭据。
	// 用于 grpc.NewServer(opt ...ServerOption) 为 gRPC Server 设置连接选项
	// 创建一个新的 gRPC 服务器 s，并将 TLS 凭证应用于该服务器
	s := grpc.NewServer(grpc.Creds(c))
	lis, err := net.Listen("tcp", ":"+PORT) //创建 Listen，监听 TCP 端口
	if err != nil {
		log.Fatalf("credentials.NewServerTLSFromFile err: %v", err)
	}
	//将 SearchService（其包含需要被调用的服务端接口）注册到 gRPC Server 的内部注册中心。
	//这样可以在接受到请求时，通过内部的服务发现，发现该服务端接口并转接进行逻辑处理
	// 注册服务接口 search.RegisterSearchServiceServer，将其绑定到 gRPC 服务器上，以便服务器可以接受客户端的请求并调用 service 结构体中的 Search 方法
	search.RegisterSearchServiceServer(s, &service{})

	//gRPC Server 开始 lis.Accept，直到 Stop 或 GracefulStop
	// 通过 s.Serve(lis) 启动 gRPC 服务器，开始监听客户端的连接请求，处理客户端发送的 gRPC 调用请求
	s.Serve(lis)
}
