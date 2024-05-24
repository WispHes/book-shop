package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/wisphes/book-shop/internal/models"
	"net/http"
)

type UserHandler struct {
	serv UserService
}

func NewUserHandler(serv UserService) *UserHandler {
	return &UserHandler{serv: serv}
}

func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.serv.CreateUser(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(fmt.Sprintf("id:%v", id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := h.serv.GenerateToken(ctx, user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(fmt.Sprintf("token: %v", token)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
