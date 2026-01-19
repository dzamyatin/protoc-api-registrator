.PHONY: build

build:
	go build -o ./bin/proto-generator-plugin ./cmd/protogenerate/main.go
	go build -o ./bin/proto-logger-plugin ./cmd/protologger/main.go
	go build -o ./bin/proto-api-registrator-plugin ./main.go