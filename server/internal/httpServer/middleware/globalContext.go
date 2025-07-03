package middleware

import (
	"example/template/internal/models"
	"github.com/gin-gonic/gin"
)

func (m *middlewareStruct) AddGlobalContext(g models.GlobalInter) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("global", g)
		c.Next()
	}
}
