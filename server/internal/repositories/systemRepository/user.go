package systemRepository

import (
	"context"
	"example/template/internal/models/systemModel"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
func (repository *UserRepository) CreateUserInfo(ctx context.Context, user *systemModel.User) (err error) {
	err = repository.db.WithContext(ctx).Model(&systemModel.User{}).Create(user).Error
	return
}
