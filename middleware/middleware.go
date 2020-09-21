package middleware

import (
	"github.com/labstack/echo/v4"
)

// ServerMiddleware ...
func ServerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		c.Response().Header().Set(echo.HeaderServer, "jwiki/2.0")
		return nil
	}
}
