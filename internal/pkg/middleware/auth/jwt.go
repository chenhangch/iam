package auth

import (
	ginjwt "github.com/appleboy/gin-jwt/v2"
	"github.com/chang144/iam/internal/pkg/middleware"
	"github.com/gin-gonic/gin"
)

// AuthzAudience defines the value of jwt audience field.
const AuthzAudience = "iam.authz.jwt.com"

type JWTStrategy struct {
	ginjwt.GinJWTMiddleware
}

func NewJWTStrategy(gjwt ginjwt.GinJWTMiddleware) JWTStrategy {
	return JWTStrategy{gjwt}
}

func (j JWTStrategy) AuthFunc() gin.HandlerFunc {
	return j.MiddlewareFunc()
}

var _ middleware.AuthStrategy = &JWTStrategy{}
