package http

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, controller UserController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", func(context echo.Context) error { return controller.GetUsers(context) })
	e.POST("/users", func(context echo.Context) error { return controller.CreateUser(context) })

	return e
}
