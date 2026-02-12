

use to log proto:
> protoc -I. -I./google --plugin=protoc-gen-logger=/home/dzamyatin/GolandProjects/protoc-api-registrator/bin/protoc-gen-logger --logger_out=../internal/grpc/generated/

use for registrator
> protoc -I. -I./google --plugin=protoc-gen-urlregistrator=/home/dzamyatin/GolandProjects/protoc-api-registrator/bin/protoc-gen-api-registrator --urlregistrator_out=../internal/grpc/generated/