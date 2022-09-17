package main

import (
	"static-echo/config"
	"static-echo/middlewares"
	"static-echo/routes"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	routes.Book(e)
	routes.User(e)
	middlewares.LogMiddleware(e)
	e.Validator = &middlewares.CustomValidator{Validator: validator.New()}
	e.Start(config.InitPort())
}