package middleware

import (
	"casbin-service/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		logger.ErrorfCtx(c, "panic | err: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	})
}
