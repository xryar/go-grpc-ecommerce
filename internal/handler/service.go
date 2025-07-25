package handler

import (
	"context"
	"fmt"

	"github.com/xryar/golang-grpc-ecommerce/internal/utils"
	"github.com/xryar/golang-grpc-ecommerce/pb/service"
)

type serviceHandler struct {
	service.UnimplementedHelloWorldServiceServer
}

func (sh *serviceHandler) HelloWorld(ctx context.Context, request *service.HelloWorldRequest) (*service.HelloWorldResponse, error) {
	validationErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}
	if validationErrors != nil {
		return &service.HelloWorldResponse{
			Base: utils.ValidationErrorResponse(validationErrors),
		}, nil
	}

	return &service.HelloWorldResponse{
		Message: fmt.Sprintf("Hello %s", request.Name),
		// Base:    utils.SuccessResponse(),
	}, nil
}

func NewServiceHandler() *serviceHandler {
	return &serviceHandler{}
}
