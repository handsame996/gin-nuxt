package systemModel

import "example/template/internal/models"

// 用户角色表：用户角色中间表
//
// role_id ： 角色id ;
//
// user_id： 用户id ;
type SysUserJob struct {
	models.GeneralModel
	JobId  string `json:"job_id" form:"job_id" gorm:"column:job_id;comment:职位id"`
	UserId string `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户id"`
}

func (SysUserJob) TableName() string {
	return "sys_user_job"
}
