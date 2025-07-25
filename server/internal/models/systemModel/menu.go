package systemModel

import "example/template/internal/models"

// 菜单权限表
type SysMenu struct {
	models.GeneralModel
	Name string `json:"name" form:"name" gorm:"column:name;comment:菜单名称"`
	//Permission string `json:"permission" form:"permission" gorm:"column:permission;comment:菜单权限标识"`
	Hidden    bool                                       `json:"hidden" gorm:"comment:是否在列表隐藏"` // 是否在列表隐藏
	Path      string                                     `json:"path" form:"path" gorm:"column:path;comment:前端URL"`
	ParentId  string                                     `json:"parent_id" form:"parent_id" gorm:"column:path;comment:父菜单ID"`
	Icon      string                                     `json:"icon" form:"icon" gorm:"column:icon;comment:图标"`
	Component string                                     `json:"component" form:"component" gorm:"column:component;comment:页面路径"`
	Sort      int                                        `json:"sort" form:"sort" gorm:"column:sort;comment:排序"`
	Meta      `json:"meta" gorm:"embedded;comment:附加属性"` // 附加属性
	KeepAlive bool                                       `json:"keep_alive" form:"keep_alive" gorm:"column:keep_alive;comment:页面缓存"`
	//Type      byte   `json:"type" form:"type" gorm:"column:type;comment:菜单类型"`
}

type Meta struct {
	ActiveName     string `json:"activeName" gorm:"comment:高亮菜单"`
	KeepAlive      bool   `json:"keepAlive" gorm:"comment:是否缓存"`           // 是否缓存
	DefaultMenu    bool   `json:"defaultMenu" gorm:"comment:是否是基础路由（开发中）"` // 是否是基础路由（开发中）
	Title          string `json:"title" gorm:"comment:菜单名"`                // 菜单名
	Icon           string `json:"icon" gorm:"comment:菜单图标"`                // 菜单图标
	CloseTab       bool   `json:"closeTab" gorm:"comment:自动关闭tab"`         // 自动关闭tab
	TransitionType string `json:"transitionType" gorm:"comment:路由切换动画"`    // 路由切换动画
}

func (SysMenu) TableName() string {
	return "sys_menu"
}
