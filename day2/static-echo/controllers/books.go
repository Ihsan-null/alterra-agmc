package controllers

import (
	"net/http"
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

	books = append(books, models.Book{Id: 1, Title: "A Time to Kill", Author: "John Grisham"})
	books = append(books, models.Book{Id: 2, Title: "The House of Mirth", Author: "Edith Wharton"})
	for _ , book := range books {
		if book.Id == id {
			data = book
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success retrieve a book",
		"book": data,
	})
}

func CreateBook(c echo.Context) error {
	var books models.Book
	var data models.Book
	err := ((&echo.DefaultBinder{}).BindQueryParams(c, &books)); if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	data = models.Book{Id: books.Id, Title: books.Title, Author: books.Author}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success create book",
		"book": data,
	})
}

func UpdateBook(c echo.Context) error {
	var updateBook models.UpdateBook
	var booksData []models.Book
	var data models.Book

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
	}

	err = ((&echo.DefaultBinder{}).BindQueryParams(c, &updateBook)); if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	booksData = append(booksData, models.Book{Id: 1, Title: "A Time to Kill", Author: "John Grisham"})
	booksData = append(booksData, models.Book{Id: 2, Title: "The House of Mirth", Author: "Edith Wharton"})
	for _ , book := range booksData {
		if book.Id == id {
			book.Id = id
			book.Title = updateBook.Title
			book.Author = updateBook.Author
			data = book
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success update a book",
		"book": data,
	})
}

func DeleteBook(c echo.Context) error {
	var delete models.UpdateBook
	var booksData []models.Book

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
	}

	err = ((&echo.DefaultBinder{}).BindQueryParams(c, &delete)); if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	booksData = append(booksData, models.Book{Id: 1, Title: "A Time to Kill", Author: "John Grisham"})
	booksData = append(booksData, models.Book{Id: 2, Title: "The House of Mirth", Author: "Edith Wharton"})
	for _ , book := range booksData {
		if book.Id == id {
			booksData = append(booksData[:id-1], booksData[id:]...)
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success delete a book",
		"book": booksData,
	})
}