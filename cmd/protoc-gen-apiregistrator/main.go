package main

import (
	//"io"
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
	PackageName  string
	Urls         []Path
	ClassPostfix string
}

type Path struct {
	Url        string
	Method     string
	UrlRuntime string
}

// //go:generate protoc -I ./proto --go_out=./proto/generated/ --go_opt=paths=source_relative plugin.proto
func main() {
	protogen.Options{
		//ParamFunc: flag.CommandLine.Set,
	}.Run(func(gen *protogen.Plugin) error {

		gen.SupportedFeatures = uint64(pb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

		tplRegistrator, err := template.New("base").Parse(templator.Template)
		if err != nil {
			return errors.Wrap(err, "failed to parse template")
		}

		tplRegistratorGlobal, err := template.New("base").Parse(templator.TemplateCommon)
		if err != nil {
			return errors.Wrap(err, "failed to parse template global")
		}

		urlAll := make([]Path, 0)

		protofiles := gen.Files
		var (
			packageName protogen.GoPackageName
			goImport    protogen.GoImportPath
		)

		for _, protofile := range protofiles {
			packageName = protofile.GoPackageName // atomWebsite

			var (
				fileName = protofile.Proto.GetPackage()
				urls     = make([]Path, 0)
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
							Url:        url,
							UrlRuntime: makeUrlAsRuntimePattern(url),
							Method:     "GET",
						})
					}

					if url := e.GetPost(); url != "" {
						urls = append(urls, Path{
							Url:        url,
							UrlRuntime: makeUrlAsRuntimePattern(url),
							Method:     "POST",
						})
					}

					if url := e.GetDelete(); url != "" {
						urls = append(urls, Path{
							Url:        url,
							UrlRuntime: makeUrlAsRuntimePattern(url),
							Method:     "DELETE",
						})
					}

					if url := e.GetPut(); url != "" {
						urls = append(urls, Path{
							Url:        url,
							UrlRuntime: makeUrlAsRuntimePattern(url),
							Method:     "PUT",
						})
					}
				}
			}

			if len(urls) == 0 {
				continue
			}

			err = writeToTpl(
				gen,
				tplRegistratorGlobal,
				Data{
					PackageName:  string(packageName),
					Urls:         urls,
					ClassPostfix: strings.ToUpper(string([]rune(fileName)[0])) + string([]rune(fileName[1:])),
				},
				fileName+"_url_registrator.go",
				protofile.GoImportPath,
			)
			if err != nil {
				return err
			}

			//content := strings.Builder{}
			//
			//err = tplRegistrator.Execute(&content, Data{
			//	PackageName:  string(packageName),
			//	Urls:         urls,
			//	ClassPostfix: strings.ToUpper(string([]rune(fileName)[0])) + string([]rune(fileName[1:])),
			//})
			//if err != nil {
			//	return errors.Wrap(err, "failed to execute template")
			//}
			//
			//fi := gen.NewGeneratedFile(
			//	fileName+"_url_registrator.go",
			//	protofile.GoImportPath,
			//)
			//_, err = fi.Write([]byte(content.String()))
			//if err != nil {
			//	return errors.Wrap(err, "failed to write template")
			//}

			urlAll = append(urlAll, urls...)
		}

		if len(urlAll) == 0 {
			return nil
		}

		fileName := "common_global"
		err = writeToTpl(
			gen,
			tplRegistrator,
			Data{
				PackageName: string(packageName),
				Urls:        urlAll,
			},
			fileName+"_url_registrator.go",
			goImport,
		)
		if err != nil {
			return err
		}

		return nil
	})
}

func writeToTpl(
	gen *protogen.Plugin,
	tplRegistrator *template.Template,
	data Data,
	fileName string,
	goImportPath protogen.GoImportPath,
) error {
	content := strings.Builder{}

	err := tplRegistrator.Execute(&content, data)
	if err != nil {
		return errors.Wrap(err, "failed to execute template")
	}

	fi := gen.NewGeneratedFile(
		fileName+"_url_registrator.go",
		goImportPath,
	)
	_, err = fi.Write([]byte(content.String()))
	if err != nil {
		return errors.Wrap(err, "failed to write template")
	}

	return nil
}

func makeUrlAsRuntimePattern(url string) string {
	res := make([]string, 0)
	for _, v := range strings.Split(url, "/") {
		if strings.HasPrefix(v, "{") && strings.HasSuffix(v, "}") {
			md, _ := strings.CutPrefix(v, "}")

			res = append(res, md+"=*}")
		}
	}

	return strings.Join(res, "/")
}
