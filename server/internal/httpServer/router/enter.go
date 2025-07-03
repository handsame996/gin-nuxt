package router

import (
	"example/template/internal/httpServer/router/system"
	"example/template/internal/models"
	"github.com/gin-gonic/gin"
)

func InitRouter(publicRouter *gin.RouterGroup, privateRouter *gin.RouterGroup, g models.GlobalInter) {
	userGroup := system.NewUser(g)
	userGroup.InitApiRouter(publicRouter, privateRouter)
}
