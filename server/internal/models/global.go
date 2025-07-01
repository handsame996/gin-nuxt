package models

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Global struct {
	DB *gorm.DB
	Logger *zap.Logger
}
