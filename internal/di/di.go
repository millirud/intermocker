package di

import (
	"os"

	"github.com/millirud/intermocker/config"
	"github.com/millirud/intermocker/pkg/logger"
	"github.com/rs/zerolog"
)

type DI struct {
	Logger *zerolog.Logger
}

func New(cfg *config.Config) *DI {
	logger := logger.New(cfg.Log.Level, os.Stdout)

	return &DI{
		Logger: logger,
	}
}
