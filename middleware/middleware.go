package middleware

import "github.com/labstack/echo"

type goMiddleware struct {
	// another stuff , may be needed by middleware
}

func (m *goMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}

func InitMiddleware() *goMiddleware {
	return &goMiddleware{}
}
