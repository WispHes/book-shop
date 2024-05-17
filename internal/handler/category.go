package handler

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/wisphes/book-shop/internal/models"
	"github.com/wisphes/book-shop/internal/service"
	"net/http"
	"strconv"
)

type CategoryHandler struct {
	catServ  *service.CategoryService
	userServ *service.UserService
}

func NewCategoryHandler(catServ *service.CategoryService, userServ *service.UserService) *CategoryHandler {
	return &CategoryHandler{
		catServ:  catServ,
		userServ: userServ,
	}
}

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()
	categories, err := h.catServ.GetCategories(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(categories); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(`Content-Type`, "application/json")

	ctx := context.Background()
	vars := mux.Vars(r)["id"]
	catId, err := strconv.Atoi(vars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	category, err := h.catServ.GetCategory(ctx, catId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(category); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(`Content-Type`, "application/json")

	ctx := context.Background()
	header := r.Header.Get("Authorization")
	userId, err := h.userServ.UserIdentity(ctx, header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if _, err = h.userServ.IsAdmin(ctx, userId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var category models.Category
	if err = json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newCategory, err := h.catServ.CreateCategory(ctx, category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(newCategory); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(`Content-Type`, "application/json")

	ctx := context.Background()
	header := r.Header.Get("Authorization")
	userId, err := h.userServ.UserIdentity(ctx, header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if _, err = h.userServ.IsAdmin(ctx, userId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var updCat models.Category

	if err = json.NewDecoder(r.Body).Decode(&updCat); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)["id"]
	updCat.Id, err = strconv.Atoi(vars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newCategory, err := h.catServ.UpdateCategory(ctx, updCat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(newCategory); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(`Content-Type`, "application/json")

	ctx := context.Background()
	header := r.Header.Get("Authorization")

	userId, err := h.userServ.UserIdentity(ctx, header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	if _, err = h.userServ.IsAdmin(ctx, userId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	vars := mux.Vars(r)["id"]
	catId, err := strconv.Atoi(vars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if err = h.catServ.DeleteCategory(ctx, catId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)

}
