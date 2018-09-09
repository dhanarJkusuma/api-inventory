package http

import (
	"net/http"

	"github.com/labstack/echo"
)

type ResponseError struct {
	Message string `json:"message"`
}

type ProductHandler struct {
}

func (p *ProductHandler) Ping(c echo.Context) error {
	response := struct {
		Message string
	}{
		Message: "Hello From Product",
	}
	return c.JSON(http.StatusOK, response)
}

func NewProductHandler(e *echo.Echo) {
	handler := &ProductHandler{}
	e.GET("/product/ping", handler.Ping)
}
