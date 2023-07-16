package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

// Unary用のInterceptor
func myUnaryServerInterceptor1(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 前処理
	log.Println("[pre] myUnaryServerInterceptor1: ", info.FullMethod)

	// handlerの呼び出し
	res, err := handler(ctx, req)

	// 後処理
	log.Println("[post] myUnaryServerInterceptor1: ", res)

	return res, err
}

// myUnaryServerInterceptor1と同じ処理を行うInterceptor
func myUnaryServerInterceptor2(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("[pre] myUnaryServerInterceptor2: ", info.FullMethod)
	res, err := handler(ctx, req)
	log.Println("[post] myUnaryServerInterceptor2: ", res)
	return res, err
}
