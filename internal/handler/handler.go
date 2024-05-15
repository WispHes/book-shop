package handler

import (
	"github.com/gorilla/mux"
	"net/http"
)

const applicationJson = "application/json"

type Handler struct {
	CatHandler  CategoryHandler
	UserHandler UserHandler
}

func NewHandler(userHand UserHandler, catHand CategoryHandler) *Handler {
	return &Handler{
		CatHandler:  catHand,
		UserHandler: userHand,
	}
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
	r.HandleFunc(`/api/category/{id}`, h.UpdCategory).Methods(http.MethodPut)

	//// Book
	//r.HandleFunc(`/api/books`, h.GetBooks).Methods(http.MethodGet)
	//r.HandleFunc(`/api/book/{id}`, h.GetBook).Methods(http.MethodGet)
	//r.HandleFunc(`/api/book`, h.CreateBook).Methods(http.MethodPost)

	return r
}
