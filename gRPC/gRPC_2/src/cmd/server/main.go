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

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// GreetingServiceServer の実装
type myServer struct {
	hellopb.UnimplementedGreetingServiceServer
}

// Unary RPCのメソッドの実装
func (s *myServer) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	log.Println("Hello() called with: ", req)

	// metadataの取得
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Println("metadata: ", md)
	}

	// サーバー側でmetadataを付与
	headerMD := metadata.New(map[string]string{"type": "unary", "from": "server", "in": "header"})
	if err := grpc.SendHeader(ctx, headerMD); err != nil {
		return nil, err
	}
	trailerMD := metadata.New(map[string]string{"type": "unary", "from": "server", "in": "trailer"})
	if err := grpc.SetTrailer(ctx, trailerMD); err != nil {
		return nil, err
	}

	// サーバー側でエラーを発生させる
	if req.GetName() == "error" {
		stat := status.New(codes.Unknown, "unknown error occurred")
		// エラーの詳細を追加
		stat, _ = stat.WithDetails(&errdetails.DebugInfo{
			Detail: "Hello() called with: " + req.GetName(),
		})
		err := stat.Err()
		return nil, err
	}

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

// Bidirectional Streaming RPCのメソッドの実装
func (s *myServer) HelloBiStreams(stream hellopb.GreetingService_HelloBiStreamsServer) error {
	// metadataの取得
	if md, ok := metadata.FromIncomingContext(stream.Context()); ok {
		log.Println("metadata: ", md)
	}

	// サーバー側でmetadataを付与
	headerMD := metadata.New(map[string]string{"type": "stream", "from": "server", "in": "header"})
	// すぐにheaderを送信する場合
	if err := stream.SendHeader(headerMD); err != nil {
		return err
	}
	// 本来headerを送るタイミングで送信する場合
	// if err := stream.SetHeader(headerMD); err != nil {
	// 	return err
	// }
	trailerMD := metadata.New(map[string]string{"type": "stream", "from": "server", "in": "trailer"})
	stream.SetTrailer(trailerMD)

	// 1リクエストに対して1レスポンスを返す
	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}
		message := fmt.Sprintf("Hello %s", req.GetName())
		if err := stream.Send(&hellopb.HelloResponse{Message: message}); err != nil {
			return err
		}
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
	s := grpc.NewServer(
		// インターセプターの登録
		grpc.ChainUnaryInterceptor(myUnaryServerInterceptor1, myUnaryServerInterceptor2),
		grpc.StreamInterceptor(myStreamServerInterceptor1),
	)

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
