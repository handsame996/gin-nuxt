package systemModel

import "example/template/internal/models"

// 职位表
type SysJob struct {
	models.GeneralModel
	Name   string `json:"name" form:"name" gorm:"column:name;comment:职位名称"`
	DeptId string `json:"dept_id" form:"dept_id" gorm:"column:dept_id;comment:部门id"`
}

func (SysJob) TableName() string {
	return "sys_job"
}
