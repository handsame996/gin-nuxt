package models

import "gorm.io/gorm"

type Global struct {
	DB *gorm.DB
}
