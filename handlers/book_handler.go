package handlers

import (
	"context"
	"myapp/core/models"
	"myapp/core/ports"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	bookService ports.BookService
}

func NewBookHandler(bookService ports.BookService) *BookHandler {
	return &BookHandler{
		bookService: bookService,
	}
}

func (bookHandler *BookHandler) Create(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*5)
	defer cancel()

	var book models.Book
	if err := c.Bind(&book); err != nil {
		c.JSON(http.StatusBadRequest, "invalid data")
		return err
	}

	err := bookHandler.bookService.Create(ctx, &book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		return err
	}

	return c.JSON(http.StatusCreated, book)
}

func (bookHandler *BookHandler) Update(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*5)
	defer cancel()

	isbn := c.Param("isbn")

	var book models.Book
	if err := c.Bind(&book); err != nil {
		c.JSON(http.StatusBadRequest, "invalid data")
		return err
	}

	bookUpdate := new(models.Book)
	bookUpdate.Isbn = isbn
	bookUpdate.Title = book.Title
	bookUpdate.Author = book.Author
	bookUpdate.Price = book.Price

	errUpdate := bookHandler.bookService.Update(ctx, bookUpdate)
	if errUpdate != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		return errUpdate
	}

	return c.JSON(http.StatusOK, bookUpdate)
}

func (bookHandler *BookHandler) Delete(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*5)
	defer cancel()

	isbn := c.Param("isbn")

	book, err := bookHandler.bookService.FindById(ctx, isbn)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{"error": "Data Not Found"})
		return err
	}

	errDelete := bookHandler.bookService.Delete(ctx, &book)
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		return errDelete
	}

	return c.JSON(http.StatusOK, book)
}

func (bookHandler *BookHandler) FindAll(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*5)
	defer cancel()

	books, err := bookHandler.bookService.FindAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		return err
	}

	return c.JSON(http.StatusOK, books)
}

func (bookHandler *BookHandler) FindById(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*5)
	defer cancel()

	isbn := c.Param("isbn")

	book, err := bookHandler.bookService.FindById(ctx, isbn)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{"error": "Data Not Found"})
		return err
	}

	return c.JSON(http.StatusOK, book)
}
