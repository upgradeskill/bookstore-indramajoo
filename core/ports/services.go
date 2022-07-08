package ports

import (
	"context"
	"myapp/core/models"
)

type BookService interface {
	Create(ctx context.Context, book *models.Book) error
	Update(ctx context.Context, book *models.Book) error
	Delete(ctx context.Context, book *models.Book) error
	FindAll(ctx context.Context) ([]models.Book, error)
	FindById(ctx context.Context, isbn string) (models.Book, error)
}
