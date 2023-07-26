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

func main() {
	//start := time.Now() // 记录开始时间
	//cc, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	log.Fatalf("did not connect: %v", err)
	//}
	//defer cc.Close()
	//// 建立连接
	//client := pb.NewSayHelloClient(cc)
	//// 执行 rpc 调用（这个方法在服务器端来实现并返回结果）
	//resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "kuangshen"})
	//fmt.Println(resp.GetResponseMsg())
	//elapsed := time.Since(start) // 计算经过的时间
	//fmt.Println("程序运行耗时:", elapsed)

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
