package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"

	svc "github.com/Dorogobid/marketplace-backend/internal/service"
)

type Server struct {
	e       *echo.Echo
	svc     *svc.Service
	xAPIKey string
}

func NewServer(svc *svc.Service, key string) *Server {
	s := &Server{
		e:       echo.New(),
		svc:     svc,
		xAPIKey: key,
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
	s.e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")

	s.setupRoutes()
}

func (s *Server) Logger() echo.Logger {
	return s.e.Logger
}

func (s *Server) Start(address string) error {
	return s.e.StartAutoTLS(address)
}
