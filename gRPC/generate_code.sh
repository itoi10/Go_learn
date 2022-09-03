#!/bin/sh

protoc proto/helloworld/prot \
        -js_out=import_style=commonjs:client/src/helloworld \
        --grpc-web_out=import_style=commonjs+dts,mode=grpcwebtext:client/src/helloworld
