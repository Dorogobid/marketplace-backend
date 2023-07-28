package repository

import (
	"database/sql"

	"github.com/Dorogobid/marketplace-backend/internal/repository/storage"
)

type Repository struct {
	q *storage.Queries
}

func New(db *sql.DB) *Repository {
	return &Repository{q: storage.New(db)}
}
