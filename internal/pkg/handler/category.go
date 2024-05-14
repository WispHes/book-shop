package handler

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/wisphes/book-shop/internal/models"
	"net/http"
	"strconv"
)

func (h *Handler) GetCategories(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", applicationJson)

	categories, err := h.services.Category.GetCategories(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(categories); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(`Content-Type`, applicationJson)
	ctx := context.Background()

	vars := mux.Vars(r)["id"]
	id, err := strconv.Atoi(vars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	category, err := h.services.Category.GetCategory(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(category); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(`Content-Type`, applicationJson)
	ctx := context.Background()

	id, err := h.UserIdentity(w, r)
	if id == 0 {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if _, err = h.services.Authorization.IsAdmin(ctx, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var category models.Category
	if err = json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newCategory, err := h.services.Category.CreateCategory(ctx, category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(newCategory); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(`Content-Type`, applicationJson)
	ctx := context.Background()

	userId, err := h.UserIdentity(w, r)
	if userId == 0 {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if _, err = h.services.Authorization.IsAdmin(ctx, userId); err != nil {
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newCategory, err := h.services.Category.UpdateCategory(ctx, updCat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(newCategory); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
