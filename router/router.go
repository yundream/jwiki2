package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// New ...
func New() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	return e
}
