package service

import (
	"context"
	"github.com/wisphes/book-shop/internal/models"
)

type BasketService struct {
	repo BasketRepository
}

func NewBasketService(repo BasketRepository) *BasketService {
	return &BasketService{repo: repo}
}

func (s *BasketService) GetBasket(ctx context.Context, userId int) (models.Basket, error) {
	return s.repo.GetBasket(ctx, userId)
}

func (s *BasketService) ToPayBasket(ctx context.Context, userId int) error {
	return s.repo.ToPayBasket(ctx, userId)
}

func (s *BasketService) UpdateBasket(ctx context.Context, userId, bookId int, method string) (models.Basket, error) {
	return s.repo.UpdateBasket(ctx, userId, bookId, method)
}
