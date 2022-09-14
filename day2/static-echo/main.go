package main

import (
	"static-echo/config"
	"static-echo/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	routes.New(e)

	e.Start(config.InitPort())
}