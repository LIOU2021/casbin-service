package router

import (
	"casbin-service/middleware"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

var (
	engine *gin.Engine
)

func Init() (r *gin.Engine) {
	gin.SetMode(gin.ReleaseMode)
	engine = gin.New()
	engine.Use(
		requestid.New(),
		middleware.SystemLogFormatMiddleware,
		middleware.RecoveryMiddleware(),
	)

	systemApi()
	return engine
}
