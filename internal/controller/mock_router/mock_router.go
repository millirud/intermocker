package mock_router

import (
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/millirud/intermocker/internal/di"
	"github.com/millirud/intermocker/internal/entity"
	"github.com/millirud/intermocker/internal/service/http_routers_factory"
)

func New(handler *gin.Engine, Di *di.DI) {
	service := http_routers_factory.New(Di.Logger, handler)

	sign := entity.NewRouteSignatureBuilder(http.MethodGet, "/test/api").New()

	response := entity.NewHttpResponseBuilder(http.StatusBadGateway, "application-json", func() (io.Reader, int64, string, error) {
		data := `{"status": "ok"}`
		return strings.NewReader(data), int64(len(data)), "application/json", nil
	}).New()

	service.AddRoute(sign, response)
}
