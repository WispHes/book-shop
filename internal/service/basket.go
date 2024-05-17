package service

import (
	"context"
	"github.com/wisphes/book-shop/internal/models"
	"github.com/wisphes/book-shop/internal/repository"
)

type Basket interface {
	GetBasket(ctx context.Context, userId int) (models.Basket, error)
	ToPayBasket(ctx context.Context, userId int) error
	UpdateBasket(ctx context.Context, userId, bookId int, method string) (models.Basket, error)
}

type BasketService struct {
	repo *repository.BasketPostgres
}

func NewBasketService(repo *repository.BasketPostgres) *BasketService {
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
