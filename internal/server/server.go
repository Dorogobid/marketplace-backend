package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	svc "github.com/Dorogobid/marketplace-backend/internal/service"
)

type Server struct {
	e       *echo.Echo
	svc     *svc.Service
	xAPIKey string
	baseURL string
}

func NewServer(svc *svc.Service, key, url string) *Server {
	s := &Server{
		e:       echo.New(),
		svc:     svc,
		xAPIKey: key,
		baseURL: url,
	}

	s.setupServer()
	return s
}

func (s *Server) setupServer() {
	s.e.Use(middleware.Recover())
	s.e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time": "${time_rfc3339}", "method": "${method}", "uri": "${uri}", "status": "${status}", "remote_ip": "${remote_ip}"}` + "\n"}))
	s.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}}))
	s.e.Static("/static", "static")
	s.setupRoutes()
}

func (s *Server) Logger() echo.Logger {
	return s.e.Logger
}

func (s *Server) Start(address string) error {
	return s.e.Start(address)
}
