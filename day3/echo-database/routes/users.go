package routes

import (
	"echo-database/controllers"
	"echo-database/middlewares"

	"github.com/labstack/echo/v4"
)

func User(e *echo.Echo) {
	g := e.Group("/v1")
	g.POST("/users", controllers.CreateUser)
	g.POST("/login", controllers.Login)
	g.Use(middlewares.JWTMiddleware(middlewares.JwtClaims{}, middlewares.JWT_SECRET))
	g.GET("/users", controllers.GetUsers)
	g.GET("/users/:id", controllers.GetUserById)
	g.PUT("/users/:id", controllers.UpdateUser)
	g.DELETE("/users/:id", controllers.DeleteUser)
}