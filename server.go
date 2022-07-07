package main

import (
	"myapp/config"
	"myapp/core/services"
	"myapp/handlers"
	"myapp/repositories/mysql"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	DB, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	bookRepository := mysql.NewBookRepository(DB)
	bookService := services.NewBookService(bookRepository)
	bookHandler := handlers.NewBookHandler(bookService)

	e.POST("/books", bookHandler.Create)
	e.PUT("/books/:isbn", bookHandler.Update)
	e.DELETE("/books/:isbn", bookHandler.Delete)
	e.GET("/books", bookHandler.FindAll)
	e.GET("/books/:isbn", bookHandler.FindById)

	e.Logger.Fatal(e.Start(":3000"))
}
