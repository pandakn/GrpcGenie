package generator

import (
	"os"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/pandakn/GrpcGenie/internal/templates"
)

type Method struct {
	Name       string
	InputType  string
	OutputType string
}

type TemplateData struct {
	PackageName     string
	GrpcPackageName string
	ServiceName     string
	GoPackagePath   string
	Methods         []Method
}

func GenerateHandler(protoPath, outputPath string, data TemplateData) error {
	tmpl, err := templates.GetHandlerTemplate()
	if err != nil {
		return err
	}
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()
	return tmpl.Execute(file, data)
}

func ParseProtoFile(protoPath string) (*desc.FileDescriptor, error) {
	parser := protoparse.Parser{}
	descriptors, err := parser.ParseFiles(protoPath)
	if err != nil {
		return nil, err
	}
	return descriptors[0], nil
}
