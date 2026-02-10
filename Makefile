.PHONY: build,proto

build:
	go build -o ./bin/protoc-gen-logger ./cmd/protologger/main.go
	go build -o ./bin/protoc-gen-api-registrator ./main.go

#protoc:
#	cd proto && protoc --go_out=./generated/ --go_opt=paths=source_relative ./googleapis/google/api/http.proto && protoc --go_out=./generated/ --go_opt=paths=source_relative ./googleapis/google/api/annotations.proto
protoc:
	protoc --proto_path=./googleapis/ --go_out=./generated/ --go_opt=paths=source_relative ./googleapis/google/api/annotations.proto
	protoc --proto_path=./googleapis/ --go_out=./generated/ --go_opt=paths=source_relative ./googleapis/google/api/http.proto