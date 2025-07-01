package middleware

import (
	"github.com/gin-gonic/gin"
)

func (m *middlewareStruct) Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {


		c.Next()
	}
}
