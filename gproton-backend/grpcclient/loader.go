package grpcclient

import (
	"log"
	"os"

	"github.com/jhump/protoreflect/desc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func LoadProtoSet(filePath string) *ProtoSet {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fds := &descriptorpb.FileDescriptorSet{}
	if err := proto.Unmarshal(data, fds); err != nil {
		log.Fatal(err)
	}
	files, err := desc.CreateFileDescriptorsFromSet(fds)
	if err != nil {
		log.Fatal(err)
	}

	ps := &ProtoSet{}
	for _, fd := range files {
		for _, svc := range fd.GetServices() {
			s := Service{Name: svc.GetName()}
			for _, m := range svc.GetMethods() {
				s.Methods = append(s.Methods, Method{
					Name:       m.GetName(),
					InputType:  m.GetInputType().GetFullyQualifiedName(),
					OutputType: m.GetOutputType().GetFullyQualifiedName(),
					Desc:       m, // ✅ اینجا MethodDescriptor واقعی است
				})
			}
			ps.Services = append(ps.Services, s)
		}
	}

	return ps
}
