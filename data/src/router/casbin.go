package router

import "casbin-service/controller/casbin"

func casbinApi() {
	group := engine.Group("/casbin")

	group.POST("/role", casbin.InsertRole)
	group.GET("/role", casbin.GetRole)
	group.POST("/role/find", casbin.GetRoleForUser)

	group.POST("/policy", casbin.InsertPolicy)
	group.GET("/policy", casbin.GetPolicy)

	group.GET("/model", casbin.GetModel)

	group.POST("/enforce", casbin.Enforce)

	group.GET("/info", casbin.GetInfo)
}
