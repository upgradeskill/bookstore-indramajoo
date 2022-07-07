package mysql

import (
	"context"
	"myapp/core/models"
	"myapp/core/ports"

	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(DB *gorm.DB) ports.BookRepository {
	return &BookRepository{DB: DB}
}

func (m *BookRepository) Create(ctx context.Context, book *models.Book) error {
	if err := m.DB.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func (m *BookRepository) Update(ctx context.Context, book *models.Book) error {
	if err := m.DB.Where("isbn", book.Isbn).UpdateColumns(book).Error; err != nil {
		return err
	}
	return nil
}

func (m *BookRepository) Delete(ctx context.Context, book *models.Book) error {
	if err := m.DB.Where("isbn", book.Isbn).Delete(book).Error; err != nil {
		return err
	}
	return nil
}

func (m *BookRepository) FindAll(ctx context.Context) ([]models.Book, error) {
	var books []models.Book
	if err := m.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (m *BookRepository) FindById(ctx context.Context, isbn string) (models.Book, error) {
	var book models.Book
	if err := m.DB.Where("isbn", isbn).First(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}
