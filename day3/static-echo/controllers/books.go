package controllers

import (
	"net/http"
	"static-echo/middlewares"
	"static-echo/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GetBooks(c echo.Context) error {
	var books []models.Book
	books = append(books, models.Book{Id: 1, Title: "A Time to Kill", Author: "John Grisham"})
	books = append(books, models.Book{Id: 2, Title: "The House of Mirth", Author: "Edith Wharton"})
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success retrieve book",
		"book": books,
	})
}

func GetBookById(c echo.Context) error {
	var books []models.Book
	var data models.Book

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
	}

	books = append(books, models.Book{Id: 1, Title: "A Time to Kill", Author: "John Grisham", UserId: 1})
	books = append(books, models.Book{Id: 2, Title: "The House of Mirth", Author: "Edith Wharton", UserId: 2})
	for _ , book := range books {
		if book.Id == id {
			data = book
		}
	}
	if data.Id == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Book not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success retrieve a book",
		"book": data,
	})
}

func CreateBook(c echo.Context) error {
	var books models.Book
	var data models.Book

	authHeader := c.Request().Header.Get("Authorization")
	jwtClaims, err := middlewares.ParseJWTToken(authHeader)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "You are unauthorize to access this API",
		})
	}

	err = ((&echo.DefaultBinder{}).BindQueryParams(c, &books)); if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	if err = c.Validate(books); err != nil {
		return err
	}
	data = models.Book{Id: books.Id, Title: books.Title, Author: books.Author, UserId: jwtClaims.UserId}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success create book",
		"book": data,
	})
}

func UpdateBook(c echo.Context) error {
	var updateBook models.UpdateBook
	var booksData []models.Book
	var data models.Book

	authHeader := c.Request().Header.Get("Authorization")
	jwtClaims, err := middlewares.ParseJWTToken(authHeader)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "You are unauthorize to access this API",
		})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
	}

	err = ((&echo.DefaultBinder{}).BindQueryParams(c, &updateBook)); if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	if err = c.Validate(updateBook); err != nil {
		return err
	}

	booksData = append(booksData, models.Book{Id: 1, Title: "A Time to Kill", Author: "John Grisham", UserId: 1})
	booksData = append(booksData, models.Book{Id: 2, Title: "The House of Mirth", Author: "Edith Wharton", UserId: 2})
	for _ , book := range booksData {
		if book.Id == id {
			if book.UserId != jwtClaims.UserId {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "You are unauthorize to update this data",
				})
			}
			book.Id = id
			book.Title = updateBook.Title
			book.Author = updateBook.Author
			data = book
		}
	}

	if data.Id == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Book not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success update a book",
		"book": data,
	})
}

func DeleteBook(c echo.Context) error {
	var booksData []models.Book

	authHeader := c.Request().Header.Get("Authorization")
	jwtClaims, err := middlewares.ParseJWTToken(authHeader)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "You are unauthorize to access this API",
		})
	}

	var id int
	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
	}

	booksData = append(booksData, models.Book{Id: 1, Title: "A Time to Kill", Author: "John Grisham", UserId: 1})
	booksData = append(booksData, models.Book{Id: 2, Title: "The House of Mirth", Author: "Edith Wharton", UserId: 2})
	var checkExist bool
	for _ , book := range booksData {
		if book.Id == id {
			if book.UserId != jwtClaims.UserId {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "You are unauthorize to delete this data",
				})
			}
			booksData = append(booksData[:id-1], booksData[id:]...)
			checkExist = true
		}
	}

	if !checkExist {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Book not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success delete a book",
		"book": booksData,
	})
}