package usecase

import (
	"context"
	"fmt"
	"github.com/mhaqiw/haqi-coba-golang/domain"
	"time"
)

type productStockUsecase struct {
	productStockRepo domain.ProductStockRepository
	productRepo      domain.ProductRepository
	warehouseRepo    domain.WarehouseRepository
	contextTimeout   time.Duration
}

func NewProductStockUsecase(ps domain.ProductStockRepository, p domain.ProductRepository, w domain.WarehouseRepository, timeout time.Duration) domain.ProductStockUsecase {
	return &productStockUsecase{
		productStockRepo: ps,
		productRepo:      p,
		warehouseRepo:    w,
		contextTimeout:   timeout,
	}
}

func (p productStockUsecase) Update(ctx context.Context, request domain.ProductStockUpdateRequest) (int,error) {
	err := p.productRepo.CheckProductByID(request.ProductID)
	if err != nil {
		return 0, err
	}

	err = p.warehouseRepo.CheckWarehouseByID(request.WarehouseID)
	if err != nil {
		return 0, err
	}

	ps, err := p.productStockRepo.GetByWarehouseAndProduct(ctx, request.WarehouseID, request.ProductID)
	if err != nil {
		return 0, err
	}

	newStock := ps.Stock + request.AddStock
	if newStock <= 0 {
		return 0, fmt.Errorf("product stock cannot less than 1")
	}

	ps.Stock = newStock
	ps.UpdatedAt = time.Now()
	err = p.productStockRepo.Update(ctx, ps)
	if err != nil {
		return 0, err
	}

	return ps.ID,nil
}
