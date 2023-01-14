package http_routers_factory

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func InitRoutes(logger *zerolog.Logger, handler *gin.Engine) {
	service := New(logger, handler)

	service.Init()
}

func New(logger *zerolog.Logger, handler *gin.Engine) *HttpRoutersFactory {
	return &HttpRoutersFactory{
		logger:  logger,
		handler: handler,
	}
}

type HttpRoutersFactory struct {
	logger  *zerolog.Logger
	handler *gin.Engine
}

func (r *HttpRoutersFactory) Init() {

}
