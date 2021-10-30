package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/labstack/gommon/log"

	"github.com/mhaqiw/haqi-coba-golang/domain"
)

type productStockRepository struct {
	Conn *sql.DB
}

func NewMProductStockRepository(Conn *sql.DB) domain.ProductStockRepository {
	return &productStockRepository{Conn}
}

func (p productStockRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]domain.ProductStock, error) {
	rows, err := p.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	result := make([]domain.ProductStock, 0)
	for rows.Next() {
		t := domain.ProductStock{}
		err = rows.Scan(
			&t.ID,
			&t.ProductID,
			&t.WarehouseID,
			&t.Stock,
			&t.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (p productStockRepository) GetByWarehouseAndProduct(ctx context.Context, warehouseId int, productId int) (res domain.ProductStock, err error) {
	query := `SELECT id, product_id, warehouse_id, stock, updated_at FROM product_stock WHERE warehouse_id = $1 and product_id = $2`

	list, err := p.fetch(ctx, query, warehouseId, productId)
	if err != nil {
		return domain.ProductStock{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return domain.ProductStock{}, fmt.Errorf("product_stock not found")
	}

	return
}

func (p *productStockRepository) Update(ctx context.Context, ps domain.ProductStock) (err error) {
	query := `UPDATE product_stock set warehouse_id= $1, product_id= $2, stock= $3, updated_at= $4 WHERE id= $5`

	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, ps.WarehouseID, ps.ProductID, ps.Stock, ps.UpdatedAt, ps.ID)
	if err != nil {
		return
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affect != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", affect)
		return
	}

	return
}


func (p productStockRepository) GetByProductWarehouse(warehouseId string, productId string) (domain.ProductStock, error) {
	var productStock domain.ProductStock
	err := p.Conn.QueryRow(`SELECT id, product_id, warehouse_id, stocks, updated_at FROM product_stock WHERE warehouse_id  
	= ? and product_id = ?`, warehouseId, productId).Scan(&productStock.ID,
		&productStock.ProductID,
		&productStock.WarehouseID,
		&productStock.Stock,
		&productStock.UpdatedAt)
	return productStock, err
}
