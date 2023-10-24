package casbin

import (
	"casbin-service/core"
	"casbin-service/logger"
	"casbin-service/model"
	"encoding/json"
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

	if !ok {
		b, _ := json.Marshal(req)
		logger.ErrorfCtx(c, "InsertPolicy AddPolicy fail | req= %s", string(b))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.String(200, "SUCCESS")
}

// 创建角色
func InsertRole(c *gin.Context) {
	req := &model.CasbinAddRole{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		logger.ErrorfCtx(c, "InsertRole request unmarshal fail | err= %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ok, err := core.Enforcer.AddGroupingPolicy(req.Rsub, req.Psub)
	if err != nil {
		logger.ErrorfCtx(c, "InsertRole AddGroupingPolicy error | err= %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if !ok {
		b, _ := json.Marshal(req)
		logger.ErrorfCtx(c, "InsertRole AddGroupingPolicy fail | req= %s", string(b))
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
