package systemModel

import "example/template/internal/models"

// 用户角色表：用户角色中间表
//
// role_id ： 角色id ;
//
// user_id： 用户id ;
type SysUserRole struct {
	models.GeneralModel
	RoleId string `json:"role_id" form:"role_id" gorm:"column:role_id;comment:角色id"`
	UserId string `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户id"`
}

func (SysUserRole) TableName() string {
	return "sys_user_role"
}
