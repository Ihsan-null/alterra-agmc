package routes

import (
	"static-echo/controllers"

	"github.com/labstack/echo/v4"
)

func User(e *echo.Echo) {
	g := e.Group("/v1")
	g.POST("/login", controllers.Login)
}