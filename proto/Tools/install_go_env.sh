#!/bin/bash
set -eu

if ! command -v go; then
  echo "Installing goenv"
  brew install goenv
  goenv install 1.21.5
  goenv global 1.21.5
  goenv rehash
  go version
fi

if ! command -v protoc-gen-go; then
  echo "Installing protoc-gen-go"
  go install google.golang.org/protobuf/cmd/protoc-gen-go
fi

if ! command -v protoc-gen-go-grpc; then
  echo "Installing protoc-gen-go-grpc"
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
fi

if ! command -v buf; then
  echo "Installing buf"
  go install github.com/bufbuild/buf/cmd/buf@latest
fi

if ! command -v grpcurl; then
  echo "Installing grpcurl"
  go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
fi

if ! command -v protoc-gen-connect-go; then
  echo "Installing protoc-gen-connect-go"
  go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
fi
