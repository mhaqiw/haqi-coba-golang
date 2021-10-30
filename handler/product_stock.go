package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mhaqiw/haqi-coba-golang/domain"
)

type ResponseError struct {
	Message string `json:"message"`
}

type ProductStockHandler struct {
	PSUsecase domain.ProductStockUsecase
}

func NewProductStockHandler(e *echo.Echo, ps domain.ProductStockUsecase) {
	handler := &ProductStockHandler{
		PSUsecase: ps,
	}
	e.PUT("/product-stock", handler.Update)
}

func (a *ProductStockHandler) Update(c echo.Context) error {
	var request domain.ProductStockUpdateRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	id, err := a.PSUsecase.Update(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, id)
}
