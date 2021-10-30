// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/mhaqiw/haqi-coba-golang/domain"
	mock "github.com/stretchr/testify/mock"
)

// ProductStockRepository is an autogenerated mock type for the ProductStockRepository type
type ProductStockRepository struct {
	mock.Mock
}

// GetByWarehouseAndProduct provides a mock function with given fields: ctx, warehouseId, productId
func (_m *ProductStockRepository) GetByWarehouseAndProduct(ctx context.Context, warehouseId int, productId int) (domain.ProductStock, error) {
	ret := _m.Called(ctx, warehouseId, productId)

	var r0 domain.ProductStock
	if rf, ok := ret.Get(0).(func(context.Context, int, int) domain.ProductStock); ok {
		r0 = rf(ctx, warehouseId, productId)
	} else {
		r0 = ret.Get(0).(domain.ProductStock)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, warehouseId, productId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, ps
func (_m *ProductStockRepository) Update(ctx context.Context, ps domain.ProductStock) error {
	ret := _m.Called(ctx, ps)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ProductStock) error); ok {
		r0 = rf(ctx, ps)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
