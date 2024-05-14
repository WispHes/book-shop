package handler

import (
	"github.com/gorilla/mux"
	"github.com/wisphes/book-shop/internal/service"
	"net/http"
)

const applicationJson = "application/json"

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	// Authorization
	r.HandleFunc(`/auth/sign-up`, h.signUp).Methods(http.MethodPost)
	r.HandleFunc(`/auth/sign-in`, h.signIn).Methods(http.MethodPost)

	// Category
	r.HandleFunc(`/api/categories`, h.GetCategories).Methods(http.MethodGet)
	r.HandleFunc(`/api/category/{id}`, h.GetCategory).Methods(http.MethodGet)
	r.HandleFunc(`/api/category`, h.CreateCategory).Methods(http.MethodPost)
	r.HandleFunc(`/api/category/{id}`, h.UpdateCategory).Methods(http.MethodPut)

	//// Book
	//r.HandleFunc(`/api/books`, h.GetBooks).Methods(http.MethodGet)
	//r.HandleFunc(`/api/book/{id}`, h.GetBook).Methods(http.MethodGet)
	//r.HandleFunc(`/api/book`, h.CreateBook).Methods(http.MethodPost)

	return r
}
