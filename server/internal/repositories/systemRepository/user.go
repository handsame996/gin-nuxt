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

func (repository *UserRepository) GetUserInfoByUserName(ctx context.Context, user *systemModel.User) (userInfo systemModel.User, err error) {
	err = repository.db.WithContext(ctx).Where("user_name = ?", user.UserName).Debug().First(&userInfo).Error
	return
	//if err != nil {
	//	return false, err
	//}
	//password, err := utils.VerifyPassword(user.Password, userInfo.Salt)
	//if err != nil {
	//	return false, err
	//}
	//if password != userInfo.Password {
	//	return false, nil
	//}
	//return true, nil
}
