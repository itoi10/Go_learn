
サーバ go

プロキシ envoy

フロント svelte

## メモ

client作成
```npx degit sveltejs/template client```
```yarn```
```yarn dev```

protocエラー
```sh
$ sh generate_code.sh
protoc-gen-js: program not found or is not executable
```

grpc-client
```sh
$ grpc_cli ls localhost:9090 helloworld.Greeter -l
filename: proto/helloworld.proto
package: helloworld;
service Greeter {
  rpc SayHello(helloworld.HelloRequest) returns (helloworld.HelloReply) {}
  rpc SayRepeatHello(helloworld.RepeatHelloRequest) returns (stream helloworld.HelloReply) {}
}
```

起動

go api
```sh
$ go run server.go
```

svelte client
```sh
$ yarn dev
```

envoy proxy
```sh
$ docker built -t envoy .
$ docker run -it -p 8080:8080 envoy
```

## 参考記事

### gRPC Go Quick start
https://grpc.io/docs/languages/go/quickstart/

### gRPC Web
https://github.com/grpc/grpc-web

### 「 gRPC Web 」で gRPC 実践！ Go と gRPC で WebAPI を作ってみよう！！
https://www.youtube.com/watch?v=hlyNZoaXvqU

### Generate js file from proto file with protoc v21.1
https://github.com/protocolbuffers/protobuf-javascript/issues/127
