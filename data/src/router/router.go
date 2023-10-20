package router

import (
	"casbin-service/middleware"

	"github.com/gin-gonic/gin"
)

var (
	engine *gin.Engine
)

func Init() (r *gin.Engine) {
	gin.SetMode(gin.ReleaseMode)
	engine = gin.New()
	engine.Use(
		gin.Recovery(),
		middleware.SystemLogFormatMiddleware,
		// gin.Logger(),
	)

	systemApi()
	return engine
}
