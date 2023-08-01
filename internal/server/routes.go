package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/Dorogobid/marketplace-backend/docs"
)

func (s *Server) setupRoutes() {
	s.e.GET("/docs/*", echoSwagger.WrapHandler)

	apiG := s.e.Group("/api/v1", middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:X-API-KEY",
		Validator: func(key string, c echo.Context) (bool, error) {
			return key == s.xAPIKey, nil
		}}))

	apiG.GET("/category", s.ListCategories)
	apiG.GET("/category/parent", s.GetParentCategoriesWithCount)
	apiG.GET("/category/child", s.GetCategoriesWithCountByParentID)
	apiG.POST("/category", s.CreateCategory)
	apiG.PATCH("/category/:id", s.UpdateCategory)
	apiG.DELETE("/category/:id", s.DeleteCategory)
}
