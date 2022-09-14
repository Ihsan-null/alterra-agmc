package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitPort() string{
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env file")
	}
	
	e := ":" + os.Getenv("APP_PORT")

	return e
}