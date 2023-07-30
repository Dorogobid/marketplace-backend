package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/Dorogobid/marketplace-backend/internal/model"
	svc "github.com/Dorogobid/marketplace-backend/internal/service"
)

// ListCategories godoc
//
//	@Summary		List all categories
//	@Description	get all categories
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]model.Category
//	@Failure		401	{object}	echo.HTTPError
//	@Failure		404	{object}	echo.HTTPError
//	@Failure		500	{object}	echo.HTTPError
//	@Security		ApiKeyAuth
//	@Router			/api/v1/category [get]
func (s *Server) ListCategories(c echo.Context) error {
	categories, err := s.svc.ListCategories(c.Request().Context())
	if err != nil {
		if errors.Is(err, svc.ErrCategoriesNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, categories)
}

// GetParentCategoriesWithCount godoc
//
//	@Summary		Get all active parent categories with child count
//	@Description	get all active parent categories with child count
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]model.CategoryWithCount
//	@Failure		401	{object}	echo.HTTPError
//	@Failure		404	{object}	echo.HTTPError
//	@Failure		500	{object}	echo.HTTPError
//	@Security		ApiKeyAuth
//	@Router			/api/v1/category/parent [get]
func (s *Server) GetParentCategoriesWithCount(c echo.Context) error {
	categories, err := s.svc.GetParentCategoriesWithCount(c.Request().Context())
	if err != nil {
		if errors.Is(err, svc.ErrCategoriesNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, categories)
}

// GetCategoriesWithCountByParentID godoc
//
//	@Summary		Get all active categories with child count by parent id
//	@Description	get all active categories with child count by parent id
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			parent_id	query		string	true	"parent category id"
//	@Success		200			{object}	[]model.CategoryWithCount
//	@Failure		400			{object}	echo.HTTPError
//	@Failure		401			{object}	echo.HTTPError
//	@Failure		404			{object}	echo.HTTPError
//	@Failure		500			{object}	echo.HTTPError
//	@Security		ApiKeyAuth
//	@Router			/api/v1/category/child [get]
func (s *Server) GetCategoriesWithCountByParentID(c echo.Context) error {
	var id string
	if err := echo.QueryParamsBinder(c).String("parent_id", &id).BindErrors(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if _, err := uuid.Parse(id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("parse id: %w", err))
	}

	categories, err := s.svc.GetCategoriesWithCountByParentID(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, svc.ErrCategoriesNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, categories)
}

// CreateCategory godoc
//
//	@Summary		Create a new category
//	@Description	create a new category
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.Category	true	"category to create"
//	@Success		200		{object}	[]model.Category
//	@Failure		400		{object}	echo.HTTPError
//	@Failure		401		{object}	echo.HTTPError
//	@Failure		404		{object}	echo.HTTPError
//	@Failure		500		{object}	echo.HTTPError
//	@Security		ApiKeyAuth
//	@Router			/api/v1/category [post]
func (s *Server) CreateCategory(c echo.Context) error {
	cat := new(model.Category)
	if err := c.Bind(cat); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := cat.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	category, err := s.svc.CreateCategory(c.Request().Context(), *cat)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, category)
}

// UpdateCategory godoc
//
//	@Summary		Update an exist category
//	@Description	Update an exist category
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string			true	"category id"
//	@Param			request	body		model.Category	true	"category to update"
//	@Success		200		{object}	[]model.Category
//	@Failure		400		{object}	echo.HTTPError
//	@Failure		401		{object}	echo.HTTPError
//	@Failure		404		{object}	echo.HTTPError
//	@Failure		500		{object}	echo.HTTPError
//	@Security		ApiKeyAuth
//	@Router			/api/v1/category/:id [patch]
func (s *Server) UpdateCategory(c echo.Context) error {
	var id string
	if err := echo.PathParamsBinder(c).String("id", &id).BindErrors(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	cat := new(model.Category)
	if err := c.Bind(cat); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := cat.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if _, err := uuid.Parse(id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("parse id: %w", err))
	}

	category, err := s.svc.UpdateCategory(c.Request().Context(), id, *cat)
	if err != nil {
		if errors.Is(err, svc.ErrCategoriesNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, category)
}

// DeleteCategory godoc
//
//	@Summary		Delete an exist category
//	@Description	Delete an exist category
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"category id"
//	@Success		200	{object}	echo.HTTPError
//	@Failure		400	{object}	echo.HTTPError
//	@Failure		401	{object}	echo.HTTPError
//	@Failure		404	{object}	echo.HTTPError
//	@Failure		500	{object}	echo.HTTPError
//	@Security		ApiKeyAuth
//	@Router			/api/v1/category/:id [delete]
func (s *Server) DeleteCategory(c echo.Context) error {
	var id string
	if err := echo.PathParamsBinder(c).String("id", &id).BindErrors(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if _, err := uuid.Parse(id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("parse id: %w", err))
	}

	if err := s.svc.DeleteCategory(c.Request().Context(), id); err != nil {
		if errors.Is(err, svc.ErrCategoriesNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "OK")
}
