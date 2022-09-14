package main

import (
	"echo-database/config"
	"echo-database/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	routes.New(e)

	config.InitDB()
	e.Start(config.InitPort())
}