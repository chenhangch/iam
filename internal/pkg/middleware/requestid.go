package middleware

import "github.com/gin-gonic/gin"

const (
	// XRequestIDKey defines X-Request-ID key string.
	XRequestIDKey = "X-Request-ID"
)

// RequestID 是一个中间件，它在每个请求的上下文和请求/响应头中注入一个'X-Request-ID'。
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
