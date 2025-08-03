package service

import (
	"context"

	"github.com/xryar/golang-grpc-ecommerce/pb/order"
)

type IOrderService interface {
	CreateOrder(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error)
}

type orderService struct {
}

func (os *orderService) CreateOrder(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {

}

func NewOrderService() IOrderService {
	return &orderService{}
}
