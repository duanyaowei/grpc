package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
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

type TokenServer struct {
	pb.UnimplementedTokenHelloServer
}

func (t TokenServer) TokenHello(ctx context.Context, req *pb.TokenRequest) (*pb.TokenResponse, error) {
	// 获取元数据信息
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("未传输 token")
	}
	var appId string
	var appKey string
	if v, ok := md["appid"]; ok {
		appId = v[0]
		fmt.Println("appID:", appId)
	}
	if v, ok := md["appkey"]; ok {
		appKey = v[0]
		fmt.Println("appKey:", appKey)
	}
	if appId != "kuangshen" || appKey != "123123" {
		return nil, errors.New("token 不正确")
	}

	fmt.Printf("hello" + req.RequestName)
	return &pb.TokenResponse{ResponseMsg: "hello" + req.RequestName}, nil
}

// -------------------------------初始生成测试-------------------------------

func ServerSayHello() {
	// 开启端口
	listen, _ := net.Listen("tcp", ":9090")
	// 创建 grpc 服务
	grpcServer := grpc.NewServer()
	// 在 grpc 服务中注册我们自己编写的服务
	pb.RegisterSayHelloServer(grpcServer, &server{})
	// 启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}

// -------------------------------新增测试-------------------------------

func DuanServerTest() {
	listen, _ := net.Listen("tcp", ":9001") //开启端口
	grpcServer := grpc.NewServer()          //创建grpc服务
	pb.RegisterDuanTestServer(grpcServer, &DuanSerer{})
	err := grpcServer.Serve(listen) //启动服务
	if err != nil {
		fmt.Printf("failed to serve: #{err}")
		return
	}
}

// ----------token认证 ：gRPC 提供了一个接口，位于credentials包下，需要客户端来实现-----

// type PerRPCCredentials interface {
//	GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
//	RequireTransportSecurity() bool
//}

func AuthToken() {
	// 开启端口
	listen, _ := net.Listen("tcp", ":9002")
	// 创建 grpc 服务
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	// 在 grpc 服务中注册我们自己编写的服务
	pb.RegisterTokenHelloServer(grpcServer, &TokenServer{})
	// 启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}

}

func main() {
	AuthToken()
}
