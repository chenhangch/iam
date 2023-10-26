package middleware

import (
	"github.com/chang144/iam/pkg/log"
	"github.com/gin-gonic/gin"
)

const UsernameKey = "username"

// Context 上下文
func Context() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(log.KeyRequestID, c.GetString(XRequestIDKey))
		c.Set(log.KeyUsername, c.GetString(UsernameKey))
		c.Next()
	}
}
