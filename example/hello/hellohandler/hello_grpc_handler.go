package hellohandler

import (
	"context"

	"github.com/pandakn/GrpcGenie/example/hello/hellopd"
)

type (
	GreeterHandler struct {
		hellopd.UnimplementedGreeterServer
	}
)

func NewGreeterHandler() GreeterHandler {
	return GreeterHandler{}
}


func (g *GreeterHandler) SayHello(ctx context.Context, req *hellopd.HelloRequest) (*hellopd.HelloReply, error) {
	return nil, nil
}

func (g *GreeterHandler) Seeya(ctx context.Context, req *hellopd.GoodbyeRequest) (*hellopd.GoodbyeReply, error) {
	return nil, nil
}

