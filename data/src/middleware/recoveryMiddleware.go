package middleware

import (
	"casbin-service/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		logger.ErrorName("panic", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	})
}
