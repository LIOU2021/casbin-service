package router

import (
	"casbin-service/controller/system"
)

func systemApi() {
	engine.GET("ping", system.Ping)
}
