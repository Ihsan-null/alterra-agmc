package routes

import (
	"static-echo/controllers"
	"static-echo/middlewares"

	"github.com/labstack/echo/v4"
)

func Book(e *echo.Echo) {
	g := e.Group("/v1")
	g.GET("/books", controllers.GetBooks)
	g.GET("/books/:id", controllers.GetBookById)
	g.Use(middlewares.JWTMiddleware(middlewares.JwtClaims{}, middlewares.JWT_SECRET))
	g.POST("/books", controllers.CreateBook)
	g.PUT("/books/:id", controllers.UpdateBook)
	g.DELETE("/books/:id", controllers.DeleteBook)
}