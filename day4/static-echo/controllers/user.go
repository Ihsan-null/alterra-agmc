package controllers

import (
	"net/http"
	"static-echo/middlewares"
	"static-echo/models"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error{
	var listUser []models.User
	var user models.User

	if err := c.Validate(user); err != nil {
		return err
	}

	listUser = append(listUser, models.User{Id: 1, Name: "Nunu", Email: "nunu@example.com", Password: "12345678"})
	listUser = append(listUser, models.User{Id: 2, Name: "Fian", Email: "fian@example.com", Password: "12345678"})
	var userId uint
	for _ , v := range listUser {
		if v.Email == user.Email && v.Password == user.Password {
			userId = v.Id
		}
	}
	if userId == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "User not found",
		})
	}
	token := middlewares.CreateToken(userId)
	
	return c.JSON(http.StatusOK, map[string]string{
		"message": "success",
		"token": token,
	})
}