package middleware

import (
	"example/template/internal/auth/jwt"
	"example/template/internal/httpServer/responses"
	"github.com/gin-gonic/gin"
	"strings"
)

func (m *middlewareStruct) Jwt(jwt jwt.JWTInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			responses.NoAuth("未登录或非法访问", c)
			c.Abort()
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			responses.NoAuth("令牌格式错误", c)
			c.Abort()
			return
		}
		userId, err := jwt.ParseToken(parts[1])
		if err != nil {
			responses.NoAuth(err.Error(), c)
			c.Abort()
			return
		}
		c.Set("userId", userId)
		c.Next()
	}
}
