package handler

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/wisphes/book-shop/internal/models"
	"net/http"
	"strconv"
)

type BookHandler struct {
	bookServ BookService
	userServ UserService
}

func NewBookHandler(bookServ BookService, userServ UserService) *BookHandler {
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

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()
	header := r.Header.Get("Authorization")
	userId, err := h.userServ.UserIdentity(ctx, header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if err = h.userServ.IsAdmin(ctx, userId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)["id"]
	bookId, err := strconv.Atoi(vars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = h.bookServ.DeleteBook(ctx, bookId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()
	header := r.Header.Get("Authorization")
	userId, err := h.userServ.UserIdentity(ctx, header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if err = h.userServ.IsAdmin(ctx, userId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var book models.Book
	if err = json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	book, err = h.bookServ.CreateBook(ctx, book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()
	header := r.Header.Get("Authorization")
	userId, err := h.userServ.UserIdentity(ctx, header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if err = h.userServ.IsAdmin(ctx, userId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var book models.Book
	if err = json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	vars := mux.Vars(r)["id"]
	if book.Id, err = strconv.Atoi(vars); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book, err = h.bookServ.UpdateBook(ctx, book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
