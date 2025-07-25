package gormDB

import (
	"example/template/internal/models/systemModel"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&systemModel.User{}, &systemModel.SysDept{}, &systemModel.SysJob{}, &systemModel.SysMenu{},
		&systemModel.SysPermissions{}, &systemModel.SysRole{}, &systemModel.SysRoleMenu{}, &systemModel.SysRolePermissions{},
		&systemModel.SysUserJob{}, &systemModel.SysUserRole{})
	if err != nil {
		return err
	}
	setupGlobalHooks(db)
	return nil
}

func setupGlobalHooks(db *gorm.DB) {
	db.Callback().Create().Before("gorm:before_create").Register("before_create_hook", beforeCreateHook)
	db.Callback().Create().Before("gorm:before_update").Register("before_update_hook", beforeUpdateHook)
}

func beforeCreateHook(db *gorm.DB) {
	if db.Error != nil {
		return
	}
	if db.Statement.Schema != nil {
		idField := db.Statement.Schema.PrimaryFields[0]
		if idField.DataType == "uint32" {
			_, zero := idField.ValueOf(db.Statement.Context, db.Statement.ReflectValue)
			if zero {
				idField.Set(db.Statement.Context, db.Statement.ReflectValue, uuid.New().ID())
				db.Statement.SetColumn("created_at", time.Now().Format("2006-01-02 15:04:05"))
				db.Statement.AddClause(clause.OnConflict{
					DoNothing: true,
				})
			}
		}
		//db.Statement.SetColumn("id", uuid.New().String())
		//db.Statement.SetColumn("created_at", time.Now().Format("2006-01-02 15:04:05"))
	}
}

func beforeUpdateHook(db *gorm.DB) {
	if db.Error != nil {
		return
	}
	if db.Statement.Schema != nil {
		db.Statement.SetColumn("updated_at", time.Now().Format("2006-01-02 15:04:05"))
	}
}
