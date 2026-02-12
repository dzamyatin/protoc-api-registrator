package main

import (
	"context"
	"strconv"
	"time"

	"io"
	"os"

	"google.golang.org/grpc/encoding"
	_ "google.golang.org/grpc/encoding/proto"
	pb "google.golang.org/protobuf/types/pluginpb"
)

var protologgerFile = "/home/dzamyatin/GolandProjects/protoc-api-registrator/var/protologger/"

func main() {
	data := getData()

	f, err := os.OpenFile(
		protologgerFile+
			"file"+
			strconv.FormatInt(time.Now().UnixNano(), 10)+
			".txt",
		os.O_WRONLY|os.O_CREATE,
		0666,
	)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Write(data)

	res := createResponse()
	mb, err := encoding.GetCodecV2("proto").Marshal(&res)
	if err != nil {
		panic(err)
	}

	out := os.Stdout
	out.Write(mb.Materialize())
}

func createResponse() pb.CodeGeneratorResponse {
	i := uint64(1)
	return pb.CodeGeneratorResponse{
		Error:             nil,
		SupportedFeatures: &i,
		File:              []*pb.CodeGeneratorResponse_File{},
	}
}

func getData() []byte {
	var read = make(chan []byte, 1)

	go func() {
		defer close(read)
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		read <- data
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var data []byte
	select {
	case data = <-read:
	case <-ctx.Done():
		panic("cant read from stdin in time")
	}

	return data
}
