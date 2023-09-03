package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// UploadFile godoc
//
//	@Summary		Upload static file
//	@Description	Upload static file (pictures)
//	@Tags			Static
//	@Produce		json
//	@Param			file	formData	file	true	"Static file"
//	@Success		200		{object}	echo.HTTPError
//	@Failure		401		{object}	echo.HTTPError
//	@Failure		404		{object}	echo.HTTPError
//	@Failure		500		{object}	echo.HTTPError
//	@Security		ApiKeyAuth
//	@Router			/api/v1/upload [POST]
func (s *Server) UploadFile(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	path, err := s.svc.UploadFile(file, s.baseURL)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, path)
}
