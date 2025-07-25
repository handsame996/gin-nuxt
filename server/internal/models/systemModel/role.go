package systemModel

import "example/template/internal/models"

// 系统角色表
type SysRole struct {
	models.GeneralModel
	RoleName string `json:"role_name" form:"role_name" gorm:"column:role_name;comment:角色名称"`
	RoleCode string `json:"role_code" form:"role_code" gorm:"column:role_code;comment:角色代码"`
	RoleDesc string `json:"role_desc" form:"role_desc" gorm:"column:role_desc;comment:角色详情"`
	ParentId *uint  `json:"parentId" gorm:"column:parent_id;comment:父角色ID"`
}

func (SysRole) TableName() string {
	return "sys_role"
}
