package repository

import (
	"fmt"
	"github.com/mhaqiw/haqi-coba-golang/domain"
	"net/http"
)

type productRepository struct {
}

func NewProductRepository() domain.ProductRepository {
	return &productRepository{}
}

func (p productRepository) CheckProductByID(id int) error {
	responseProductDetail, err := http.Get(fmt.Sprintf("https://617a0c5dcb1efe001700fc3f.mockapi.io/api/v1/products/%d", id))
	if err != nil {
		return err
	}
	if responseProductDetail.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid Product")
	}

	return nil
}