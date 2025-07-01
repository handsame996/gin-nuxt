package router

import (
	"example/template/internal/httpServer/router/system"
	"github.com/gin-gonic/gin"
)

func InitRouter(publicRouter *gin.RouterGroup, privateRouter *gin.RouterGroup) {
	system.Base{}.InitApiRouter(publicRouter)
}
