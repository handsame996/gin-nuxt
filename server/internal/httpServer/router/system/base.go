package system

import (
	systemHandle "example/template/internal/httpServer/handlers/system"
	"example/template/internal/models"
	"example/template/internal/repositories/systemRepository"
	"example/template/internal/services/systemService"
	"github.com/gin-gonic/gin"
)

type User struct {
	g models.GlobalInter
}

func NewUser(g models.GlobalInter) *User {
	return &User{
		g: g,
	}
}

func (u *User) InitApiRouter(publicRouter *gin.RouterGroup, privateRouter *gin.RouterGroup) {
	// 依赖注入
	repository := systemRepository.NewUserRepository(u.g.GetGormDB())
	service := systemService.NewUserService(repository)
	var handler = systemHandle.NewUserHandler(service, u.g.GetLogger())
	// 设置router group
	apiPublicRouter := publicRouter.Group("base")
	apiPrivateRouter := privateRouter.Group("base")

	{
		apiPrivateRouter.POST("/createUser", handler.CreateUser)
	}
	{
		apiPublicRouter.POST("/login")

	}
}
