.PHONY: build

build:
	go build -o ./bin/protoc-gen-logger ./cmd/protologger/main.go
	go build -o ./bin/protoc-gen-api-registrator ./main.go