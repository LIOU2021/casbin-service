package model

type CasbinAddPolicy struct {
	Sub string `json:"sub" form:"sub" binding:"required"`
	Obj string `json:"obj" form:"obj" binding:"required"`
	Act string `json:"act" form:"act" binding:"required"`
}

type CasbinAddRole struct {
	Rsub string `json:"r_sub" form:"p_sub" binding:"required"` // 用戶
	Psub string `json:"p_sub" form:"r_sub" binding:"required"` // 權限角色
}

type CasbinEnforce struct {
	Sub string `json:"sub" form:"sub" binding:"required"`
	Obj string `json:"obj" form:"obj" binding:"required"`
	Act string `json:"act" form:"act" binding:"required"`
}

type CasbinFindRoleForUser struct {
	User string `json:"user" form:"user" binding:"required"` // 就是sub
}

// 移除p.sub
type CasbinDeleteRole struct {
	Sub string `json:"sub" form:"sub" binding:"required"`
}

// 移除用戶的某個權限腳色
type CasbinDeleteRoleForUser struct {
	Rsub string `json:"r_sub" form:"p_sub" binding:"required"` // 用戶
	Psub string `json:"p_sub" form:"r_sub" binding:"required"` // 權限角色
}

// 刪除用戶p.sub
type CasbinDeleteUser struct {
	User string `json:"user" form:"user" binding:"required"`
}
