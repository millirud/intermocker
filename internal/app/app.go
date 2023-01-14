// Package app configures and runs application.
package app

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/millirud/intermocker/config"
	"github.com/millirud/intermocker/internal/controller/mock_router"
	system_routes "github.com/millirud/intermocker/internal/controller/system"
	"github.com/millirud/intermocker/internal/di"
	"github.com/millirud/intermocker/pkg/httpserver"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	var err error

	Di := di.New(cfg)
	l := Di.Logger

	// HTTP Server
	handler := gin.Default()

	handler.GET("/_mocker/healthz", system_routes.NewLivenessProbe())

	mock_router.New(handler, Di)

	httpServer := httpserver.New(
		handler,
		httpserver.Port(cfg.HTTP.Port),
		httpserver.ReadTimeout(time.Duration(cfg.HTTP.ReadTimeout)*time.Millisecond),
		httpserver.WriteTimeout(time.Duration(cfg.HTTP.WriteTimeout)*time.Millisecond),
	)

	l.Info().Msgf("server started on port %s", cfg.HTTP.Port)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info().Msgf("app - Run - signal: %s", s.String())
	case err = <-httpServer.Notify():
		l.Error().Msgf("app - Run - httpServer.Notify: %w", err)
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error().Msgf("app - Run - httpServer.Shutdown: %w", err)
	}
}
