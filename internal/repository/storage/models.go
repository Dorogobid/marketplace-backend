// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package storage

import (
	"github.com/google/uuid"
)

type Category struct {
	ID           uuid.UUID
	ChildOf      uuid.NullUUID
	CategoryName string
	ImageUrl     string
	IsActive     bool
	SortIndex    int64
}
