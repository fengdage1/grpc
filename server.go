package main

import (
	"net"
	"fmt"
	pb "protobuf/grpctest" // 引入编译生成的包

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

var HelloService = helloService{}

type helloService struct{}


func (this helloService) SayHello(ctx context.Context,in *pb.HelloRequest)(*pb.HelloReply,error){
	resp := new(pb.HelloReply)
	resp.Message = "hello"+in.Name+"."
	return resp,nil
}


func main(){
	listen,err:=net.Listen("tcp",Address)
	if err != nil{
		grpclog.Fatalf("failed to listen: %v", err)
	}
	s:=grpc.NewServer()
	pb.RegisterHelloServer(s,HelloService)

	grpclog.Println("Listen on " + Address)
	s.Serve(listen)
	fmt.Println("1")
}
