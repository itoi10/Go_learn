# gRPC


## memo

- RPC - Remote Procedure Call
  - リモートの関数呼び出し. gRPCはRPCを実現する方式の1つ


- protoからgoのコードを自動生成

ディレクトリ構成
```
src
├── api
│   └── hello.proto // protoファイル
├── go.mod
├── go.sum
└── pkg
    └── grpc // protoから生成されたgoのコードが入る
        ├── hello.pb.go       // Request,Responseの型やGetterなどを定義したコード
        └── hello_grpc.pb.go  // Serviceのサーバサイド,クライアントサイドのInterfaceなどのコード
```

前提 protobuf, protoc-gen-go-grpcインストール済み
```
$ brew install protobuf
$ go mod init mygrpc
$ go get -u google.golang.org/grpc
$ go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

コマンド
```
$ cd api
$ protoc --go_out=../pkg/grpc --go_opt=paths=source_relative \
        --go-grpc_out=../pkg/grpc --go-grpc_opt=paths=source_relative \
        hello.proto
```
