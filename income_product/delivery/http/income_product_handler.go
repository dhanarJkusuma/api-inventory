package http

import (
	"net/http"

	"github.com/labstack/echo"
)

type ResponseError struct {
	Message string `json:"message"`
}

type IncomeProductHandler struct {
}

func (i *IncomeProductHandler) Ping(c echo.Context) error {
	response := struct {
		Message string
	}{
		Message: "Hello From Income Product",
	}
	return c.JSON(http.StatusOK, response)
}

func NewIncomeProductHandler(e *echo.Echo) {
	handler := &IncomeProductHandler{}
	e.GET("/income-product/ping", handler.Ping)
}
