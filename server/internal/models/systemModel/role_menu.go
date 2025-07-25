package systemModel

import "example/template/internal/models"

// 角色菜单表：角色菜单中间表
//
// role_id ： 角色id ;
//
// menu_id： 菜单id ;
type SysRoleMenu struct {
	models.GeneralModel
	RoleId string `json:"role_id" form:"role_id" gorm:"column:role_id;comment:角色id"`
	MenuId string `json:"menu_id" form:"menu_id" gorm:"column:menu_id;comment:菜单id"`
}

func (SysRoleMenu) SysPermissions() string {
	return "sys_role_menu"
}
