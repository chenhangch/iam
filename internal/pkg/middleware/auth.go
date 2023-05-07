package middleware

import "github.com/gin-gonic/gin"

// AuthStrategy defines the set of methods used to do resource authentication.
// 定义用于执行资源身份验证的方法集
type AuthStrategy interface {
	AuthFunc() gin.HandlerFunc
}

// AuthOperator used to switch between different authentication strategy.
// 策略中心
type AuthOperator struct {
	strategy AuthStrategy
}

func (operator *AuthOperator) SetStrategy(strategy AuthStrategy) {
	operator.strategy = strategy
}

func (operator *AuthOperator) AuthFunc() gin.HandlerFunc {
	return operator.strategy.AuthFunc()
}
