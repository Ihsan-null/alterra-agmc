package controllers

import (
	"echo-database/lib/database"
	"echo-database/middlewares"
	"echo-database/models"
	"echo-database/utils"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func Login(c echo.Context) error {
	var user models.Login
	err := ((&echo.DefaultBinder{}).BindBody(c, &user)); if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if err = c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	checkUser, err := database.GetUserByEmail(user.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid email or password",
		})
	}

	checkPassword := utils.CompareHashPassword(user.Password, checkUser.Password)
	if !checkPassword {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid email or password",
		})
	}
	token := middlewares.CreateToken(checkUser.ID)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "login success",
		"token": token,
	})
}

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
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	  }
	
	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	insertDb := models.User{
		Name: user.Name,
		Email: user.Email,
		Password: hash,
	}

	_, err = database.CreateUser(insertDb)
	if err != nil {
		log.Error(err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success create user",
	})
}

func UpdateUser(c echo.Context) error {
	var params models.RequestUpdateUser

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
	}

	authHeader := c.Request().Header.Get("Authorization")
	jwtClaims, err := middlewares.ParseJWTToken(authHeader)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "You are unauthorize to access this API",
		})
	}

	if int(jwtClaims.UserId) != id {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "You are unauthorize to access this API",
		})
	}

	err = ((&echo.DefaultBinder{}).BindBody(c, &params)); if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	hash, err := utils.HashPassword(params.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	updateDb := models.User{
		Name: params.Name,
		Email: params.Email,
		Password: hash,
	}
	data, err := database.UpdateUser(&id, &updateDb)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success update a user",
		"user": data,
	})
}

func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	authHeader := c.Request().Header.Get("Authorization")
	jwtClaims, err := middlewares.ParseJWTToken(authHeader)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "You are unauthorize to access this API",
		})
	}

	if int(jwtClaims.UserId) != id {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "You are unauthorize to access this API",
		})
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