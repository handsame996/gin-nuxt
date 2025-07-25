package systemModel

import "example/template/internal/models"

// 角色权限表：角色权限中间表
//
// role_id ： 角色id ;
//
// permissions_id： 权限id ;
type SysRolePermissions struct {
	models.GeneralModel
	RoleId        string `json:"role_id" form:"role_id" gorm:"column:role_id;comment:角色id"`
	PermissionsId string `json:"permissions_id" form:"permissions_id" gorm:"column:permissions_id;comment:权限id"`
}

func (SysRolePermissions) SysPermissions() string {
	return "sys_role_permissions"
}
