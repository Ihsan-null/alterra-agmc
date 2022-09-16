package main

import (
	"echo-database/config"
	"echo-database/middlewares"
	"echo-database/routes"
	"log"

	"github.com/joho/godotenv"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env file")
	}
	
	e := echo.New()
	routes.User(e)
	middlewares.LogMiddleware(e)
	e.Validator = &middlewares.CustomValidator{Validator: validator.New()}
	config.InitDB()
	e.Start(config.InitPort())
}