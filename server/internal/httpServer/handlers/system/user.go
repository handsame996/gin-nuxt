package systemHandle

import (
	"example/template/internal/httpServer/responses"
	"example/template/internal/models/systemModel"
	"example/template/internal/services/systemService"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	service *systemService.UserService
	logger  *zap.Logger
}

func NewUserHandler(service *systemService.UserService, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		service: service,
		logger:  logger,
	}
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	var user systemModel.User
	if err := c.ShouldBindJSON(&user); err != nil {
		responses.FailWithMessage(err.Error(), c)
		return
	}

	if err := u.service.CreateUser(c, user); err != nil {
		u.logger.Error("创建失败!", zap.Error(err))
		responses.FailWithMessage("创建失败", c)
		return
	}
	responses.OkWithMessage("创建成功", c)
}
