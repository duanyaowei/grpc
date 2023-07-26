package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "stay_grpc/hello-server/proto"
	"time"
)

type ClientTokenAuth struct {
}

func (c ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appId":  "kuangshen",
		"appKey": "123123",
	}, nil
}

func (c ClientTokenAuth) RequireTransportSecurity() bool {
	return false
}

// -------------------初始化测试--------------

func ServerHello() {
	start := time.Now() // 记录开始时间
	cc, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer cc.Close()
	// 建立连接
	client := pb.NewSayHelloClient(cc)
	// 执行 rpc 调用（这个方法在服务器端来实现并返回结果）
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "kuangshen"})
	fmt.Println(resp.GetResponseMsg())
	elapsed := time.Since(start) // 计算经过的时间
	fmt.Println("程序运行耗时:", elapsed)
}

// -------------------新增测试--------------

func DuanTest() {
	start := time.Now() // 记录开始时间
	cc, err := grpc.Dial("127.0.0.1:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer cc.Close()
	// 建立连接
	client := pb.NewDuanTestClient(cc)
	// 执行 rpc 调用（这个方法在服务器端来实现并返回结果）
	resp, _ := client.DuanTest(context.Background(), &pb.DuanParams{Name: "KD", Age: 23, Address: "美国超音速"})
	fmt.Println(resp.GetResponseCode())
	elapsed := time.Since(start) // 计算经过的时间
	fmt.Println("程序运行耗时:", elapsed)
}

// -------------------token认证--------------

func AuthToken() {
	start := time.Now() // 记录开始时间

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))

	cc, err := grpc.Dial("127.0.0.1:9002", opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer cc.Close()

	// 建立连接
	client := pb.NewTokenHelloClient(cc)
	// 执行 rpc 调用（这个方法在服务器端来实现并返回结果）
	resp, _ := client.TokenHello(context.Background(), &pb.TokenRequest{RequestName: "kuangshen"})
	fmt.Println(resp.GetResponseMsg())

	elapsed := time.Since(start) // 计算经过的时间
	fmt.Println("程序运行耗时:", elapsed)

}

func main() {
	AuthToken()
}
