package repository

import (
	"fmt"
	"github.com/mhaqiw/haqi-coba-golang/domain"
	"net/http"
)

type warehouseRepository struct {
}

func (p warehouseRepository) CheckWarehouseByID(id int) error {
	responseWarehouseDetail, err := http.Get(fmt.Sprintf("https://617a0c5dcb1efe001700fc3f.mockapi.io/api/v1/warehouse/%d", id))
	if err != nil {
		return err
	}
	if responseWarehouseDetail.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid Warehouse")
	}

	return nil}

func NewWarehouseRepository() domain.WarehouseRepository {
	return &warehouseRepository{}
}