package router

import (
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
		r = gin.New()
		r.Use(
			gin.Recovery(),
			// gin.Logger(),
		)

		systemApi()
	})

	return engine
}
