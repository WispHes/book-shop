package dep

import (
	"github.com/jmoiron/sqlx"
	"github.com/wisphes/book-shop/internal/handler"
	"github.com/wisphes/book-shop/internal/repository"
	"github.com/wisphes/book-shop/internal/service"
)

type Dependencies struct {
	UserRepo   *repository.UserPostgres
	CatRepo    *repository.CategoryPostgres
	BookRepo   *repository.BookPostgres
	BasketRepo *repository.BasketPostgres

	UserService   *service.UserService
	CatService    *service.CategoryService
	BookService   *service.BookService
	BasketService *service.BasketService

	UserHandler   *handler.UserHandler
	CatHandler    *handler.CategoryHandler
	BookHandler   *handler.BookHandler
	BasketHandler *handler.BasketHandler

	_ struct{}
}

func NewDependencies(db *sqlx.DB) *Dependencies {
	// repository
	userRepo := repository.NewUserPostgres(db)
	catRepo := repository.NewCategoryPostgres(db)
	bookRepo := repository.NewBookPostgres(db)
	basketRepo := repository.NewBasketPostgres(db)

	// service
	userService := service.NewUserService(userRepo)
	catService := service.NewCategoryService(catRepo)
	bookService := service.NewBookService(bookRepo)
	basketService := service.NewBasketService(basketRepo)

	// handler
	userHandler := handler.NewUserHandler(userService)
	catHandler := handler.NewCategoryHandler(catService, userService)
	bookHandler := handler.NewBookHandler(bookService, userService)
	basketHandler := handler.NewBasketHandler(basketService, userService)

	return &Dependencies{
		UserRepo:   userRepo,
		CatRepo:    catRepo,
		BookRepo:   bookRepo,
		BasketRepo: basketRepo,

		UserService:   userService,
		CatService:    catService,
		BookService:   bookService,
		BasketService: basketService,

		UserHandler:   userHandler,
		CatHandler:    catHandler,
		BookHandler:   bookHandler,
		BasketHandler: basketHandler,
	}
}
