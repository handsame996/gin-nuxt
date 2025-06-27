package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (middlewareStruct) Jwt(e *casbin.SyncedCachedEnforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.GetHeader("X-User-ID")
		if user == "" {
			user = "anonymous"
		}

		obj := c.Request.URL.Path
		act := c.Request.Method

		ok, err := e.Enforce(user, obj, act)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "权限检查失败"})
			c.Abort()
			return
		}

		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "权限不足"})
			c.Abort()
			return
		}

		c.Next()
	}
}
