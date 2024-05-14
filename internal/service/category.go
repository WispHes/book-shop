package service

import (
	"context"
	"github.com/wisphes/book-shop/internal/models"
	"github.com/wisphes/book-shop/internal/repository"
)

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetCategories(ctx context.Context) ([]models.Category, error) {
	return s.repo.GetCategories(ctx)
}

func (s *CategoryService) GetCategory(ctx context.Context, id int) (models.Category, error) {
	return s.repo.GetCategory(ctx, id)
}

func (s *CategoryService) CreateCategory(ctx context.Context, category models.Category) (models.Category, error) {
	return s.repo.CreateCategory(ctx, category)
}

func (s *CategoryService) UpdateCategory(ctx context.Context, category models.Category) (models.Category, error) {
	return s.repo.UpdateCategory(ctx, category)
}
