package service

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/xryar/golang-grpc-ecommerce/internal/entity"
	jwtentity "github.com/xryar/golang-grpc-ecommerce/internal/entity/jwt"
	"github.com/xryar/golang-grpc-ecommerce/internal/repository"
	"github.com/xryar/golang-grpc-ecommerce/internal/utils"
	"github.com/xryar/golang-grpc-ecommerce/pb/product"
)

type IProductService interface {
	CreateProduct(ctx context.Context, request *product.CreateProductRequest) (*product.CreateProductResponse, error)
	DetailProduct(ctx context.Context, request *product.DetailProductRequest) (*product.DetailProductResponse, error)
	EditProduct(ctx context.Context, request *product.EditProductRequest) (*product.EditProductResponse, error)
}

type productService struct {
	productRepository repository.IProductRepository
}

func (ps *productService) CreateProduct(ctx context.Context, request *product.CreateProductRequest) (*product.CreateProductResponse, error) {
	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if claims.Role != entity.UserRoleAdmin {
		return nil, utils.UnauthenticatedResponse()
	}

	imagePath := filepath.Join("storage", "product", request.ImageFileName)
	_, err = os.Stat(imagePath)
	if err != nil {
		if os.IsNotExist(err) {
			return &product.CreateProductResponse{
				Base: utils.BadRequestResponse("File not found"),
			}, nil
		}

		return nil, err
	}

	productEntity := entity.Product{
		Id:            uuid.NewString(),
		Name:          request.Name,
		Description:   request.Description,
		Price:         request.Price,
		ImageFileName: request.ImageFileName,
		CreatedAt:     time.Now(),
		CreatedBy:     claims.Fullname,
	}
	err = ps.productRepository.CreateNewProduct(ctx, &productEntity)
	if err != nil {
		return nil, err
	}

	return &product.CreateProductResponse{
		Base: utils.SuccessResponse("Product is created"),
		Id:   productEntity.Id,
	}, nil
}

func (ps *productService) DetailProduct(ctx context.Context, request *product.DetailProductRequest) (*product.DetailProductResponse, error) {
	productEntity, err := ps.productRepository.GetProductById(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	if productEntity == nil {
		return &product.DetailProductResponse{
			Base: utils.NotFoundResponse("Product not found"),
		}, nil
	}

	return &product.DetailProductResponse{
		Base:        utils.SuccessResponse("Success Get detail product"),
		Id:          productEntity.Id,
		Name:        productEntity.Name,
		Description: productEntity.Description,
		Price:       productEntity.Price,
		ImageUrl:    fmt.Sprintf("%s/product/%s", os.Getenv("STORAGE_SERVICE_URL"), productEntity.ImageFileName),
	}, nil
}

func (ps *productService) EditProduct(ctx context.Context, request *product.EditProductRequest) (*product.EditProductResponse, error) {
	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if claims.Role != entity.UserRoleAdmin {
		return nil, utils.UnauthenticatedResponse()
	}

	productEntity, err := ps.productRepository.GetProductById(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	if productEntity == nil {
		return &product.EditProductResponse{
			Base: utils.NotFoundResponse("Product not found"),
		}, nil
	}

	if productEntity.ImageFileName != request.ImageFileName {
		newImagePath := filepath.Join("storage", "product", request.ImageFileName)
		_, err := os.Stat(newImagePath)
		if err != nil {
			if os.IsNotExist(err) {
				return &product.EditProductResponse{
					Base: utils.BadRequestResponse("Image not found"),
				}, nil
			}

			return nil, err
		}

		oldImagePath := filepath.Join("storage", "product", productEntity.ImageFileName)
		err = os.Remove(oldImagePath)
		if err != nil {
			return nil, err
		}
	}

	newProduct := entity.Product{
		Id:            request.Id,
		Name:          request.Name,
		Description:   request.Description,
		Price:         request.Price,
		ImageFileName: request.ImageFileName,
		UpdatedAt:     time.Now(),
		UpdatedBy:     &claims.Fullname,
	}

	err = ps.productRepository.UpdateProduct(ctx, &newProduct)
	if err != nil {
		return nil, err
	}

	return &product.EditProductResponse{
		Base: utils.SuccessResponse("Edit Product Success"),
		Id:   request.Id,
	}, nil
}

func NewProductService(productRepository repository.IProductRepository) IProductService {
	return &productService{
		productRepository: productRepository,
	}
}
