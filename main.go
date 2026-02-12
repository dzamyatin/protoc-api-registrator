package main

import (
	"io"
	"strings"
	"text/template"

	"github.com/dzamyatin/protoc-api-registrator/internal/templator"
	//pluginpb "github.com/dzamyatin/protoc-api-registrator/proto/generated"

	_ "github.com/dzamyatin/protoc-api-registrator/proto/generated/google/api"          //authomatically init to make oriti iotuion available to parse
	annotation "github.com/dzamyatin/protoc-api-registrator/proto/generated/google/api" //authomatically init to make oriti iotuion available to parse
	_ "google.golang.org/grpc/encoding/proto"
	"google.golang.org/protobuf/compiler/protogen"
	_ "google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	pb "google.golang.org/protobuf/types/pluginpb"

	"github.com/pkg/errors"
)

type Data struct {
	PackageName string
	Urls        []string
}

//go:generate protoc -I ./proto --go_out=./proto/generated/ --go_opt=paths=source_relative plugin.proto
func main() {
	protogen.Options{
		//ParamFunc: flag.CommandLine.Set,
	}.Run(func(gen *protogen.Plugin) error {
		protofiles := gen.Files
		for _, protofile := range protofiles {
			for _, service := range protofile.Services {
				for _, method := range service.Methods {
					_ = method
					option, ok := method.Desc.Options().(*descriptorpb.MethodOptions)
					if !ok {
						continue
					}

					e, ok := proto.GetExtension(option, annotation.E_Http).(*annotation.HttpRule)
					if !ok {

						continue
					}
					p := e.GetPattern()
					_ = p

					//
					var (
						packageName = protofile.GoPackageName // atomWebsite
						fileName    = protofile.Proto.GetPackage()
					)

					t, err := template.New("base").Parse(templator.Template)
					if err != nil {
						return errors.Wrap(err, "failed to parse template")
					}

					content := strings.Builder{}

					t.Execute(&content, Data{
						PackageName: string(packageName),
						Urls: []string{
							"",
						},
					})

					res := gen.Response()
					n := fileName + ".go"
					f := pb.CodeGeneratorResponse_File{
						Name: &n,
						//InsertionPoint:    nil,
						Content: nil,
						//GeneratedCodeInfo: nil,
					}
					res.File = append(res.File, &f)
				}
			}
		}

		return nil
	})
}
