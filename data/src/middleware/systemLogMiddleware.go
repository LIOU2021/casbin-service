package middleware

import (
	"casbin-service/logger"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func SystemLogFormatMiddleware(c *gin.Context) {
	gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		accessLog := fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			c.ClientIP(),
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)

		if param.Latency > (time.Millisecond * 250) {
			logger.WarnName("api slow query", accessLog)
		}
		return accessLog
	})(c)
}
