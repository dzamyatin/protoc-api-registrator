.PHONY: build,proto

build:
	go build -o ./bin/protoc-gen-logger ./cmd/protologger/main.go
	go build -o ./bin/protoc-gen-api-registrator ./main.go

#protoc:
#	cd proto && protoc --go_out=./generated/ --go_opt=paths=source_relative ./googleapis/google/api/http.proto && protoc --go_out=./generated/ --go_opt=paths=source_relative ./googleapis/google/api/annotations.proto
protoc:
	protoc --proto_path=./googleapis/ --go_out=./generated/ --go_opt=paths=source_relative ./googleapis/google/api/annotations.proto
	protoc --proto_path=./googleapis/ --go_out=./generated/ --go_opt=paths=source_relative ./googleapis/google/api/http.proto


#proto: https://github.com/googleapis/googleapis/tree/master/google/api

#//go:generate protoc -I. -I./google --plugin=protoc-gen-logger=/home/dzamyatin/GolandProjects/protoc-api-registrator/bin/protoc-gen-logger --logger_out=../internal/grpc/generated/ --openapiv2_out . --go_out=../internal/grpc/generated/ --go_opt=paths=source_relative --go-grpc_out=../internal/grpc/generated/ --go-grpc_opt=paths=source_relative --grpc-gateway_out ../internal/grpc/generated/ --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true auth.proto shop.proto