package http

import (
	"net/http"

	"github.com/labstack/echo"
)

type ResponseError struct {
	Message string `json:"message"`
}

type TransactionHandler struct {
}

func (t *TransactionHandler) Ping(c echo.Context) error {
	response := struct {
		Message string
	}{
		Message: "Hello From Transaction",
	}
	return c.JSON(http.StatusOK, response)
}

func NewTransactionHandler(e *echo.Echo) {
	handler := &TransactionHandler{}
	e.GET("/transaction/ping", handler.Ping)
}
