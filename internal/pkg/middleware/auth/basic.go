package auth

import (
	"encoding/base64"
	"github.com/chang144/golunzi/errors"
	"github.com/chang144/iam/internal/pkg/code"
	"github.com/chang144/iam/internal/pkg/middleware"
	"github.com/chang144/iam/pkg/core"
	"github.com/gin-gonic/gin"
	"strings"
)

// BasicStrategy 定义 Basic authentication strategy
// 基础认证（Basic认证）就是最简单的认证方式，简单地将“用户名、密码”进行base64编码后，放进HTTP Authorization header中
type BasicStrategy struct {
	compare basicCompare
}

var _ middleware.AuthStrategy = &BasicStrategy{}

type basicCompare func(username, password string) bool

func NewBasicStrategy(compare basicCompare) BasicStrategy {
	return BasicStrategy{
		compare: compare,
	}
}

// AuthFunc 实现 basic 认证的
func (b BasicStrategy) AuthFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			core.WriteResponse(
				c,
				errors.WithCode(code.ErrSignatureInvalid, "Authorization header format is wrong"),
				nil,
			)
			c.Abort()

			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !b.compare(pair[0], pair[1]) {
			core.WriteResponse(
				c,
				errors.WithCode(code.ErrSignatureInvalid, "Authorization header format is wrong."),
				nil,
			)
			c.Abort()

			return
		}
		c.Set(middleware.UsernameKey, pair[0])

		c.Next()
	}
}
