package templates

import (
	_ "embed"
	"text/template"
)

//go:embed files/grpc_handler_template.go.tmpl
var handlerTemplate string

func GetHandlerTemplate() (*template.Template, error) {
	return template.New("handler").Parse(handlerTemplate)
}
