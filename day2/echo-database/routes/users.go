package routes

import (
	"echo-database/controllers"

	"github.com/labstack/echo/v4"
)

func New(e *echo.Echo) {
	g := e.Group("/v1")
	g.GET("/users", controllers.GetUsers)
	g.GET("/users/:id", controllers.GetUserById)
	g.POST("/users", controllers.CreateUser)
	g.PUT("/users/:id", controllers.UpdateUser)
	g.DELETE("/users/:id", controllers.DeleteUser)
}