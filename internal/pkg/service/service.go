package service

import (
	"context"
	"github.com/wisphes/book-shop/internal/models"
	"github.com/wisphes/book-shop/internal/pkg/repository"
)

type Authorization interface {
	CreateUser(ctx context.Context, user models.User) (int, error)
	GenerateToken(ctx context.Context, email, password string) (string, error)
	ParseToken(ctx context.Context, token string) (int, error)
	IsAdmin(ctx context.Context, id int) (bool, error)
}

type Basket interface {
}

type Book interface {
}

type Category interface {
	GetCategories(ctx context.Context) ([]models.Category, error)
	GetCategory(ctx context.Context, id int) (models.Category, error)
	CreateCategory(ctx context.Context, category models.Category) (models.Category, error)
	UpdateCategory(ctx context.Context, category models.Category) (models.Category, error)
}

type Service struct {
	Authorization
	Basket
	Book
	Category
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Category:      NewCategoryService(repos.Category),
	}
}
