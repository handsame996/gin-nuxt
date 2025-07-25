package systemModel

import (
	"example/template/internal/models"
	"time"
)

type User struct {
	models.GeneralModel
	UserName   string     `json:"user_name" form:"user_name" gorm:"index:;column:user_name;comment:用户名"`
	Password   string     `json:"password" form:"password" gorm:"column:password;comment:密码"`
	Salt       string     `json:"salt" form:"salt" gorm:"column:salt;comment:盐"`
	FullName   string     `json:"full_name" form:"full_name" gorm:"column:full_name;comment:全名"`
	Email      string     `json:"email" form:"email" gorm:"index:;column:email;comment:邮箱"`
	Phone      string     `json:"phone" form:"phone" gorm:"index:;column:phone;comment:手机号"`
	Gender     int        `json:"gender" form:"gender" gorm:"type:tinyint(1);column:gender;comment:性别：0-未知，1-男，2-女"`
	Birthday   *time.Time `json:"birthday" form:"birthday" gorm:"type:date;column:birthday;comment:生日"`
	AvatarUrl  string     `json:"avatar_url" form:"avatar_url" gorm:"column:avatar_url;comment:头像url"`
	DeptId     string     `json:"dept_id" form:"dept_id" gorm:"column:dept_id;comment:部门id"`
	JobId      string     `json:"job_id" form:"job_id" gorm:"column:job_id;comment:职位id"`
	Enable     int        `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
	MiniOpenid string     `json:"mini_openid" form:"mini_openid" gorm:"column:mini_openid;comment:小程序id"`
	WxOpenid   string     `json:"wx_openid" form:"wx_openid" gorm:"column:wx_openid;comment:wx openid"`
	QQOpenid   string     `json:"qq_openid" form:"qq_openid" gorm:"column:qq_openid;comment:qq openid"`
}

func (User) TableName() string {
	return "sys_users"
}
