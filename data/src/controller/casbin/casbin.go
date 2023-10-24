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
	req := &model.Casbin{}
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
	core.Enforcer.AddGroupingPolicy()
	c.String(200, "hi insert Role")
}

func GetPolicy(c *gin.Context) {
	c.JSON(200, core.Enforcer.GetPolicy())
}

func GetModel(c *gin.Context) {
	c.String(200, core.Enforcer.GetModel().ToText())
}
