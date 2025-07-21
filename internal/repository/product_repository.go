package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/xryar/golang-grpc-ecommerce/internal/entity"
	"github.com/xryar/golang-grpc-ecommerce/pb/common"
)

type IProductRepository interface {
	CreateNewProduct(ctx context.Context, product *entity.Product) error
	GetProductById(ctx context.Context, id string) (*entity.Product, error)
	UpdateProduct(ctx context.Context, product *entity.Product) error
	DeleteProduct(ctx context.Context, id string, deletedAt time.Time, deleteBy string) error
	GetProductsPagination(ctx context.Context, pagination *common.PaginationRequest) ([]*entity.Product, *common.PaginationResponse, error)
	GetProductsPaginationAdmin(ctx context.Context, pagination *common.PaginationRequest) ([]*entity.Product, *common.PaginationResponse, error)
	GetProductHighlight(ctx context.Context) ([]*entity.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func (repo *productRepository) CreateNewProduct(ctx context.Context, product *entity.Product) error {
	_, err := repo.db.ExecContext(
		ctx,
		"INSERT INTO product (id, name, description, price, image_file_name, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, is_deleted) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)",
		product.Id,
		product.Name,
		product.Description,
		product.Price,
		product.ImageFileName,
		product.CreatedAt,
		product.CreatedBy,
		product.UpdatedAt,
		product.UpdatedBy,
		product.DeletedAt,
		product.DeletedBy,
		product.IsDeleted,
	)
	if err != nil {
		return err
	}

	return nil
}

func (repo *productRepository) GetProductById(ctx context.Context, id string) (*entity.Product, error) {
	var productEntity entity.Product
	row := repo.db.QueryRowContext(
		ctx,
		"SELECT id, name, description, price, image_file_name FROM product WHERE id = $1 AND is_deleted = false",
		id,
	)
	if row.Err() != nil {
		return nil, row.Err()
	}

	err := row.Scan(
		&productEntity.Id,
		&productEntity.Name,
		&productEntity.Description,
		&productEntity.Price,
		&productEntity.ImageFileName,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &productEntity, nil
}

func (repo *productRepository) UpdateProduct(ctx context.Context, product *entity.Product) error {
	_, err := repo.db.ExecContext(
		ctx,
		"UPDATE product SET name = $1, description= $2, price= $3, image_file_name= $4, updated_at= $5, updated_by= $6 WHERE id= $7",
		product.Name,
		product.Description,
		product.Price,
		product.ImageFileName,
		product.UpdatedAt,
		product.UpdatedBy,
		product.Id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (repo *productRepository) DeleteProduct(ctx context.Context, id string, deletedAt time.Time, deleteBy string) error {
	_, err := repo.db.ExecContext(
		ctx,
		"UPDATE product SET deleted_at = $1, deleted_by = $2, is_deleted = true WHERE id = $3",
		deletedAt,
		deleteBy,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (repo *productRepository) GetProductsPagination(ctx context.Context, pagination *common.PaginationRequest) ([]*entity.Product, *common.PaginationResponse, error) {
	row := repo.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM product WHERE is_deleted = false")
	if row.Err() != nil {
		return nil, nil, row.Err()
	}

	var totalCount int
	err := row.Scan(&totalCount)
	if err != nil {
		return nil, nil, err
	}

	offset := (pagination.CurrentPage - 1) * pagination.ItemPerPage
	totalPages := (totalCount + int(pagination.ItemPerPage) - 1) / int(pagination.ItemPerPage)

	rows, err := repo.db.QueryContext(
		ctx,
		"SELECT id, name, description, price, image_file_name FROM product WHERE is_deleted = false ORDER BY created_at DESC LIMIT $1 OFFSET $2",
		pagination.ItemPerPage,
		offset,
	)
	if err != nil {
		return nil, nil, err
	}

	var products []*entity.Product = make([]*entity.Product, 0)
	for rows.Next() {
		var product entity.Product

		err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.ImageFileName,
		)
		if err != nil {
			return nil, nil, err
		}

		products = append(products, &product)
	}

	paginationResponse := &common.PaginationResponse{
		CurrentPage:    pagination.CurrentPage,
		ItemPerPage:    pagination.ItemPerPage,
		TotalItemCount: int32(totalCount),
		TotalPageCount: int32(totalPages),
	}
	return products, paginationResponse, nil
}

func (repo *productRepository) GetProductsPaginationAdmin(ctx context.Context, pagination *common.PaginationRequest) ([]*entity.Product, *common.PaginationResponse, error) {
	row := repo.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM product WHERE is_deleted = false")
	if row.Err() != nil {
		return nil, nil, row.Err()
	}

	var totalCount int
	err := row.Scan(&totalCount)
	if err != nil {
		return nil, nil, err
	}

	offset := (pagination.CurrentPage - 1) * pagination.ItemPerPage
	totalPages := (totalCount + int(pagination.ItemPerPage) - 1) / int(pagination.ItemPerPage)

	allowedSorts := map[string]bool{
		"name":        true,
		"description": true,
		"price":       true,
	}

	orderQuery := "ORDER BY created_at DESC"
	if pagination.Sort != nil && allowedSorts[pagination.Sort.Field] {
		direction := "asc"
		if pagination.Sort.Direction == "desc" {
			direction = "desc"
		}

		orderQuery = fmt.Sprintf("ORDER BY %s %s", pagination.Sort.Field, direction)
	}

	baseQuery := fmt.Sprintf("SELECT id, name, description, price, image_file_name FROM product WHERE is_deleted = false %s LIMIT $1 OFFSET $2", orderQuery)
	rows, err := repo.db.QueryContext(
		ctx,
		baseQuery,
		pagination.ItemPerPage,
		offset,
	)
	if err != nil {
		return nil, nil, err
	}

	var products []*entity.Product = make([]*entity.Product, 0)
	for rows.Next() {
		var product entity.Product

		err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.ImageFileName,
		)
		if err != nil {
			return nil, nil, err
		}

		products = append(products, &product)
	}

	paginationResponse := &common.PaginationResponse{
		CurrentPage:    pagination.CurrentPage,
		ItemPerPage:    pagination.ItemPerPage,
		TotalItemCount: int32(totalCount),
		TotalPageCount: int32(totalPages),
	}
	return products, paginationResponse, nil
}

func (repo *productRepository) GetProductHighlight(ctx context.Context) ([]*entity.Product, error) {
	rows, err := repo.db.QueryContext(
		ctx,
		"SELECT id, name, description, price, image_file_name FROM product WHERE is_deleted = false ORDER BY created_at DESC LIMIT 3",
	)
	if err != nil {
		return nil, err
	}

	var products []*entity.Product = make([]*entity.Product, 0)
	for rows.Next() {
		var productEntity entity.Product

		err = rows.Scan(
			&productEntity.Id,
			&productEntity.Name,
			&productEntity.Description,
			&productEntity.Price,
			&productEntity.ImageFileName,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, &productEntity)
	}

	return products, nil
}

func NewProductRepository(db *sql.DB) IProductRepository {
	return &productRepository{
		db: db,
	}
}
