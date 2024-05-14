package handler

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) UserIdentity(w http.ResponseWriter, r *http.Request) (int, error) {
	ctx := context.Background()
	header := r.Header.Get(authorizationHeader)
	if header == "" {
		return 0, errors.New("invalid token")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return 0, errors.New("invalid token")
	}
	userId, err := h.services.Authorization.ParseToken(ctx, headerParts[1])
	if err != nil {
		return 0, errors.New("invalid token")
	}
	return userId, err
}
