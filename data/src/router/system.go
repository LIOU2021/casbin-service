package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func systemApi() {
	engine.GET("ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}
