package controllers

import (
	"net/http"
	"net/http/httptest"
	"static-echo/middlewares"
	"strings"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	bookJson = `{"id": 1, "title": "The Psychology of Money", "author": "Morgan Housel"}`
)

func TestGetBooks(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books")
	
	if assert.NoError(t, GetBooks(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `{"book":[{"UserId":0, "author":"John Grisham", "id":1, "title":"A Time to Kill"},
		{"UserId":0, "author":"Edith Wharton", "id":2, "title":"The House of Mirth"}],
		"message":"Success retrieve book"}`, rec.Body.String())
	}
}

func TestGetBookByIdSuccess(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books/1")
	c.SetParamNames("id")
	c.SetParamValues("1")
	
	if assert.NoError(t, GetBookById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `{"book":{ "UserId":1, "author":"John Grisham", "id":1, "title":"A Time to Kill"},
		"message":"Success retrieve a book"}`, rec.Body.String())
	}
}

func TestGetBookByIdFailed(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books/1")
	c.SetParamNames("id")
	c.SetParamValues("0")
	
	if assert.NoError(t, GetBookById(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, `{"message":"Book not found"}`, rec.Body.String())
	}
}

func TestCreateBookSuccess(t *testing.T) {
	e := echo.New()
	e.Validator = &middlewares.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(bookJson))
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2NjM1NzQ4NjV9.YOFmVb-42-I2chou56c1A5DJ7p0J4RnXU7yB0LHY9aQ")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books")
	
	if assert.NoError(t, CreateBook(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `{"book":{"Id":1, "Title":"The Psychology of Money", "Author":"Morgan Housel"}, "message":"Success create book"}`, rec.Body.String())
	}
}

func TestCreateBookJWTFail(t *testing.T) {
	e := echo.New()
	e.Validator = &middlewares.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books/1")
	
	if assert.NoError(t, CreateBook(c)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.JSONEq(t, `{"message":"You are unauthorize to access this API"}`, rec.Body.String())
	}
}

func TestUpdateBook(t *testing.T) {
	e := echo.New()
	e.Validator = &middlewares.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2NjM1NzQ4NjV9.YOFmVb-42-I2chou56c1A5DJ7p0J4RnXU7yB0LHY9aQ")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books")
	c.SetParamNames("id", "title", "author")
	c.SetParamValues("1", "The Psychology of Money", "Morgan Housel")
	
	if assert.NoError(t, UpdateBook(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `{"book":{"Id":1, "Title":"The Psychology of Money", "Author":"Morgan Housel"}, "message":"Success create book"}`, rec.Body.String())
	}
}

func TestUpdateBookFail(t *testing.T) {
	e := echo.New()
	e.Validator = &middlewares.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books/1")
	
	if assert.NoError(t, UpdateBook(c)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.JSONEq(t, `{"message":"You are unauthorize to access this API"}`, rec.Body.String())
	}
}

func TestDeleteBookSuccess(t *testing.T) {
	e := echo.New()
	e.Validator = &middlewares.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2NjM1NzQ4NjV9.YOFmVb-42-I2chou56c1A5DJ7p0J4RnXU7yB0LHY9aQ")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books/1")
	c.SetParamNames("id")
	c.SetParamValues("1")
	
	if assert.NoError(t, DeleteBook(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `{"book": [{"UserId":2, "author":"Edith Wharton", "id":2, "title":"The House of Mirth"}],"message":"Success delete a book"}`, rec.Body.String())
	}
}

func TestDeleteFailed(t *testing.T) {
	e := echo.New()
	e.Validator = &middlewares.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2NjM1NzQ4NjV9.YOFmVb-42-I2chou56c1A5DJ7p0J4RnXU7yB0LHY9aQ")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books/1")
	c.SetParamNames("id")
	c.SetParamValues("3")
	
	if assert.NoError(t, DeleteBook(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, `{"message":"Book not found"}`, rec.Body.String())
	}
}