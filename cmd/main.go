package main

import (
	_ "github.com/lib/pq"
	bookshop "github.com/wisphes/book-shop"
	"github.com/wisphes/book-shop/configs"
	"github.com/wisphes/book-shop/internal/pg"
	"github.com/wisphes/book-shop/internal/pkg/handler"
	"github.com/wisphes/book-shop/internal/pkg/repository"
	"github.com/wisphes/book-shop/internal/pkg/service"
	"log"
)

func main() {
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

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	r := handler.NewHandler(services)

	srv := &bookshop.Server{}
	log.Fatal(srv.Run(cfg.ServerPort, r.InitRoutes()))

}
