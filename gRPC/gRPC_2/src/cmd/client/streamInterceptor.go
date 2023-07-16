package main

import (
	"context"
	"errors"
	"io"
	"log"

	"google.golang.org/grpc"
)

func myStreamClientInterceptor1(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	// ストリームがopenされたときに行われる前処理
	log.Println("[pre] myStreamClientInterceptor1: ", method)

	stream, err := streamer(ctx, desc, cc, method, opts...)
	return &myClientStreamWrapper1{stream}, err
}

type myClientStreamWrapper1 struct {
	grpc.ClientStream
}

// リクエスト送信時の前処理
func (s *myClientStreamWrapper1) SendMsg(m interface{}) error {
	// リクエストを送信する前の前処理
	log.Println("[pre message] myStreamClientInterceptor1: ", m)
	// リクエスト送信
	return s.ClientStream.SendMsg(m)
}

// レスポンス受信時の後処理
func (s *myClientStreamWrapper1) RecvMsg(m interface{}) error {
	// レスポンス受信
	err := s.ClientStream.RecvMsg(m)
	// レスポンス受信後の後処理
	if !errors.Is(err, io.EOF) {
		log.Println("[post message] myStreamClientInterceptor1: ", m)
	}
	return err
}

// ストリームがcloseされたときに行われる後処理
func (s *myClientStreamWrapper1) CloseSend() error {
	err := s.ClientStream.CloseSend()
	log.Println("[post] myStreamClientInterceptor1: err:", err)
	return err
}
