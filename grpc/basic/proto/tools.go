//go:build tools
// +build tools

package tools

import (
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)

//go:generate protoc --go_out=. --go-grpc_out=. hello.proto
