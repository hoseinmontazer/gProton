package grpcclient

import "github.com/jhump/protoreflect/desc"

type Method struct {
	Name       string
	InputType  string
	OutputType string
	Desc       *desc.MethodDescriptor
}

type Service struct {
	Name    string
	Methods []Method
}

type ProtoSet struct {
	Services []Service
	Files    []*desc.FileDescriptor
}
