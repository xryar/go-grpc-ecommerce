package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/xryar/golang-grpc-ecommerce/internal/entity"
)

type IOrderRepository interface {
	GetNumbering(ctx context.Context, module string) (*entity.Numbering, error)
}

type orderRepository struct {
	db *sql.DB
}

func (or *orderRepository) GetNumbering(ctx context.Context, module string) (*entity.Numbering, error) {
	row := or.db.QueryRowContext(
		ctx,
		"SELECT module, number FROM numbering WHERE module = $1",
		module,
	)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var numbering entity.Numbering
	err := row.Scan(
		&numbering.Module,
		&numbering.Number,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &numbering, nil
}

func NewOrderRepository(db *sql.DB) IOrderRepository {
	return &orderRepository{
		db: db,
	}
}
