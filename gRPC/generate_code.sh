#!/bin/sh

protoc proto/helloworld.proto \
    --plugin=protoc-gen-js \
    --js_out=import_style=commonjs:client/src/helloworld \
    --grpc-web_out=import_style=commonjs+dts,mode=grpcwebtext:client/src/helloworld
