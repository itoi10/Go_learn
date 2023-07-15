package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
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
		fmt.Printf("Connection failed: %v", err)
		return
	}
	defer conn.Close()

	// gRPCクライアント生成 (この関数はprotoファイルから自動生成されている)
	client = hellopb.NewGreetingServiceClient(conn)

	for {
		fmt.Println("1: send Request(Unary)")
		fmt.Println("2: send Request(Server Streaming)")
		fmt.Println("3: send Request(Client Streaming)")
		fmt.Println("4: send Request(Bidirectional Streaming)")
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
		case "3":
			HelloClientStream()
		case "4":
			HelloBiStreams()

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
		fmt.Printf("HelloUnary failed: %v", err)
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
		fmt.Printf("HelloServerStream failed: %v", err)
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
			fmt.Printf("Response Recv failed: %v", err)
			return
		}
		fmt.Printf("Response: %v\n", res.GetMessage())
	}
}

func HelloClientStream() {
	stream, err := client.HelloClientStream(context.Background())
	if err != nil {
		fmt.Printf("HelloClientStream failed: %v", err)
		return
	}

	// 複数回サーバーに送信する
	sendCount := 5
	fmt.Printf("input your name %d times:\n", sendCount)
	for i := 0; i < sendCount; i++ {
		scanner.Scan()
		name := scanner.Text()

		if err := stream.Send(&hellopb.HelloRequest{Name: name}); err != nil {
			fmt.Printf("Request Send failed: %v", err)
			return
		}
	}

	// サーバーからのレスポンスを受け取る
	res, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Printf("Response CloseAndRecv failed: %v", err)
		return
	}
	fmt.Printf("Response: %v\n", res.GetMessage())
}

func HelloBiStreams() {
	stream, err := client.HelloBiStreams(context.Background())
	if err != nil {
		fmt.Printf("HelloBiStreams failed: %v", err)
		return
	}

	sendNum := 5
	fmt.Printf("input your name %d times:\n", sendNum)

	sendEnd := false
	recvEnd := false
	sendCount := 0
	for !(sendEnd && recvEnd) {
		// 送信処理
		if !sendEnd {
			scanner.Scan()
			name := scanner.Text()

			sendCount += 1
			if err := stream.Send(&hellopb.HelloRequest{Name: name}); err != nil {
				fmt.Println("Request Send failed:", err)
				sendEnd = true
			}

			if sendCount >= sendNum {
				// 送信回数が上限に達したら送信終了
				if err := stream.CloseSend(); err != nil {
					fmt.Println("Request CloseSend failed:", err)
				}
				sendEnd = true
			}
		}

		// 受信処理
		if !recvEnd {
			if res, err := stream.Recv(); err != nil {
				if !errors.Is(err, io.EOF) {
					fmt.Println("Response Recv failed:", err)
				}
				recvEnd = true
			} else {
				fmt.Printf("Response: %v\n", res.GetMessage())
			}
		}
	}

}
