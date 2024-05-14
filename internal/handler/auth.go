package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/wisphes/book-shop/internal/models"
	"net/http"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", applicationJson)
	ctx := context.Background()
	var input models.User

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.services.Authorization.CreateUser(ctx, input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(fmt.Sprintf("id:%v", id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", applicationJson)
	ctx := context.Background()
	var input signInInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := h.services.Authorization.GenerateToken(ctx, input.Email, input.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(fmt.Sprintf("token: %v", token)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
