package service

import (
	"context"
	"github.com/wisphes/book-shop/internal/models"
)

type CategoryService struct {
	repo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetCategories(ctx context.Context) ([]models.Category, error) {
	return s.repo.GetCategories(ctx)
}

func (s *CategoryService) GetCategory(ctx context.Context, catId int) (models.Category, error) {
	return s.repo.GetCategory(ctx, catId)
}

func (s *CategoryService) CreateCategory(ctx context.Context, category models.Category) (models.Category, error) {
	return s.repo.CreateCategory(ctx, category)
}

func (s *CategoryService) UpdateCategory(ctx context.Context, category models.Category) (models.Category, error) {
	return s.repo.UpdateCategory(ctx, category)
}

func (s *CategoryService) DeleteCategory(ctx context.Context, catId int) error {
	return s.repo.DeleteCategory(ctx, catId)
}
