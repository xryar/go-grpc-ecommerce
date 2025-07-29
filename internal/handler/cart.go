package handler

import (
	"context"

	"github.com/xryar/golang-grpc-ecommerce/internal/service"
	"github.com/xryar/golang-grpc-ecommerce/internal/utils"
	"github.com/xryar/golang-grpc-ecommerce/pb/cart"
)

type cartHandler struct {
	cart.UnimplementedCartServiceServer

	cartService service.ICartService
}

func (ch *cartHandler) AddProductToCart(ctx context.Context, request *cart.AddProductToCartRequest) (*cart.AddProductToCartResponse, error) {
	validationErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}
	if validationErrors != nil {
		return &cart.AddProductToCartResponse{
			Base: utils.ValidationErrorResponse(validationErrors),
		}, nil
	}

	res, err := ch.cartService.AddProductToCart(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewCartHandler(cartService service.ICartService) *cartHandler {
	return &cartHandler{
		cartService: cartService,
	}
}
