package model

import (
	"errors"

	"github.com/google/uuid"
)

type Category struct {
	ID           string `json:"id"`
	ChildOf      string `json:"child_of"`
	CategoryName string `json:"category_name"`
	ImageUrl     string `json:"image_url"`
	IsActive     bool   `json:"is_active"`
	SortIndex    int64  `json:"sort_index"`
}

func (c *Category) Validate() error {
	if c.ChildOf != "" {
		if _, err := uuid.Parse(c.ChildOf); err != nil {
			return errors.New("child_of is not valid uuid")
		}
	}

	if c.CategoryName == "" {
		return errors.New("category_name is required")
	}

	return nil
}

type CategoryWithCount struct {
	ID           string `json:"id"`
	ChildOf      string `json:"child_of"`
	CategoryName string `json:"category_name"`
	ImageUrl     string `json:"image_url"`
	IsActive     bool   `json:"is_active"`
	SortIndex    int64  `json:"sort_index"`
	ChildCount   int64  `json:"child_count"`
}
