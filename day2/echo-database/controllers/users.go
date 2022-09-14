package controllers

import (
	"echo-database/lib/database"
	"echo-database/models"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GetUsers(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		log.Error(err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success retrieve users",
		"user": users,
	})
}

func GetUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
	}

	users, err := database.GetUserById(&id)
	if err != nil {
		log.Error(err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success retrieve users",
		"user": users,
	})
}

func CreateUser(c echo.Context) error {
	var user models.User
	err := ((&echo.DefaultBinder{}).BindBody(c, &user)); if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	data, err := database.CreateUser(user)
	if err != nil {
		log.Error(err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success create user",
		"user": data,
	})
}

func UpdateUser(c echo.Context) error {
	var params models.UpdateUser

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
	}
	err = ((&echo.DefaultBinder{}).BindQueryParams(c, &params)); if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	data, err := database.UpdateUser(&id, &params)
	if err != nil {
		log.Error(err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success update a user",
		"user": data,
	})
}

func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
	}

	data, err := database.DeleteUser(&id)
	if err != nil {
		log.Error(err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success delete a user",
		"user": data,
	})
}