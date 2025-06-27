package gormDB

import (
	"example/templalte/internal/configs"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func InitializeGorm(param configs.Mysql) *gorm.DB {
	// 配置日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // 日志输出
		logger.Config{
			SlowThreshold:             time.Second, // 慢查询阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略记录未找到错误
			ParameterizedQueries:      true,        // 参数化查询
			Colorful:                  true,        // 彩色日志
		},
	)
	fmt.Println(param.DSN())
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: param.DSN(), // DSN data source name
	}), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Panicf("gorm connect mysql err! err  = %s", err)
		return nil
	}
	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(param.MaxIdleConns)  // 最大空闲连接数
	sqlDB.SetMaxOpenConns(param.MaxOpenConns)  // 最大打开连接数
	sqlDB.SetConnMaxLifetime(10 * time.Minute) // 连接最大存活时间
	return db
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate()
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
