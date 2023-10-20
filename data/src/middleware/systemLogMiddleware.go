package middleware

import (
	"casbin-service/logger"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SystemLogFormatMiddleware(c *gin.Context) {
	start := time.Now()
	slow := false

	c.Next()
	d := time.Since(start)
	if d > (time.Millisecond * 250) {
		slow = true
	}

	logger.Info("api access",
		zap.Int("status", c.Writer.Status()),
		zap.String("ip", c.ClientIP()),
		zap.String("ts", start.Format(time.RFC3339Nano)),
		zap.String("proto", c.Request.Proto),
		zap.String("method", c.Request.Method),
		zap.String("url", c.Request.URL.Path+"?"+c.Request.URL.Query().Encode()),
		zap.String("latency", d.String()),
		zap.String("ua", c.Request.UserAgent()),
		zap.Error(c.Err()),
		zap.Bool("slow", slow),
	)
}
