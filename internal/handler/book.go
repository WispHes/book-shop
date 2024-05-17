package handler

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/wisphes/book-shop/internal/service"
	"net/http"
	"strconv"
)

type BookHandler struct {
	bookServ *service.BookService
	userServ *service.UserService
}

func NewBookHandler(bookServ *service.BookService, userServ *service.UserService) *BookHandler {
	return &BookHandler{
		bookServ: bookServ,
		userServ: userServ,
	}
}

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()
	books, err := h.bookServ.GetBooks(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()
	vars := mux.Vars(r)["id"]
	id, err := strconv.Atoi(vars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book, err := h.bookServ.GetBook(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
