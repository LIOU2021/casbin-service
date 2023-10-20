package router

import (
	"casbin-service/middleware"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	engine     *gin.Engine
	engineOnce sync.Once
)

func New() (r *gin.Engine) {
	engineOnce.Do(func() {
		gin.SetMode(gin.ReleaseMode)
		engine = gin.New()
		engine.Use(
			gin.Recovery(),
			middleware.SystemLogFormatMiddleware,
			// gin.Logger(),
		)

		systemApi()
	})

	return engine
}
