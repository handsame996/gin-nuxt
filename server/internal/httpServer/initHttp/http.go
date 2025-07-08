package initHttp

import (
	"example/template/internal/configs"
	"example/template/internal/httpServer/middleware"
	"example/template/internal/httpServer/router"
	"example/template/internal/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func Server(g models.GlobalInter, param configs.Http) {
	r := gin.New()
	r.Use(gin.Recovery())

	m := middleware.NewMiddlewareStruct()
	r.Use(m.Cors(), m.Logger(g.GetLogger()))
	r.POST("ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "pong")
	})

	publicRouter := r.Group("")
	privateRouter := r.Group("")

	//privateRouter.Use(m.Casbin(casbinEnf.GetCasbin()))
	privateRouter.Use(m.Jwt(g.GetJWTInterface()))
	router.InitRouter(publicRouter, privateRouter, g)
	privateRouter.POST("ping1", func(context *gin.Context) {
		context.JSON(http.StatusOK, "pong")
	})
	s := &http.Server{
		Addr:           param.Port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println(s.ListenAndServe())
}
