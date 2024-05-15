package dep

import (
	"github.com/wisphes/book-shop/configs"
	"github.com/wisphes/book-shop/internal/handler"
	"github.com/wisphes/book-shop/internal/pkg/pg"
	"github.com/wisphes/book-shop/internal/repository"
	"github.com/wisphes/book-shop/internal/service"
	"log"
)

func InitDep() *handler.Handler {
	cfg, err := configs.NewParsedConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := pg.NewPostgresDB(pg.Config{
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		User:     cfg.Database.User,
		DBName:   cfg.Database.DBName,
		Password: cfg.Database.Password,
		SSLMode:  cfg.Database.SSLMode,
	})
	if err != nil {
		log.Fatal(err)
	}

	repUser := repository.NewAuthPostgres(db)
	servUser := service.NewAuthService(repUser)
	handUser := handler.NewUserHandler(servUser)

	repCat := repository.NewCategoryPostgres(db)
	servCat := service.NewCategoryService(repCat)
	handCat := handler.NewCategoryHandler(servCat)

	return handler.NewHandler(*handUser, *handCat)
}
