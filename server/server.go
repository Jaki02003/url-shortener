package server

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"go-redis-url-shortener/config"
	"go-redis-url-shortener/utils/log"
)

type Server struct {
	framework *echo.Echo
	config    *config.Config
}

func New(config *config.Config, framework *echo.Echo) *Server {
	return &Server{
		framework: framework,
		config:    config,
	}
}

func (s *Server) Start() {
	e := s.framework
	// start routes server
	go func() {
		e.Logger.Fatal(e.Start(":" + s.config.App.Port))
	}()
	// graceful shutdown
	s.GracefulShutdown()
}

// GracefulShutdown server will gracefully shut down within 5 sec
func (s *Server) GracefulShutdown() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	log.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ = s.framework.Shutdown(ctx)
	log.Info("server shutdowns gracefully")
}
