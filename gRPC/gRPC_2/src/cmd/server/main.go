package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"time"

	hellopb "mygrpc/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// GreetingServiceServer の実装
type myServer struct {
	hellopb.UnimplementedGreetingServiceServer
}

// Unary RPCのメソッドの実装
func (s *myServer) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	log.Printf("Hello() called with: %v", req)
	return &hellopb.HelloResponse{Message: "Hello " + req.GetName()}, nil
}

// Server Streaming RPCのメソッドの実装
func (s *myServer) HelloServerStream(req *hellopb.HelloRequest, stream hellopb.GreetingService_HelloServerStreamServer) error {
	// サーバーからクライアントへ5回レスポンスを送信
	resCount := 5
	for i := 0; i < resCount; i++ {
		if err := stream.Send(&hellopb.HelloResponse{
			Message: fmt.Sprintf("[%d] Hello %s", i, req.GetName()),
		}); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	// メソッドの終了がストリームの終了を意味する
	return nil
}

// Client Streaming RPCのメソッドの実装
func (s *myServer) HelloClientStream(stream hellopb.GreetingService_HelloClientStreamServer) error {
	nameList := make([]string, 0)
	// クライアントからのリクエストを受信
	for {
		req, err := stream.Recv()
		// クライアントからのリクエストが終了したら、レスポンスを返して終了
		if errors.Is(err, io.EOF) {
			message := fmt.Sprintf("Hello %s", strings.Join(nameList, ", "))
			return stream.SendAndClose(&hellopb.HelloResponse{Message: message})
		}
		if err != nil {
			return err
		}
		nameList = append(nameList, req.GetName())
	}
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
