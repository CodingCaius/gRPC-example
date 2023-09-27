## 环境准备

> Protocol Buffer官方文档：https://protobuf.dev/getting-started/gotutorial/
>
> protobuf下载：
>
> https://github.com/protocolbuffers/protobuf
>
> 安装gRPC核心库：
>
> ```go
> go get google.golang.org/grpc
> ```
>
> 安装go语言的代码生成工具
>
> ```go
> go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
> go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
> ```




> 生成
>
> ```
> protoc --go_out=. hello.proto  // 生成go代码
> protoc --go-grpc_out=. hello.proto // 生成gRPC的代码，用来调用
> ```
>
> 或者同时生成hello.pb.go 和 hello_grpc.pb.go
>
> ```
> protoc --go-grpc_out=. --go_out=. hello.proto
> ```
>
> 




