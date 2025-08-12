package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/xryar/golang-grpc-ecommerce/internal/entity"
	"github.com/xryar/golang-grpc-ecommerce/pb/common"
	"github.com/xryar/golang-grpc-ecommerce/pkg/database"
)

type IOrderRepository interface {
	WithTransaction(tx *sql.Tx) IOrderRepository
	GetNumbering(ctx context.Context, module string) (*entity.Numbering, error)
	CreateOrder(ctx context.Context, order *entity.Order) error
	UpdateNumbering(ctx context.Context, numbering *entity.Numbering) error
	CreateOrderItem(ctx context.Context, orderItem *entity.OrderItem) error
	GetOrderById(ctx context.Context, orderId string) (*entity.Order, error)
	UpdateOrder(ctx context.Context, order *entity.Order) error
	GetListOrderAdminPagination(ctx context.Context, pagination *common.PaginationRequest) ([]*entity.Order, *common.PaginationResponse, error)
	GetListOrderPagination(ctx context.Context, pagination *common.PaginationRequest, userId string) ([]*entity.Order, *common.PaginationResponse, error)
}

type orderRepository struct {
	db database.DatabaseQuery
}

func (or *orderRepository) WithTransaction(tx *sql.Tx) IOrderRepository {
	return &orderRepository{
		db: tx,
	}
}

