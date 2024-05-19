package main

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/wisphes/book-shop/cmd/dep"
	"github.com/wisphes/book-shop/configs"
	"github.com/wisphes/book-shop/internal/pkg/pg"
	"log"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, routes *mux.Router) error {

	s.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: routes,
	}
	return s.httpServer.ListenAndServe()
}

func buildRoutes(r *mux.Router, dep *dep.Dependencies) *mux.Router {
	// Authorization
	r.HandleFunc(`/auth/sign-up`, dep.UserHandler.SignUp).Methods(http.MethodPost)
	r.HandleFunc(`/auth/sign-in`, dep.UserHandler.SignIn).Methods(http.MethodPost)

	// Category
	r.HandleFunc(`/api/categories`, dep.CatHandler.GetCategories).Methods(http.MethodGet)
	r.HandleFunc(`/api/category`, dep.CatHandler.CreateCategory).Methods(http.MethodPost)
	r.HandleFunc(`/api/category/{id}`, dep.CatHandler.GetCategory).Methods(http.MethodGet)
	r.HandleFunc(`/api/category/{id}`, dep.CatHandler.UpdateCategory).Methods(http.MethodPut)
	r.HandleFunc(`/api/category/{id}`, dep.CatHandler.DeleteCategory).Methods(http.MethodDelete)

	// Book
	r.HandleFunc(`/api/books`, dep.BookHandler.GetBooks).Methods(http.MethodGet)
	r.HandleFunc(`/api/book`, dep.BookHandler.CreateBook).Methods(http.MethodPost)
	r.HandleFunc(`/api/book/{id}`, dep.BookHandler.GetBook).Methods(http.MethodGet)
	r.HandleFunc(`/api/book/{id}`, dep.BookHandler.UpdateBook).Methods(http.MethodPut)
	r.HandleFunc(`/api/book/{id}`, dep.BookHandler.DeleteBook).Methods(http.MethodDelete)

	// Basket
	r.HandleFunc(`/api/basket`, dep.BasketHandler.GetBasket).Methods(http.MethodGet)
	r.HandleFunc(`/api/basket/{id}`, dep.BasketHandler.UpdateBasket).Methods(http.MethodDelete, http.MethodPut)
	r.HandleFunc(`/api/basket/pay`, dep.BasketHandler.ToPayBasket).Methods(http.MethodPost)

	return r
}

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

	dependencies := dep.NewDependencies(db)
	r := buildRoutes(mux.NewRouter(), dependencies)
	srv := &Server{}

	log.Fatal(srv.Run(cfg.ServerPort, r))
}
