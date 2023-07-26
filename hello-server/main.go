package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	pb "stay_grpc/hello-server/proto"
)

type server struct {
	pb.UnimplementedSayHelloServer
}

func (s server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Printf("hello" + req.RequestName)
	return &pb.HelloResponse{ResponseMsg: "hello" + req.RequestName}, nil
}

type DuanSerer struct {
	pb.UnimplementedDuanTestServer
}

func (d DuanSerer) DuanTest(ctx context.Context, req *pb.DuanParams) (*pb.DuanResponse, error) {
	var duan = &pb.DuanParams{}
	duan.Name = req.Name
	duan.Age = req.Age
	duan.Address = req.Address

	return &pb.DuanResponse{ResponseMsg: "响应成功", ResponseCode: 200}, nil
}

func main() {

	//// 开启端口
	//listen, _ := net.Listen("tcp", ":9090")
	//// 创建 grpc 服务
	//grpcServer := grpc.NewServer()
	//// 在 grpc 服务中注册我们自己编写的服务
	//pb.RegisterSayHelloServer(grpcServer, &server{})
	//// 启动服务
	//err := grpcServer.Serve(listen)
	//if err != nil {
	//	fmt.Printf("failed to serve: %v", err)
	//	return
	//}

	//--------------------------duan test
	listen, _ := net.Listen("tcp", ":9001") //开启端口
	grpcServer := grpc.NewServer()          //创建grpc服务
	pb.RegisterDuanTestServer(grpcServer, &DuanSerer{})
	err := grpcServer.Serve(listen) //启动服务
	if err != nil {
		fmt.Printf("failed to serve: #{err}")
		return
	}

}
