package main

import (
	"fmt"

	pluginpb "github.com/dzamyatin/protoc-api-registrator/proto/generated"
	"google.golang.org/grpc/encoding"
	_ "google.golang.org/grpc/encoding/proto"
	"google.golang.org/protobuf/types/descriptorpb"

	"os"
)

const fileToWrite = "/home/dzamyatin/GolandProjects/protoc-api-registrator/test.txt"

func main() {
	req := getReq()

	res, err := encoding.GetCodecV2("proto").Marshal(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res.Materialize()))
	fmt.Printf("%b", res.Materialize())

	f, err := os.OpenFile(fileToWrite, os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Write(res.Materialize())
}

func getReq() any {
	parameter := "parameter"
	hello := "hello"
	p := "package"

	return &pluginpb.CodeGeneratorRequest{
		FileToGenerate: []string{"fileToGenerate"},
		Parameter:      &parameter,
		ProtoFile: []*descriptorpb.FileDescriptorProto{
			{
				Name:             &hello,
				Package:          &p,
				Dependency:       []string{},
				PublicDependency: []int32{},
				WeakDependency:   []int32{},
				OptionDependency: []string{"OptionDependency"},
				MessageType:      []*descriptorpb.DescriptorProto{},
				EnumType:         []*descriptorpb.EnumDescriptorProto{},
				Service:          []*descriptorpb.ServiceDescriptorProto{},
				Extension:        []*descriptorpb.FieldDescriptorProto{},
				//Options:          new("Options"),
				//SourceCodeInfo:   new("SourceCodeInfo"),
				//Syntax:           new("Syntax"),
				//Edition:          new("Edition"),
			},
		},
		CompilerVersion: nil,
	}
}
