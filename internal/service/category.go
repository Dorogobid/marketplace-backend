package svc

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/Dorogobid/marketplace-backend/internal/model"
	"github.com/Dorogobid/marketplace-backend/internal/repository"
)

var (
	ErrCategoriesNotFound = errors.New("categories not found")
)

func (s *Service) ListCategories(ctx context.Context) ([]model.Category, error) {
	categories, err := s.repo.ListCategories(ctx)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrCategoriesNotFound
		}

		return nil, fmt.Errorf("service list categories: %w", err)
	}

	return categories, nil
}

func (s *Service) GetParentCategoriesWithCount(ctx context.Context) ([]model.CategoryWithCount, error) {
	categories, err := s.repo.GetParentCategoriesWithCount(ctx)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrCategoriesNotFound
		}

		return nil, fmt.Errorf("service get parent categories with count: %w", err)
	}

	return categories, nil
}

func (s *Service) GetCategoriesWithCountByParentID(ctx context.Context, id string) ([]model.CategoryWithCount, error) {
	categories, err := s.repo.GetCategoriesWithCountByParentID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrCategoriesNotFound
		}

		return nil, fmt.Errorf("service get categories with count by parent id: %w", err)
	}

	return categories, nil
}

func (s *Service) CreateCategory(ctx context.Context, cat model.Category) (model.Category, error) {
	if cat.ID == "" {
		cat.ID = uuid.NewString()
	}

	category, err := s.repo.CreateCategory(ctx, cat)
	if err != nil {
		return model.Category{}, fmt.Errorf("service create category: %w", err)
	}

	return category, nil
}

func (s *Service) UpdateCategory(ctx context.Context, id string, cat model.Category) (model.Category, error) {
	if _, err := s.repo.GetCategoryByID(ctx, id); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return model.Category{}, ErrCategoriesNotFound
		}

		return model.Category{}, fmt.Errorf("service get category by id: %w", err)
	}

	category, err := s.repo.UpdateCategory(ctx, cat)
	if err != nil {
		return model.Category{}, fmt.Errorf("service update category: %w", err)
	}

	return category, nil
}

func (s *Service) DeleteCategory(ctx context.Context, id string) error {
	if _, err := s.repo.GetCategoryByID(ctx, id); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrCategoriesNotFound
		}

		return fmt.Errorf("service get category by id: %w", err)
	}

	if err := s.repo.DeleteCategory(ctx, id); err != nil {
		return fmt.Errorf("service delete category: %w", err)
	}

	return nil
}
