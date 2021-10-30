package domain

import (
	"context"
	"time"
)

type ProductStock struct {
	ID          int       `json:"id"`
	ProductID   int    `json:"product_id"`
	WarehouseID int    `json:"warehouse_id"`
	Stock       int       `json:"stock"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductStockUpdateRequest struct {
	ProductID   int `json:"product_id"`
	WarehouseID int `json:"warehouse_id"`
	AddStock    int    `json:"add_stock"`
}

type ProductStockRepository interface {
	GetByWarehouseAndProduct(ctx context.Context, warehouseId int, productId int) (res ProductStock, err error)
	Update(ctx context.Context, ps ProductStock) (err error)
}

type ProductStockUsecase interface {
	Update(ctx context.Context, request ProductStockUpdateRequest) (int,error)
}