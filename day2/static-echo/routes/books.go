package routes

import (
	"static-echo/controllers"

	"github.com/labstack/echo/v4"
)

func New(e *echo.Echo) {
	g := e.Group("/v1")
	g.GET("/books", controllers.GetBooks)
	g.GET("/books/:id", controllers.GetBookById)
	g.POST("/books", controllers.CreateBook)
	g.PUT("/books/:id", controllers.UpdateBook)
	g.DELETE("/books/:id", controllers.DeleteBook)
}