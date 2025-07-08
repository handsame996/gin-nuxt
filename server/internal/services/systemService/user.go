package systemService

import (
	"context"
	"example/template/internal/auth/jwt"
	"example/template/internal/models/systemModel"
	"example/template/internal/repositories/systemRepository"
	"example/template/pkg/utils"
)

type UserService struct {
	repository *systemRepository.UserRepository
	jwt        jwt.JWTInterface
}

func NewUserService(repository *systemRepository.UserRepository, jwt jwt.JWTInterface) *UserService {
	return &UserService{
		repository: repository,
		jwt:        jwt,
	}
}

func (server *UserService) CreateUser(ctx context.Context, user systemModel.User) (err error) {
	salt, pw, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = pw
	user.Salt = salt
	err = server.repository.CreateUserInfo(ctx, &user)
	if err != nil {
		return
	}
	return
}

func (server *UserService) Login(ctx context.Context, user systemModel.User) (token string, err error) {
	userInfo, err := server.repository.GetUserInfoByUserName(ctx, &user)
	if err != nil {
		return "", err
	}
	password, err := utils.VerifyPassword(user.Password, userInfo.Salt)
	if err != nil {
		return "", err
	}
	if password != userInfo.Password {
		return "", nil
	}
	token, err = server.jwt.CreateToken(userInfo.ID)
	if err != nil {
		return "", err
	}
	return
}
