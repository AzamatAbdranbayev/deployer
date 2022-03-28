package http

import (
	"context"

	configs "github.com/AzamatAbdranbayev/deployer/internal/config"
	"github.com/AzamatAbdranbayev/deployer/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	cfg     *configs.Config
	handler *handler.Manager
}

func NewServer(cfg *configs.Config, handler *handler.Manager) *Server {
	return &Server{cfg: cfg, handler: handler}
}

func (h *Server) StartHTTPServer(ctx context.Context) error {
	e := h.EchoEngine()
	return e.Start(h.cfg.Server.Addr)
}

func (h *Server) EchoEngine() *echo.Echo {

	app := echo.New()

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	//
	//insService, _ := service2.NewManager(h.cfg, service2.CreateHttpClient())
	//insHandler := handler.NewManager(insService)

	app.POST("/deploy/react", h.handler.Static.DeployStatic)
	return app
}
