module github.com/hiddenmarten/go-notes/grpc/basic/server

go 1.22.6

replace github.com/hiddenmarten/go-notes/grpc/basic/proto => ../proto

require (
	github.com/hiddenmarten/go-notes/grpc/basic/proto v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.65.0
)

require (
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240604185151-ef581f913117 // indirect
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.5.1 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)
