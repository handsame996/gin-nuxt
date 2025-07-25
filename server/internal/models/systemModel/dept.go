package systemModel

import "example/template/internal/models"

// 部门表：用户部门
type SysDept struct {
	models.GeneralModel
	Name         string `json:"name" form:"name" gorm:"column:name;comment:部门名称"`
	ContactEmail string `json:"contact_email" form:"contact_email" gorm:"column:contact_email;comment:联系邮箱"`
	ContactPhone string `json:"contact_phone" form:"contact_phone" gorm:"column:contact_phone;comment:联系电话"`
	ParentId     string `json:"parent_id" form:"parent_id" gorm:"column:parent_id;comment:父部门ID"`
}

func (SysDept) TableName() string {
	return "sys_dept"
}
