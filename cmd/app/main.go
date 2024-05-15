package main

import (
	_ "github.com/lib/pq"
	bookshop "github.com/wisphes/book-shop"
	"github.com/wisphes/book-shop/cmd/dep"
	"github.com/wisphes/book-shop/configs"
	"log"
)

func main() {
	cfg, err := configs.NewParsedConfig()
	if err != nil {
		log.Fatal(err)
	}

	srv := &bookshop.Server{}
	r := dep.InitDep()
	log.Fatal(srv.Run(cfg.ServerPort, r.InitRoutes()))

}
