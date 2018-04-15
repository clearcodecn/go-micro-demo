#!/bin/bash
out=./pb
protoc
  # -I. \
  # -I$GOPATH/src \
  # -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=$out \
  --micro_out=$out \
  *.proto
