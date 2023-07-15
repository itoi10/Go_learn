package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	hellopb "mygrpc/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	scanner *bufio.Scanner
	client  hellopb.GreetingServiceClient
)

func main() {
	fmt.Println("start gRPC Client.")

	// 標準入力のスキャナーを作成
	scanner = bufio.NewScanner(os.Stdin)

	// gRPCサーバーとのコネクションを確立
	address := "localhost:8080"
	conn, err := grpc.Dial(
		address,
		// SSL/TLSを使わない
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// コネクションが確立されるまで待機する
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("Connection failed: %v", err)
		return
	}
	defer conn.Close()

	// gRPCクライアント生成 (この関数はprotoファイルから自動生成されている)
	client = hellopb.NewGreetingServiceClient(conn)

	for {
		fmt.Println("1: send Request(Unary)")
		fmt.Println("2: send Request(Server Streaming)")
		fmt.Println("9: exit")
		fmt.Print("-> ")

		// 標準入力から入力を受け取る
		scanner.Scan()
		in := scanner.Text()

		switch in {
		case "1":
			HelloUnary()
		case "2":
			HelloServerStream()

		case "9":
			fmt.Println("bye.")
			goto END
		}
	}

END:
}

func HelloUnary() {
	fmt.Println("input your name:")
	scanner.Scan()
	name := scanner.Text()

	// HelloRequest型のリクエストを作成
	req := &hellopb.HelloRequest{
		Name: name,
	}
	// Helloメソッドを実行し、HelloResponse型のレスポンスを受け取る
	res, err := client.Hello(context.Background(), req)
	if err != nil {
		log.Fatalf("HelloUnary failed: %v", err)
		return
	}

	fmt.Printf("Response: %v\n", res.GetMessage())
}

func HelloServerStream() {
	fmt.Println("input your name:")
	scanner.Scan()
	name := scanner.Text()

	req := &hellopb.HelloRequest{
		Name: name,
	}
	// HelloServerStreamメソッドを実行し、HelloResponse型のレスポンスを受け取る
	stream, err := client.HelloServerStream(context.Background(), req)
	if err != nil {
		log.Fatalf("HelloServerStream failed: %v", err)
		return
	}

	// サーバーからのレスポンスを受け取る
	for {
		res, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			// サーバーからのレスポンスが終了したら終了
			fmt.Println("all response received.")
			break
		}
		if err != nil {
			log.Fatalf("Response Recv failed: %v", err)
			return
		}
		fmt.Printf("Response: %v\n", res.GetMessage())
	}
}
