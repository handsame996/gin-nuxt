package middleware

import (
	"example/template/internal/models"
	"github.com/gin-gonic/gin"
)

func (m *middlewareStruct) AddGlobalContext(g *models.Global) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("global", g)
		c.Next()
	}
}
