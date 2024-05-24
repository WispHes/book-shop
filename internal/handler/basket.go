package handler

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/wisphes/book-shop/internal/models"
	"net/http"
	"strconv"
)

type BasketHandler struct {
	basketServ BasketService
	userServ   UserService
}

func NewBasketHandler(basketServ BasketService, userServ UserService) *BasketHandler {
	return &BasketHandler{
		basketServ: basketServ,
		userServ:   userServ,
	}
}

func (h *BasketHandler) GetBasket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()
	header := r.Header.Get("Authorization")

	userId, err := h.userServ.UserIdentity(ctx, header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	basket, err := h.basketServ.GetBasket(ctx, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(basket.Books); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *BasketHandler) ToPayBasket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()
	header := r.Header.Get("Authorization")

	userId, err := h.userServ.UserIdentity(ctx, header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if err = h.basketServ.ToPayBasket(ctx, userId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *BasketHandler) UpdateBasket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()
	method := r.Method
	header := r.Header.Get("Authorization")
	userId, err := h.userServ.UserIdentity(ctx, header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)["id"]
	bookId, err := strconv.Atoi(vars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var basket models.Basket
	if basket, err = h.basketServ.UpdateBasket(ctx, userId, bookId, method); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(basket.Books); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
