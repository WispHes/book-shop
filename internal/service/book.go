package service

import (
	"context"
	"github.com/wisphes/book-shop/internal/models"
)

type BookService struct {
	repo BookRepository
}

func NewBookService(repo BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetBooks(ctx context.Context) ([]models.Book, error) {
	return s.repo.GetBooks(ctx)
}

func (s *BookService) GetBook(ctx context.Context, bookId int) (models.Book, error) {
	return s.repo.GetBook(ctx, bookId)
}

func (s *BookService) CreateBook(ctx context.Context, book models.Book) (models.Book, error) {
	return s.repo.CreateBook(ctx, book)
}

func (s *BookService) UpdateBook(ctx context.Context, book models.Book) (models.Book, error) {
	return s.repo.UpdateBook(ctx, book)
}

func (s *BookService) DeleteBook(ctx context.Context, bookId int) error {
	return s.repo.DeleteBook(ctx, bookId)
}
