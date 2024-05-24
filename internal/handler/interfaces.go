package handler

import (
	"context"
	"github.com/wisphes/book-shop/internal/models"
)

type BookService interface {
	GetBooks(ctx context.Context) ([]models.Book, error)
	GetBook(ctx context.Context, bookId int) (models.Book, error)
	CreateBook(ctx context.Context, book models.Book) (models.Book, error)
	UpdateBook(ctx context.Context, book models.Book) (models.Book, error)
	DeleteBook(ctx context.Context, bookId int) error
}

type UserService interface {
	CreateUser(ctx context.Context, user models.User) (int, error)
	GenerateToken(ctx context.Context, email, password string) (string, error)
	ParseToken(ctx context.Context, token string) (int, error)
	IsAdmin(ctx context.Context, userId int) error
	UserIdentity(ctx context.Context, header string) (int, error)
}

type CategoryService interface {
	GetCategories(ctx context.Context) ([]models.Category, error)
	GetCategory(ctx context.Context, catId int) (models.Category, error)
	CreateCategory(ctx context.Context, category models.Category) (models.Category, error)
	UpdateCategory(ctx context.Context, category models.Category) (models.Category, error)
	DeleteCategory(ctx context.Context, catId int) error
}

type BasketService interface {
	GetBasket(ctx context.Context, userId int) (models.Basket, error)
	ToPayBasket(ctx context.Context, userId int) error
	UpdateBasket(ctx context.Context, userId, bookId int, method string) (models.Basket, error)
}
