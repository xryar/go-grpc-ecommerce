package service

import (
	"context"

	"github.com/xryar/golang-grpc-ecommerce/pb/cart"
)

type ICartService interface {
	AddProductToCart(ctx context.Context, request *cart.AddProductToCartRequest) (*cart.AddProductToCartResponse, error)
}
