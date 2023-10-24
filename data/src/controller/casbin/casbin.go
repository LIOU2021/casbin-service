package casbin

import (
	"casbin-service/core"
	"casbin-service/logger"
	"casbin-service/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 创建政策
func InsertPolicy(c *gin.Context) {
	req := &model.CasbinAddPolicy{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		logger.ErrorfCtx(c, "InsertPolicy request unmarshal fail | err= %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ok, err := core.Enforcer.AddPolicy(req.Sub, req.Obj, req.Act)
	if err != nil {
		logger.ErrorfCtx(c, "InsertPolicy AddPolicy error | err= %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	defer func() {
		logger.InfofCtx(c, "InsertPolicy | ok= %t | sub: %s, obj: %s, act: %s", ok, req.Sub, req.Obj, req.Act)
	}()

	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.String(200, "SUCCESS")
}

// 创建用戶或群組關聯
func AddGroupingPolicy(c *gin.Context) {
	req := &model.CasbinAddRole{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		logger.ErrorfCtx(c, "AddGroupingPolicy request unmarshal fail | err= %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ok, err := core.Enforcer.AddGroupingPolicy(req.Rsub, req.Psub)
	if err != nil {
		logger.ErrorfCtx(c, "AddGroupingPolicy AddGroupingPolicy error | err= %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	defer func() {
		logger.InfofCtx(c, "AddGroupingPolicy | ok= %t | r.sub: %s, p.sub: %s", ok, req.Rsub, req.Psub)
	}()

	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.String(200, "SUCCESS")
}

// 验证政策
func Enforce(c *gin.Context) {
	req := &model.CasbinEnforce{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		logger.ErrorfCtx(c, "enforce request unmarshal fail | err= %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ok, err := core.Enforcer.Enforce(req.Sub, req.Obj, req.Act)

	defer func() {
		logger.InfofCtx(c, "enforce access | ok= %t | sub: %s, obj: %s, act: %s", ok, req.Sub, req.Obj, req.Act)
	}()

	if err != nil {
		logger.ErrorfCtx(c, "enforce error | sub: %s, obj: %s, act: %s | err= %v", ok, req.Sub, req.Obj, req.Act, err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if !ok {
		c.String(http.StatusForbidden, "FORBIDDEN")
		return
	}

	c.String(200, "SUCCESS")
}

func GetPolicy(c *gin.Context) {
	c.JSON(200, core.Enforcer.GetPolicy())
}

func GetRole(c *gin.Context) {
	c.JSON(200, core.Enforcer.GetGroupingPolicy())
}

func GetModel(c *gin.Context) {
	c.String(200, core.Enforcer.GetModel().ToText())
}

func GetInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"all_roles":   core.Enforcer.GetAllRoles(),
		"all_actions": core.Enforcer.GetAllActions(),
		"all_objects": core.Enforcer.GetAllObjects(),
	})
}

func GetRoleForUser(c *gin.Context) {
	req := &model.CasbinFindRoleForUser{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		logger.ErrorfCtx(c, "GetRoleForUser request unmarshal fail | err= %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	list, err := core.Enforcer.GetRolesForUser(req.User)
	if err != nil {
		logger.ErrorfCtx(c, "GetRoleForUser GetRolesForUser fail | err= %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(200, list)
}

// 刪除指定腳色
// 如果刪除develop，但A用戶是develop，那麼A用戶也會被刪除
func DeleteRole(c *gin.Context) {
	req := &model.CasbinDeleteRole{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		logger.ErrorfCtx(c, "DeleteRole request unmarshal fail | err= %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ok, err := core.Enforcer.DeleteRole(req.Sub)
	if err != nil {
		logger.ErrorfCtx(c, "DeleteRole fail | err= %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	defer func() {
		logger.InfofCtx(c, "DeleteRole | ok= %t | sub: %s", ok, req.Sub)
	}()

	if !ok {
		c.String(http.StatusBadRequest, "FAIL")
		return
	}
	c.String(200, "SUCCESS")
}

// 刪除指定用戶底下的特定權限腳色
// 比如A用戶具有a腳色跟b腳色，可以刪除A用戶跟a腳色的關聯，這樣A用戶就只剩下b腳色
func DeleteRoleForUser(c *gin.Context) {
	req := &model.CasbinDeleteRoleForUser{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		logger.ErrorfCtx(c, "DeleteRoleForUser request unmarshal fail | err= %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ok, err := core.Enforcer.DeleteRoleForUser(req.Rsub, req.Psub)
	if err != nil {
		logger.ErrorfCtx(c, "DeleteRoleForUser fail | err= %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	defer func() {
		logger.InfofCtx(c, "DeleteRoleForUser | ok= %t | r.sub: %s, p.sub", ok, req.Rsub, req.Psub)
	}()

	if !ok {
		c.String(http.StatusBadRequest, "FAIL")
		return
	}
	c.String(200, "SUCCESS")
}

// 刪除用戶
// 只會刪除g，不會刪除到p的
// 比如A用戶擁有a腳色跟b腳色，刪除A用戶時，A用戶與a跟b腳色關聯的資料將會刪除
// 但a跟b腳色依然存在資料庫，可以供其他用戶再行關聯，不必重新創建policy
func DeleteUser(c *gin.Context) {
	req := &model.CasbinDeleteUser{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		logger.ErrorfCtx(c, "DeleteUser request unmarshal fail | err= %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ok, err := core.Enforcer.DeleteUser(req.User)
	if err != nil {
		logger.ErrorfCtx(c, "DeleteUser fail | err= %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	defer func() {
		logger.InfofCtx(c, "DeleteUser | ok= %t | user: %s", ok, req.User)
	}()

	if !ok {
		c.String(http.StatusBadRequest, "FAIL")
		return
	}
	c.String(200, "SUCCESS")
}
