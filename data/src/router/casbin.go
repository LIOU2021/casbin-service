package router

import "casbin-service/controller/casbin"

func casbinApi() {
	group := engine.Group("/casbin")

	group.POST("/role", casbin.InsertRole)

	group.POST("/policy", casbin.InsertPolicy)
	group.GET("/policy", casbin.GetPolicy)

	group.GET("/model", casbin.GetModel)
}
