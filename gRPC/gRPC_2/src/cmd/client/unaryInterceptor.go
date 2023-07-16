package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

// client用のInterceptor
func myUnaryClientInterceptor1(ctx context.Context, method string, req, res interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// リクエスト送信前の前処理
	log.Println("[pre] myUnaryClientInterceptor1: ", method, req)

	// リクエスト送信
	err := invoker(ctx, method, req, res, cc, opts...)

	// リクエスト送信後の後処理
	log.Println("[post] myUnaryClientInterceptor1: ", res)

	return err
}
