package mock_router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/millirud/intermocker/internal/di"
	"github.com/millirud/intermocker/internal/entity"
	"github.com/millirud/intermocker/internal/service/http_routers_factory"
)

func New(handler *gin.Engine, Di *di.DI) {
	service := http_routers_factory.New(Di.Logger, handler)

	route := entity.NewRouteSignatureBuilder(http.MethodGet, "/test/api").New()

	service.AddRoute(route)
}
