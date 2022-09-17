package config

import (
	"os"
)

func InitPort() string{
	e := ":" + os.Getenv("APP_PORT")

	return e
}