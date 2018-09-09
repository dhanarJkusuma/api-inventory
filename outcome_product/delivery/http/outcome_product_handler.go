package http

import (
	"net/http"

	"github.com/labstack/echo"
)

type ResponseError struct {
	Message string `json:"message"`
}

type OutcomeProductHandler struct {
}

func (o *OutcomeProductHandler) Ping(c echo.Context) error {
	response := struct {
		Message string
	}{
		Message: "Hello From Outcome Product",
	}
	return c.JSON(http.StatusOK, response)
}

func NewOutcomeProductHandler(e *echo.Echo) {
	handler := &OutcomeProductHandler{}
	e.GET("/outcome-product/ping", handler.Ping)
}
