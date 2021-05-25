package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"golang-clean-architecture/core/controllers"
)

func NewRouter(e *echo.Echo, c controllers.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", func(context echo.Context) error { return c.User.GetUsers(context) })
	e.POST("/users", func(context echo.Context) error { return c.User.CreateUser(context) })

	return e
}
