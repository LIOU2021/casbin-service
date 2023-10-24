package model

type Casbin struct {
	Sub string `json:"sub" form:"sub" binding:"required"`
	Obj string `json:"obj" form:"obj" binding:"required"`
	Act string `json:"act" form:"act" binding:"required"`
}
