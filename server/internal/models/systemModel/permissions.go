package systemModel

import "example/template/internal/models"

// 资源权限api
type SysPermissions struct {
	models.GeneralModel
	Path        string `json:"path" form:"path" gorm:"column:path;comment:后台接口path"`
	Description string `json:"description" form:"description" gorm:"column:description;comment:后台接口path"`
	ApiGroup    string `json:"api_group" form:"api_group" gorm:"column:api_group;comment:接口分组"`
	Method      string `json:"method" form:"method" gorm:"column:method;comment:操作方式 POST"`
}

func (SysPermissions) TableName() string {
	return "sys_permissions"
}
