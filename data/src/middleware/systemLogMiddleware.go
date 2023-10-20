package middleware

import (
	"casbin-service/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func SystemLogFormatMiddleware(c *gin.Context) {
	start := time.Now()

	c.Next()

	d := time.Since(start)

	logger.GetAccessLogger().Sugar().Infof("%s - [%s] \"%s %s %s\" %d %d %s \"%s\"",
		c.ClientIP(),
		time.Now().Format(time.RFC3339Nano),
		c.Request.Method,
		c.Request.URL.String(),
		c.Request.Proto,
		c.Writer.Status(),
		c.Request.ContentLength,
		d.String(),
		c.Request.UserAgent(),
	)
}
