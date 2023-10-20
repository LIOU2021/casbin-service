package router

import (
	"casbin-service/logger"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func systemApi() {
	if engine == nil {
		logger.Error("engine was nil")
		os.Exit(1)
	}
	engine.GET("ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}
