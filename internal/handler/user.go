package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/wisphes/book-shop/internal/models"
	"github.com/wisphes/book-shop/internal/service"
	"net/http"
)

type UserHandler struct {
	serv *service.UserService
}

func NewUserHandler(serv *service.UserService) *UserHandler {
	return &UserHandler{serv: serv}
}

func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	var input models.User

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.serv.CreateUser(ctx, input)
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

func (h *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	var input signInInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := h.serv.GenerateToken(ctx, input.Email, input.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(fmt.Sprintf("token: %v", token)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
