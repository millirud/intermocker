package system_routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewLivenessProbe() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}
