package service

import (
	"context"
	"github.com/wisphes/book-shop/internal/models"
	"github.com/wisphes/book-shop/internal/repository"
)

type Book interface {
	GetBooks(ctx context.Context) ([]models.Book, error)
	GetBook(ctx context.Context, bookId int) (models.Book, error)
}

type BookService struct {
	repo *repository.BookPostgres
}

func NewBookService(repo *repository.BookPostgres) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetBooks(ctx context.Context) ([]models.Book, error) {
	return s.repo.GetBooks(ctx)
}

func (s *BookService) GetBook(ctx context.Context, bookId int) (models.Book, error) {
	return s.repo.GetBook(ctx, bookId)
}
