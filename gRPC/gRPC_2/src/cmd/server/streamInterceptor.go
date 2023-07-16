package main

import (
	"errors"
	"io"
	"log"

	"google.golang.org/grpc"
)

// StreamServer用のInterceptor
func myStreamServerInterceptor1(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// ストリームがopenされたときに行われる前処理
	log.Println("[pre] myStreamServerInterceptor1: ", info.FullMethod)

	// ストリーム処理
	err := handler(srv, &myServerStreamWrapper1{ss})

	// ストリームがcloseされたときに行われる後処理
	log.Println("[post] myStreamServerInterceptor1: err:", err)

	return err
}

type myServerStreamWrapper1 struct {
	grpc.ServerStream
}

// リクエスト受信時の前処理
func (s *myServerStreamWrapper1) RecvMsg(m interface{}) error {
	// リクエスト受信
	err := s.ServerStream.RecvMsg(m)
	// ハンドラで処理する前の前処理
	if !errors.Is(err, io.EOF) {
		log.Printf("[pre message] myStreamServerInterceptor1: %v", m)
	}
	return err
}

// レスポンス送信時の後処理
func (s *myServerStreamWrapper1) SendMsg(m interface{}) error {
	// ハンドラで作成したレスポンスを返信する直前の後処理
	log.Printf("[post message] myStreamServerInterceptor1: %v", m)
	return s.ServerStream.SendMsg(m)
}
