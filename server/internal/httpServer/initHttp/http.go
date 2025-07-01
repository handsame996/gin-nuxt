package initHttp

import (
	"example/template/internal/auth/casbinEnf"
	"example/template/internal/configs"
	"example/template/internal/httpServer/middleware"
	"example/template/internal/httpServer/router"
	"example/template/internal/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func InitHttpServer(g *models.Global, param configs.Http) {
	r := gin.New()
	r.Use(gin.Recovery())

	m := middleware.NewMiddlewareStruct()
	r.Use(m.Cors(), m.AddGlobalContext(g), m.Logger(g.Logger))

	publicRouter := r.Group("")
	privateRouter := r.Group("")
	privateRouter.Use(m.Casbin(casbinEnf.GetCasbin()))
	router.InitRouter(publicRouter, privateRouter)

	s := &http.Server{
		Addr:           param.Port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println(s.ListenAndServe())
}
