module github.com/dzamyatin/protoc-api-registrator

go 1.26rc2

tool google.golang.org/protobuf/cmd/protoc-gen-go

require (
	google.golang.org/grpc v1.78.0
	google.golang.org/protobuf v1.36.11
)

require golang.org/x/sys v0.40.0 // indirect
