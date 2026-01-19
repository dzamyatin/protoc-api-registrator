package main

import (
	"context"
	"time"

	"io"
	"os"

	pluginpb "github.com/dzamyatin/protoc-api-registrator/proto/generated"
	"google.golang.org/grpc/encoding"
	_ "google.golang.org/grpc/encoding/proto"
	"google.golang.org/grpc/mem"
)

//go:generate protoc -I ./proto --go_out=./proto/generated/ --go_opt=paths=source_relative plugin.proto
func main() {
	req := new(pluginpb.CodeGeneratorRequest)

	sl := mem.BufferSlice{}
	data := getData()
	sl = append(sl, mem.SliceBuffer(data))

	err := encoding.GetCodecV2("proto").Unmarshal(sl, req)
	if err != nil {
		panic(err)
	}

	out := os.Stdout
	out.Write([]byte("hela"))
}

func createResponse() pluginpb.CodeGeneratorResponse {

	return pluginpb.CodeGeneratorResponse{
		Error:             nil,
		SupportedFeatures: nil,
		File: []*pluginpb.CodeGeneratorResponse_File{
			{},
		},
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
