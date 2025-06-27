package initHttp

import (
	"example/templalte/internal/configs"
	"example/templalte/internal/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func InitHttpServer(g *models.Global, param configs.Http) {
	router := gin.New()
	router.Use(func(c *gin.Context) {
		c.Set("global", g)
		c.Next()
	})
	router.Use(gin.Recovery())
	s := &http.Server{
		Addr:           param.Port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println(s.ListenAndServe())
}
