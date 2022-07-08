package services

import (
	"context"
	"myapp/core/models"
	"myapp/core/ports"
)

type BookService struct {
	bookRepostory ports.BookRepository
}

func NewBookService(bookRepository ports.BookRepository) ports.BookService {
	return &BookService{
		bookRepostory: bookRepository,
	}
}

func (bookService *BookService) Create(ctx context.Context, book *models.Book) error {
	err := bookService.bookRepostory.Create(ctx, book)
	if err != nil {
		return err
	}
	return nil
}

func (bookService *BookService) Update(ctx context.Context, book *models.Book) error {
	err := bookService.bookRepostory.Update(ctx, book)
	if err != nil {
		return err
	}
	return nil
}

func (bookService *BookService) Delete(ctx context.Context, book *models.Book) error {
	err := bookService.bookRepostory.Delete(ctx, book)
	if err != nil {
		return err
	}
	return nil
}

func (bookService *BookService) FindAll(ctx context.Context) ([]models.Book, error) {
	books, err := bookService.bookRepostory.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (bookService *BookService) FindById(ctx context.Context, isbn string) (models.Book, error) {
	book, err := bookService.bookRepostory.FindById(ctx, isbn)
	if err != nil {
		return book, err
	}
	return book, nil
}
