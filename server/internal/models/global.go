package models

import (
	"example/template/internal/auth/jwt"
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
	jwt    jwt.JWTInterface
}

type GlobalInter interface {
	GetGormDB() *gorm.DB
	GetLogger() *zap.Logger
	GetJWTInterface() jwt.JWTInterface
}

func CreateGlobal(db *gorm.DB, logger *zap.Logger, jwt jwt.JWTInterface) GlobalInter {
	return &global{
		db:     db,
		logger: logger,
		jwt:    jwt,
	}
}

func (g *global) GetGormDB() *gorm.DB {
	return g.db
}

func (g *global) GetLogger() *zap.Logger {
	return g.logger
}

func (g *global) GetJWTInterface() jwt.JWTInterface {
	return g.jwt
}
