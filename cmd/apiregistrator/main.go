package main

import (
	//"io"
	"strings"
	"text/template"

	"github.com/dzamyatin/protoc-api-registrator/internal/templator"
	pluginpb "github.com/dzamyatin/protoc-api-registrator/proto/generated"

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
	Urls        []Path
}

type Path struct {
	Url    string
	Method string
}

//go:generate protoc -I ./proto --go_out=./proto/generated/ --go_opt=paths=source_relative plugin.proto
func main() {
	protogen.Options{
		//ParamFunc: flag.CommandLine.Set,
	}.Run(func(gen *protogen.Plugin) error {
		res := gen.Response()

		tplRegistrator, err := template.New("base").Parse(templator.Template)
		if err != nil {
			return errors.Wrap(err, "failed to parse template")
		}

		protofiles := gen.Files
		for _, protofile := range protofiles {
			var (
				packageName = protofile.GoPackageName // atomWebsite
				fileName    = protofile.Proto.GetPackage()
				urls        = make([]Path, 0)
			)

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

					if url := e.GetGet(); url != "" {
						urls = append(urls, Path{
							Url:    url,
							Method: "GET",
						})
					}

					if url := e.GetPost(); url != "" {
						urls = append(urls, Path{
							Url:    url,
							Method: "POST",
						})
					}

					if url := e.GetDelete(); url != "" {
						urls = append(urls, Path{
							Url:    url,
							Method: "DELETE",
						})
					}

					if url := e.GetPut(); url != "" {
						urls = append(urls, Path{
							Url:    url,
							Method: "PUT",
						})
					}
				}
			}

			content := strings.Builder{}

			err = tplRegistrator.Execute(&content, Data{
				PackageName: string(packageName),
				Urls:        urls,
			})
			if err != nil {
				return errors.Wrap(err, "failed to execute template")
			}

			contentString := content.String()

			n := fileName + "_url_registrator.go"
			f := pb.CodeGeneratorResponse_File{
				Name:    &n,
				Content: &contentString,
			}
			res.File = append(res.File, &f)
		}

		i := uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		res.SupportedFeatures = &i

		return nil
	})
}
