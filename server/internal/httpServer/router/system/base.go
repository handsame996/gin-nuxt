package system

import "github.com/gin-gonic/gin"

type Base struct{}

func (Base) InitApiRouter(router *gin.RouterGroup) {
	//var handler = Base.NewUserHandlers()
	apiRouter := router.Group("base")
	{
		apiRouter.POST("/login", func(context *gin.Context) {
			
		})
	}
}
