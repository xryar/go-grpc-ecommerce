package handler

import (
	"context"
	"fmt"

	"github.com/xryar/golang-grpc-ecommerce/pb/service"
)

type serviceHandler struct {
	service.UnimplementedHelloWorldServiceServer
}

func (sh *serviceHandler) HelloWorld(ctx context.Context, request *service.HelloWorldRequest) (*service.HelloWorldResponse, error) {
	return &service.HelloWorldResponse{
		Message: fmt.Sprintf("Hello %s", request.Name),
	}, nil
}

func NewServiceHandler() *serviceHandler {
	return &serviceHandler{}
}
