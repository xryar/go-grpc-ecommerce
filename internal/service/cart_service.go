package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/xryar/golang-grpc-ecommerce/internal/entity"
	jwtentity "github.com/xryar/golang-grpc-ecommerce/internal/entity/jwt"
	"github.com/xryar/golang-grpc-ecommerce/internal/repository"
	"github.com/xryar/golang-grpc-ecommerce/internal/utils"
	"github.com/xryar/golang-grpc-ecommerce/pb/cart"
)

type ICartService interface {
	AddProductToCart(ctx context.Context, request *cart.AddProductToCartRequest) (*cart.AddProductToCartResponse, error)
	ListCart(ctx context.Context, request *cart.ListCartRequest) (*cart.ListCartResponse, error)
}

type cartService struct {
	productRepository repository.IProductRepository
	cartRepository    repository.ICartRepository
}

func (cs *cartService) AddProductToCart(ctx context.Context, request *cart.AddProductToCartRequest) (*cart.AddProductToCartResponse, error) {
	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	// cek apakah product id ada di db
	productEntity, err := cs.productRepository.GetProductById(ctx, request.ProductId)
	if err != nil {
		return nil, err
	}
	if productEntity == nil {
		return &cart.AddProductToCartResponse{
			Base: utils.NotFoundResponse("Product not found"),
		}, nil
	}

	// cek apakah product sudah ada di cart user
	cartEntity, err := cs.cartRepository.GetCartByProductAndUserId(ctx, request.ProductId, claims.Subject)
	if err != nil {
		return nil, err
	}

	if cartEntity != nil {
		now := time.Now()
		cartEntity.Quantity += 1
		cartEntity.UpdateAt = &now
		cartEntity.UpdatedBy = &claims.Subject

		err = cs.cartRepository.UpdateCart(ctx, cartEntity)
		if err != nil {
			return nil, err
		}

		return &cart.AddProductToCartResponse{
			Base: utils.SuccessResponse("Add Product to Cart Success"),
			Id:   cartEntity.Id,
		}, nil
	}

	newCartEntity := entity.UserCart{
		Id:        uuid.NewString(),
		UserId:    claims.Subject,
		ProductId: request.ProductId,
		Quantity:  1,
		CreatedAt: time.Now(),
		CreatedBy: claims.Fullname,
	}
	err = cs.cartRepository.CreateNewCart(ctx, &newCartEntity)
	if err != nil {
		return nil, err
	}

	return &cart.AddProductToCartResponse{
		Base: utils.SuccessResponse("Add Product to Cart Success"),
		Id:   newCartEntity.Id,
	}, nil

}

func (cs *cartService) ListCart(ctx context.Context, request *cart.ListCartRequest) (*cart.ListCartResponse, error) {
	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	carts, err := cs.cartRepository.GetListCart(ctx, claims.Subject)
	if err != nil {
		return nil, err
	}

	var items []*cart.ListCartResponseItem = make([]*cart.ListCartResponseItem, 0)
	for _, cartEntity := range carts {
		item := cart.ListCartResponseItem{
			ProductId:       cartEntity.Product.Id,
			ProductName:     cartEntity.Product.Name,
			ProductImageUrl: fmt.Sprintf("%s/product/%s", os.Getenv("STORAGE_SERVICE_URL"), cartEntity.Product.ImageFileName),
			ProductPrice:    cartEntity.Product.Price,
			Quantity:        int64(cartEntity.Quantity),
		}

		items = append(items, &item)
	}

	return &cart.ListCartResponse{
		Base:  utils.SuccessResponse("Get List Cart Success"),
		Items: items,
	}, nil
}

func NewCartService(productRespository repository.IProductRepository, cartRepository repository.ICartRepository) ICartService {
	return &cartService{
		productRepository: productRespository,
		cartRepository:    cartRepository,
	}
}
