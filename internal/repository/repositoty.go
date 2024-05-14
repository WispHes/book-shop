package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/wisphes/book-shop/internal/models"
)

type Authorization interface {
	CreateUser(ctx context.Context, user models.User) (int, error)
	GetUser(ctx context.Context, email, password string) (models.User, error)
	IsAdmin(ctx context.Context, id int) (models.User, error)
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

type Repository struct {
	Authorization
	Basket
	Book
	Category
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Category:      NewCategoryPostgres(db),
	}
}
