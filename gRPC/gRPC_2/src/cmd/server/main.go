package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	hellopb "mygrpc/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// GreetingServiceServer の実装
type myServer struct {
	hellopb.UnimplementedGreetingServiceServer
}

// Hello はHelloRequestを受け取り、HelloResponseを返す
func (s *myServer) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	log.Printf("Hello() called with: %v", req)
	return &hellopb.HelloResponse{Message: "Hello " + req.GetName()}, nil
}

// myServer のコンストラクタ
func NewMyServer() *myServer {
	return &myServer{}
}

func main() {
	// リスナー作成
	port := 8080
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	// gRPCサーバー作成
	s := grpc.NewServer()

	// gRPCサーバーにGreetingServiceを登録
	hellopb.RegisterGreetingServiceServer(s, NewMyServer())

	// サーバーリフレクション有効化 (gPRCサーバからprotoファイルの定義を取得できるようになる)
	// gRPCurlはprotoファイルの定義を知らないので、サーバから定義を取得する必要がある
	reflection.Register(s)

	// サーバー起動
	go func() {
		log.Printf("start gPRC server port: %v", port)
		s.Serve(lis)
	}()

	// Ctrl + C で終了 (GracefulStop 保留中のリクエストが終わってから終了)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