func (or *orderRepository) GetNumbering(ctx context.Context, module string) (*entity.Numbering, error) {
	row := or.db.QueryRowContext(
		ctx,
		"SELECT module, number FROM numbering WHERE module = $1 FOR UPDATE",
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

func (or *orderRepository) CreateOrder(ctx context.Context, order *entity.Order) error {
	_, err := or.db.ExecContext(
		ctx,
		"INSERT INTO \"order\" (id, number, user_id, order_status_code, user_full_name, address, phone_number, notes, total, expired_at, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, is_deleted, xendit_invoice_id, xendit_invoice_url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)",
		order.Id,
		order.Number,
		order.UserId,
		order.OrderStatusCode,
		order.UserFullName,
		order.Address,
		order.PhoneNumber,
		order.Notes,
		order.Total,
		order.ExpiredAt,
		order.CreatedAt,
		order.CreatedBy,
		order.UpdatedAt,
		order.UpdatedBy,
		order.DeletedAt,
		order.DeletedBy,
		order.IsDeleted,
		order.XenditInvoiceId,
		order.XenditInvoiceUrl,
	)
	if err != nil {
		return err
	}

	return nil
}

func (or *orderRepository) UpdateNumbering(ctx context.Context, numbering *entity.Numbering) error {
	_, err := or.db.ExecContext(
		ctx,
		"UPDATE numbering SET number = $1 WHERE module = $2",
		numbering.Number,
		numbering.Module,
	)
	if err != nil {
		return err
	}

	return nil
}

func (or *orderRepository) CreateOrderItem(ctx context.Context, orderItem *entity.OrderItem) error {
	_, err := or.db.ExecContext(
		ctx,
		"INSERT INTO order_item (id, product_id, product_name, product_image_file_name, product_price, quantity, order_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, is_deleted) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)",
		orderItem.Id,
		orderItem.ProductId,
		orderItem.ProductName,
		orderItem.ProductImageFileName,
		orderItem.ProductPrice,
		orderItem.Quantity,
		orderItem.OrderId,
		orderItem.CreatedAt,
		orderItem.CreatedBy,
		orderItem.UpdatedAt,
		orderItem.UpdatedBy,
		orderItem.DeletedAt,
		orderItem.DeletedBy,
		orderItem.IsDeleted,
	)
	if err != nil {
		return err
	}

	return nil
}

func (or *orderRepository) GetOrderById(ctx context.Context, orderId string) (*entity.Order, error) {
	row := or.db.QueryRowContext(
		ctx,
		"SELECT id FROM \"order\" WHERE id = $1 AND is_deleted = false",
		orderId,
	)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var order entity.Order
	err := row.Scan(
		&order.Id,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &order, nil
}

func (or *orderRepository) UpdateOrder(ctx context.Context, order *entity.Order) error {
	_, err := or.db.ExecContext(
		ctx,
		"UPDATE \"order\" SET updated_at = $1, updated_by = $2, xendit_paid_at = $3, xendit_payment_channel = $4, xendit_payment_method = $5, order_status_code = $6 WHERE id = $7",
		order.UpdatedAt,
		order.UpdatedBy,
		order.XenditPaidAt,
		order.XenditPaymentChannel,
		order.XenditPaymentMethod,
		order.OrderStatusCode,
		order.Id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (or *orderRepository) GetListOrderAdminPagination(ctx context.Context, pagination *common.PaginationRequest) ([]*entity.Order, *common.PaginationResponse, error) {
	row := or.db.QueryRowContext(
		ctx,
		"SELECT COUNT(*) FROM \"order\" WHERE is_deleted = false",
	)
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

	allowedSorts := map[string]string{
		"number":     "number",
		"customer":   "user_full_name",
		"total":      "total",
		"created_at": "created_at",
	}
	sort := "ORDER BY created_at DESC"
	if pagination.Sort != nil {
		direction := "ASC"
		sortField, ok := allowedSorts[pagination.Sort.Field]
		if ok {
			if pagination.Sort.Direction == "desc" {
				direction = "DESC"
			}
			sort = fmt.Sprintf("ORDER BY %s %s", sortField, direction)
		}
	}

	baseQuery := fmt.Sprintf("SELECT id, number, order_status_code, total, user_full_name, created_at, expired_at FROM \"order\" WHERE is_deleted = false %s LIMIT $1 OFFSET $2", sort)
	rows, err := or.db.QueryContext(
		ctx,
		baseQuery,
		pagination.ItemPerPage,
		offset,
	)
	if err != nil {
		return nil, nil, err
	}

	orders := make([]*entity.Order, 0)
	ids := make([]string, 0)
	orderItemsMap := make(map[string][]*entity.OrderItem)
	for rows.Next() {
		var orderEntity entity.Order
		err = rows.Scan(
			&orderEntity.Id,
			&orderEntity.Number,
			&orderEntity.OrderStatusCode,
			&orderEntity.Total,
			&orderEntity.UserFullName,
			&orderEntity.CreatedAt,
			&orderEntity.ExpiredAt,
		)
		if err != nil {
			return nil, nil, err
		}

		orders = append(orders, &orderEntity)
		ids = append(ids, fmt.Sprintf("'%s'", orderEntity.Id))
		orderItemsMap[orderEntity.Id] = make([]*entity.OrderItem, 0)
	}

	if len(orders) > 0 {
		idsJoined := strings.Join(ids, ", ")
		baseOrderItemQuery := fmt.Sprintf("SELECT product_id, product_name, product_price, quantity, order_id FROM order_item WHERE is_deleted = false AND order_id IN (%s)", idsJoined)
		rows, err = or.db.QueryContext(
			ctx,
			baseOrderItemQuery,
		)
		if err != nil {
			return nil, nil, err
		}

		for rows.Next() {
			var item entity.OrderItem
			err = rows.Scan(
				&item.ProductId,
				&item.ProductName,
				&item.ProductPrice,
				&item.Quantity,
				&item.OrderId,
			)
			if err != nil {
				return nil, nil, err
			}

			orderItemsMap[item.OrderId] = append(orderItemsMap[item.OrderId], &item)
		}

		for i, o := range orders {
			orders[i].Items = orderItemsMap[o.Id]
		}
	}

	var metadata common.PaginationResponse = common.PaginationResponse{
		CurrentPage:    pagination.CurrentPage,
		TotalPageCount: int32(totalPages),
		ItemPerPage:    pagination.ItemPerPage,
		TotalItemCount: int32(totalCount),
	}

	return orders, &metadata, nil
}

func (or *orderRepository) GetListOrderPagination(ctx context.Context, pagination *common.PaginationRequest, userId string) ([]*entity.Order, *common.PaginationResponse, error) {
	row := or.db.QueryRowContext(
		ctx,
		"SELECT COUNT(*) FROM \"order\" WHERE is_deleted = false AND user_id = $1",
		userId,
	)
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

	allowedSorts := map[string]string{
		"number":     "number",
		"customer":   "user_full_name",
		"total":      "total",
		"created_at": "created_at",
	}
	sort := "ORDER BY created_at DESC"
	if pagination.Sort != nil {
		direction := "ASC"
		sortField, ok := allowedSorts[pagination.Sort.Field]
		if ok {
			if pagination.Sort.Direction == "desc" {
				direction = "DESC"
			}
			sort = fmt.Sprintf("ORDER BY %s %s", sortField, direction)
		}
	}

	baseQuery := fmt.Sprintf("SELECT id, number, order_status_code, total, user_full_name, created_at, expired_at, xendit_invoice_url FROM \"order\" WHERE is_deleted = false AND user_id = $1 %s LIMIT $2 OFFSET $3", sort)
	rows, err := or.db.QueryContext(
		ctx,
		baseQuery,
		userId,
		pagination.ItemPerPage,
		offset,
	)
	if err != nil {
		return nil, nil, err
	}

	orders := make([]*entity.Order, 0)
	ids := make([]string, 0)
	orderItemsMap := make(map[string][]*entity.OrderItem)
	for rows.Next() {
		var orderEntity entity.Order
		err = rows.Scan(
			&orderEntity.Id,
			&orderEntity.Number,
			&orderEntity.OrderStatusCode,
			&orderEntity.Total,
			&orderEntity.UserFullName,
			&orderEntity.CreatedAt,
			&orderEntity.ExpiredAt,
			&orderEntity.XenditInvoiceUrl,
		)
		if err != nil {
			return nil, nil, err
		}

		orders = append(orders, &orderEntity)
		ids = append(ids, fmt.Sprintf("'%s'", orderEntity.Id))
		orderItemsMap[orderEntity.Id] = make([]*entity.OrderItem, 0)
	}

	if len(orders) > 0 {
		idsJoined := strings.Join(ids, ", ")
		baseOrderItemQuery := fmt.Sprintf("SELECT product_id, product_name, product_price, quantity, order_id FROM order_item WHERE is_deleted = false AND order_id IN (%s)", idsJoined)
		rows, err = or.db.QueryContext(
			ctx,
			baseOrderItemQuery,
		)
		if err != nil {
			return nil, nil, err
		}

		for rows.Next() {
			var item entity.OrderItem
			err = rows.Scan(
				&item.ProductId,
				&item.ProductName,
				&item.ProductPrice,
				&item.Quantity,
				&item.OrderId,
			)
			if err != nil {
				return nil, nil, err
			}

			orderItemsMap[item.OrderId] = append(orderItemsMap[item.OrderId], &item)
		}

		for i, o := range orders {
			orders[i].Items = orderItemsMap[o.Id]
		}
	}

	var metadata common.PaginationResponse = common.PaginationResponse{
		CurrentPage:    pagination.CurrentPage,
		TotalPageCount: int32(totalPages),
		ItemPerPage:    pagination.ItemPerPage,
		TotalItemCount: int32(totalCount),
	}

	return orders, &metadata, nil
}

func NewOrderRepository(db database.DatabaseQuery) IOrderRepository {
	return &orderRepository{
		db: db,
	}
}
