package models

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

// 通用模型
type GeneralModel struct {
	ID        int            `gorm:"primarykey" json:"id"` // 主键ID
	CreatedAt time.Time      `gorm:"index" json:"-"`       // 创建时间
	UpdatedAt time.Time      `gorm:"index" json:"-"`       // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`       // 删除时间
}

type global struct {
	db     *gorm.DB
	logger *zap.Logger
}

type GlobalInter interface {
	GetGormDB() *gorm.DB
	GetLogger() *zap.Logger
}

func CreateGlobal(db *gorm.DB, logger *zap.Logger) GlobalInter {
	return &global{
		db:     db,
		logger: logger,
	}
}

func (g *global) GetGormDB() *gorm.DB {
	return g.db
}

func (g *global) GetLogger() *zap.Logger {
	return g.logger
}
