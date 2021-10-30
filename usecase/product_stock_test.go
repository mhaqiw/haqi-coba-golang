package usecase_test

import (
	"context"
	"fmt"
	"github.com/mhaqiw/haqi-coba-golang/domain"
	mocks "github.com/mhaqiw/haqi-coba-golang/mocks/domain"
	"github.com/mhaqiw/haqi-coba-golang/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func Test_productStockUsecase_Update(t *testing.T) {
	mockProductStockRepo := new(mocks.ProductStockRepository)
	mockProductRepo := new(mocks.ProductRepository)
	mockWarehouseRepo := new(mocks.WarehouseRepository)
	mockProductStock := domain.ProductStock{
		ID:          1,
		ProductID:   1,
		WarehouseID: 1,
	}

	t.Run("success", func(t *testing.T) {
		mockProductRepo.On("CheckProductByID", mock.Anything).Return(nil).Once()
		mockWarehouseRepo.On("CheckWarehouseByID", mock.Anything).Return(nil).Once()
		mockProductStockRepo.On("GetByWarehouseAndProduct", mock.Anything, mock.Anything, mock.Anything).Return(mockProductStock, nil)
		mockProductStockRepo.On("Update", mock.Anything, mock.Anything).Return(nil)

		u := usecase.NewProductStockUsecase(mockProductStockRepo, mockProductRepo, mockWarehouseRepo, time.Second*10)

		request := domain.ProductStockUpdateRequest{
			ProductID:   1,
			WarehouseID: 1,
			AddStock:    1,
		}
		id, err := u.Update(context.TODO(), request)
		assert.Equal(t, id, 1)
		assert.NotEmpty(t, id)
		assert.NoError(t, err)
	})

	t.Run("stock-under-0", func(t *testing.T) {
		mockProductRepo.On("CheckProductByID", mock.Anything).Return(nil).Once()
		mockWarehouseRepo.On("CheckWarehouseByID", mock.Anything).Return(nil).Once()
		mockProductStockRepo.On("GetByWarehouseAndProduct", mock.Anything, mock.Anything, mock.Anything).Return(mockProductStock, nil)
		mockProductStockRepo.On("Update", mock.Anything, mock.Anything).Return(nil)

		u := usecase.NewProductStockUsecase(mockProductStockRepo, mockProductRepo, mockWarehouseRepo, time.Second*10)

		request := domain.ProductStockUpdateRequest{
			ProductID:   1,
			WarehouseID: 1,
			AddStock:    -10,
		}
		id, err := u.Update(context.TODO(), request)
		assert.Empty(t, id)
		assert.Error(t, err)
	})


	t.Run("warehouse-invalid", func(t *testing.T) {

		errWarehouseInvalid := fmt.Errorf("warehouse invalid")
		mockProductRepo.On("CheckProductByID", mock.Anything).Return(nil).Once()
		mockWarehouseRepo.On("CheckWarehouseByID", mock.Anything).Return(errWarehouseInvalid).Once()
		mockProductStockRepo.On("GetByWarehouseAndProduct", mock.Anything, mock.Anything, mock.Anything).Return(mockProductStock, nil)
		mockProductStockRepo.On("Update", mock.Anything, mock.Anything).Return(nil)

		u := usecase.NewProductStockUsecase(mockProductStockRepo, mockProductRepo, mockWarehouseRepo, time.Second*10)

		request := domain.ProductStockUpdateRequest{
			ProductID:   1,
			WarehouseID: 1,
			AddStock:    1,
		}
		id, err := u.Update(context.TODO(), request)
		assert.Empty(t, id)
		assert.Error(t, err)
		assert.Equal(t, err, errWarehouseInvalid)
	})
}
