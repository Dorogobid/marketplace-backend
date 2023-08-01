package repository

import (
	"context"
	"database/sql"

	"github.com/Dorogobid/marketplace-backend/internal/model"
	"github.com/Dorogobid/marketplace-backend/internal/repository/storage"
)

type Repository interface {
	ListCategories(ctx context.Context) ([]model.Category, error)
	GetCategoryByID(ctx context.Context, id string) (model.Category, error)
	GetParentCategoriesWithCount(ctx context.Context) ([]model.CategoryWithCount, error)
	GetCategoriesWithCountByParentID(ctx context.Context, id string) ([]model.CategoryWithCount, error)
	CreateCategory(ctx context.Context, c model.Category) (model.Category, error)
	UpdateCategory(ctx context.Context, c model.Category) (model.Category, error)
	DeleteCategory(ctx context.Context, id string) error
}

func New(db *sql.DB) Repository {
	return &repository{q: storage.New(db)}
}

type repository struct {
	q *storage.Queries
}
