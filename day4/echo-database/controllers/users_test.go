package controllers

import (
	"echo-database/middlewares"
	"echo-database/config"
	"log"
	"net/http"
	"net/http/httptest"

	"strings"
	"testing"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var(
	userForm = `name=zack&email=zack@example.com&password=12345678`
)
func initTest() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading env file")
	}
	config.InitDB()
}

func TestCreateUserSuccess(t *testing.T) {
	e := echo.New()
	e.Validator = &middlewares.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userForm))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/users")
	
	if assert.NoError(t, CreateUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `{"message":"Success create user"}`, rec.Body.String())
	}
}

func TestCreateUserInvalidForm(t *testing.T) {
	e := echo.New()
	e.Validator = &middlewares.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/users")
	
	if assert.NoError(t, CreateUser(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.True(t, strings.HasPrefix(rec.Body.String(), `{"message":"Key:`))
	}
}
func TestGetUsers(t *testing.T) {
	initTest()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2NjM1NzQ4NjV9.YOFmVb-42-I2chou56c1A5DJ7p0J4RnXU7yB0LHY9aQ")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/users")
	
	if assert.NoError(t, GetUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.True(t, strings.HasPrefix(rec.Body.String(), `{"message":"Success retrieve users","user":`))
	}
}

func TestGetBookByIdSuccess(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/users/4")
	c.SetParamNames("id")
	c.SetParamValues("4")
	
	if assert.NoError(t, GetUserById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.True(t, strings.HasPrefix(rec.Body.String(), `{"message":"Success retrieve users","user":`))
	}
}