package auth

import (
	"github.com/chang144/golunzi/errors"
	"github.com/chang144/iam/internal/pkg/code"
	"github.com/chang144/iam/internal/pkg/middleware"
	"github.com/chang144/iam/pkg/core"
	"github.com/gin-gonic/gin"
	"strings"
)

const authHandlerCount = 2

type AutoStrategy struct {
	basic middleware.AuthStrategy
	jwt   middleware.AuthStrategy
}

// NewAutoStrategy create auto strategy with basic strategy and jwt strategy.
func NewAutoStrategy(basic, jwt middleware.AuthStrategy) AutoStrategy {
	return AutoStrategy{
		basic: basic,
		jwt:   jwt,
	}
}

func (a AutoStrategy) AuthFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		operation := middleware.AuthOperator{}
		authHandler := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(authHandler) != authHandlerCount {
			core.WriteResponse(c,
				errors.WithCode(code.ErrSignatureInvalid, "Authorization header format is wrong"),
				nil)
			c.Abort()
			return
		}

		switch authHandler[0] {
		case "Basic":
			operation.SetStrategy(a.basic)
		case "Bearer":
			operation.SetStrategy(a.jwt)
		default:
			core.WriteResponse(c, errors.WithCode(code.ErrSignatureInvalid, "unrecognized authorization header."), nil)
			c.Abort()

			return
		}
		operation.AuthFunc()(c)

		c.Next()
	}
}

var _ middleware.AuthStrategy = &AutoStrategy{}
