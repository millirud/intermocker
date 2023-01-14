package http_routers_factory

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/millirud/intermocker/internal/entity"
	"github.com/rs/zerolog"
)

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

func (r *HttpRoutersFactory) AddRoute(sig *entity.RouteSignature) {
	r.logger.Info().Msgf("creating route %s %s", sig.Method, sig.Path)

	r.handler.Handle(sig.Method, sig.Path, func(c *gin.Context) {
		c.String(http.StatusOK, "it works")
	})
}

// func (r *HttpRoutersFactory) useFuncByMethod(method string) gin.IRoutes {
// 	switch method {
// 	case http.MethodGet:
// 		return r.handler.GET

// 	default:
// 		return r.handler.GET
// 	}
// }

func (r *HttpRoutersFactory) Init() {

}
