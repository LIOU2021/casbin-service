package router

import "casbin-service/controller/casbin"

func casbinApi() {
	group := engine.Group("/casbin")

	group.POST("/role", casbin.AddGroupingPolicy)
	group.GET("/role", casbin.GetRole)
	group.DELETE("/role", casbin.DeleteRole)
	group.DELETE("/role/user", casbin.DeleteRoleForUser)
	group.POST("/role/find", casbin.GetRoleForUser)

	group.DELETE("/user", casbin.DeleteUser)

	group.POST("/policy", casbin.InsertPolicy)
	group.GET("/policy", casbin.GetPolicy)

	group.GET("/model", casbin.GetModel)

	group.POST("/enforce", casbin.Enforce)

	group.GET("/info", casbin.GetInfo)
}
