module github.com/dzamyatin/protoc-api-registrator

go 1.25

tool google.golang.org/protobuf/cmd/protoc-gen-go

require (
	github.com/pkg/errors v0.9.1
	google.golang.org/grpc v1.78.0
	google.golang.org/protobuf v1.36.11
)

require (
	golang.org/x/sys v0.40.0 // indirect
	golang.org/x/text v0.32.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251222181119-0a764e51fe1b // indirect
)
