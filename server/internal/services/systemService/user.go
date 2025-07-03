package systemService

import (
	"context"
	"example/template/internal/models/systemModel"
	"example/template/internal/repositories/systemRepository"
	"example/template/pkg/utils"
)

type UserService struct {
	repository *systemRepository.UserRepository
}

func NewUserService(repository *systemRepository.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (server *UserService) CreateUser(ctx context.Context, user systemModel.User) (err error) {
	password, s, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = password
	user.Salt = s
	err = server.repository.CreateUserInfo(ctx, &user)
	if err != nil {
		return
	}
	return
}
