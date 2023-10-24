package model

type CasbinAddPolicy struct {
	Sub string `json:"sub" form:"sub" binding:"required"`
	Obj string `json:"obj" form:"obj" binding:"required"`
	Act string `json:"act" form:"act" binding:"required"`
}

type CasbinAddRole struct {
	Rsub string `json:"r_sub" form:"p_sub" binding:"required"`
	Psub string `json:"p_sub" form:"r_sub" binding:"required"`
}
