//go:build tools
// +build tools

package lib

import (
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)

//go:generate protoc --go_out=. --go-grpc_out=. proto/hello.proto
