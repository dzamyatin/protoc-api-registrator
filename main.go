package main

import (
	_ "github.com/dzamyatin/protoc-api-registrator/proto/generated/google/api"          //authomatically init to make oriti iotuion available to parse
	annotation "github.com/dzamyatin/protoc-api-registrator/proto/generated/google/api" //authomatically init to make oriti iotuion available to parse
	_ "google.golang.org/grpc/encoding/proto"
	"google.golang.org/protobuf/compiler/protogen"
	_ "google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

//go:generate protoc -I ./proto --go_out=./proto/generated/ --go_opt=paths=source_relative plugin.proto
func main() {
	protogen.Options{
		//ParamFunc: flag.CommandLine.Set,
	}.Run(func(gen *protogen.Plugin) error {

		//annotations

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
					//_ = e
					//prm := option.ProtoReflect()
					//_ = prm
					//_ = option
					//s := option.String()
					//option.extensionFields

					//h := annotations.Http{}
					//proto.GetExtension(&h, option.)
					//_ = s
					//fmt.Println(option.String())
					//for _, unint := range option.String() {
					//	_ = unint
					//}
					//
					//req := new(annotations.Http)
					//
					//sl := mem.BufferSlice{}
					//data := option.String()
					//sl = append(sl, mem.SliceBuffer(data))
					//
					//err := encoding.GetCodecV2("proto").Unmarshal(sl, req)
					//if err != nil {
					//	continue
					//}
				}
			}
		}

		//sl := mem.BufferSlice{}
		//data := ""
		//sl = append(sl, mem.SliceBuffer(data))
		//
		//req := new(annotations.Http)
		//err := encoding.GetCodecV2("proto").Unmarshal(sl, req)
		//if err != nil {
		//	panic(err)
		//}

		return nil
	})

	//_, err = os.Stdout.Write(buf.Materialize())
	//if err != nil {
	//	panic(err)
	//}
}

//
//req := new(pluginpb.CodeGeneratorRequest)
//
//sl := mem.BufferSlice{}
//data := getData()
//sl = append(sl, mem.SliceBuffer(data))
//
//err := encoding.GetCodecV2("proto").Unmarshal(sl, req)
//if err != nil {
//panic(err)
//}
//
//buf, err := encoding.GetCodecV2("proto").Marshal(createResponse())
//if err != nil {
//panic(err)
//}

//
//func createResponse() *pluginpb.CodeGeneratorResponse {
//	return &pluginpb.CodeGeneratorResponse{
//		Error:             nil,
//		SupportedFeatures: nil,
//		File:              []*pluginpb.CodeGeneratorResponse_File{},
//	}
//}
//
//func getData() []byte {
//	var read = make(chan []byte, 1)
//
//	go func() {
//		defer close(read)
//		data, err := io.ReadAll(os.Stdin)
//		if err != nil {
//			panic(err)
//		}
//		read <- data
//	}()
//
//	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
//	defer cancel()
//
//	var data []byte
//	select {
//	case data = <-read:
//	case <-ctx.Done():
//		panic("cant read from stdin in time")
//	}
//
//	return data
//}
