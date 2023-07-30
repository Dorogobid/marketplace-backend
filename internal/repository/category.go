package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/Dorogobid/marketplace-backend/internal/model"
	"github.com/Dorogobid/marketplace-backend/internal/repository/storage"
)

func (r *repository) ListCategories(ctx context.Context) ([]model.Category, error) {
	dbCategories, err := r.q.ListCategories(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("repository list categories: %w", err)
	}

	if len(dbCategories) == 0 {
		return nil, ErrNotFound
	}

	categories := make([]model.Category, len(dbCategories))
	for i, c := range dbCategories {
		childV := ""
		if c.ChildOf.Valid {
			childV = c.ChildOf.UUID.String()
		}

		categories[i] = model.Category{
			ID:           c.ID.String(),
			ChildOf:      childV,
			CategoryName: c.CategoryName,
			ImageUrl:     c.ImageUrl,
			IsActive:     c.IsActive,
			SortIndex:    c.SortIndex,
		}
	}

	return categories, err
}

func (r *repository) GetCategoryByID(ctx context.Context, id string) (model.Category, error) {
	c, err := r.q.GetCategoryByID(ctx, uuid.MustParse(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Category{}, ErrNotFound
		}

		return model.Category{}, fmt.Errorf("repository get category by id: %w", err)
	}

	childV := ""
	if c.ChildOf.Valid {
		childV = c.ChildOf.UUID.String()
	}

	return model.Category{
		ID:           c.ID.String(),
		ChildOf:      childV,
		CategoryName: c.CategoryName,
		ImageUrl:     c.ImageUrl,
		IsActive:     c.IsActive,
		SortIndex:    c.SortIndex,
	}, nil
}

func (r *repository) GetParentCategoriesWithCount(ctx context.Context) ([]model.CategoryWithCount, error) {
	dbCategories, err := r.q.GetParentCategoriesWithCount(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("repository get parent categories with count: %w", err)
	}

	if len(dbCategories) == 0 {
		return nil, ErrNotFound
	}

	categories := make([]model.CategoryWithCount, len(dbCategories))
	for i, c := range dbCategories {
		childV := ""
		if c.ChildOf.Valid {
			childV = c.ChildOf.UUID.String()
		}

		categories[i] = model.CategoryWithCount{
			ID:           c.ID.String(),
			ChildOf:      childV,
			CategoryName: c.CategoryName,
			ImageUrl:     c.ImageUrl,
			IsActive:     c.IsActive,
			SortIndex:    c.SortIndex,
			ChildCount:   c.ChildCount,
		}
	}

	return categories, err
}

func (r *repository) GetCategoriesWithCountByParentID(ctx context.Context, id string) ([]model.CategoryWithCount, error) {
	parentID := uuid.NullUUID{
		UUID:  uuid.MustParse(id),
		Valid: true,
	}
	dbCategories, err := r.q.GetCategoriesWithCountByParentID(ctx, parentID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("repository get categories with count by parent id: %w", err)
	}

	if len(dbCategories) == 0 {
		return nil, ErrNotFound
	}

	categories := make([]model.CategoryWithCount, len(dbCategories))
	for i, c := range dbCategories {
		childV := ""
		if c.ChildOf.Valid {
			childV = c.ChildOf.UUID.String()
		}

		categories[i] = model.CategoryWithCount{
			ID:           c.ID.String(),
			ChildOf:      childV,
			CategoryName: c.CategoryName,
			ImageUrl:     c.ImageUrl,
			IsActive:     c.IsActive,
			SortIndex:    c.SortIndex,
			ChildCount:   c.ChildCount,
		}
	}

	return categories, err
}

func (r *repository) CreateCategory(ctx context.Context, c model.Category) (model.Category, error) {
	childOf := uuid.NullUUID{}
	if c.ChildOf != "" {
		childOf.Valid = true
		childOf.UUID = uuid.MustParse(c.ChildOf)
	}

	dbCategory, err := r.q.CreateCategory(ctx, storage.CreateCategoryParams{
		ID:           uuid.MustParse(c.ID),
		ChildOf:      childOf,
		CategoryName: c.CategoryName,
		ImageUrl:     c.ImageUrl,
		IsActive:     c.IsActive,
		SortIndex:    c.SortIndex,
	})
	if err != nil {
		return model.Category{}, fmt.Errorf("repository create category: %w", err)
	}

	child := ""
	if dbCategory.ChildOf.Valid {
		child = dbCategory.ChildOf.UUID.String()
	}

	return model.Category{
		ID:           dbCategory.ID.String(),
		ChildOf:      child,
		CategoryName: dbCategory.CategoryName,
		ImageUrl:     dbCategory.ImageUrl,
		IsActive:     dbCategory.IsActive,
		SortIndex:    dbCategory.SortIndex,
	}, nil
}

func (r *repository) UpdateCategory(ctx context.Context, c model.Category) (model.Category, error) {
	childOf := uuid.NullUUID{}
	if c.ChildOf != "" {
		childOf.Valid = true
		childOf.UUID = uuid.MustParse(c.ChildOf)
	}

	dbCategory, err := r.q.UpdateCategory(ctx, storage.UpdateCategoryParams{
		ID:           uuid.MustParse(c.ID),
		ChildOf:      childOf,
		CategoryName: c.CategoryName,
		ImageUrl:     c.ImageUrl,
		IsActive:     c.IsActive,
		SortIndex:    c.SortIndex,
	})
	if err != nil {
		return model.Category{}, fmt.Errorf("repository update category: %w", err)
	}

	child := ""
	if dbCategory.ChildOf.Valid {
		child = dbCategory.ChildOf.UUID.String()
	}

	return model.Category{
		ID:           dbCategory.ID.String(),
		ChildOf:      child,
		CategoryName: dbCategory.CategoryName,
		ImageUrl:     dbCategory.ImageUrl,
		IsActive:     dbCategory.IsActive,
		SortIndex:    dbCategory.SortIndex,
	}, nil
}

func (r *repository) DeleteCategory(ctx context.Context, id string) error {
	if err := r.q.DeleteCategory(ctx, uuid.MustParse(id)); err != nil {
		return fmt.Errorf("repository delete category: %w", err)
	}

	return nil
}
