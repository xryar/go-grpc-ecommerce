package handler

import (
	"context"

	"github.com/xryar/golang-grpc-ecommerce/internal/service"
	"github.com/xryar/golang-grpc-ecommerce/internal/utils"
	"github.com/xryar/golang-grpc-ecommerce/pb/product"
)

type productHandler struct {
	product.UnimplementedProductServiceServer

	productService service.IProductService
}

func (ph *productHandler) CreateProduct(ctx context.Context, request *product.CreateProductRequest) (*product.CreateProductResponse, error) {
	validationErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}
	if validationErrors != nil {
		return &product.CreateProductResponse{
			Base: utils.ValidationErrorResponse(validationErrors),
		}, nil
	}

	res, err := ph.productService.CreateProduct(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewProductHandler(productService service.IProductService) *productHandler {
	return &productHandler{
		productService: productService,
	}
}
