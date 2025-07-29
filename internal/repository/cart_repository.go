package repository

import (
	"context"
	"database/sql"
)

type ICartRepository interface {
	GetCartByProductAndUserId(ctx context.Context, productId, userId string)
}

type cartRepository struct {
	db *sql.DB
}

func (cr *cartRepository) GetCartByProductAndUserId(ctx context.Context, productId, userId string) {

}

func NewCartRepository(db *sql.DB) ICartRepository {
	return &cartRepository{
		db: db,
	}
}
