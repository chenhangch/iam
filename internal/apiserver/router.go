package apiserver

import (
	"fmt"
	"github.com/chang144/iam/internal/apiserver/store/mysql"
	"github.com/chang144/iam/internal/pkg/middleware/auth"
	"github.com/gin-gonic/gin"
)

func initRouter(g *gin.Engine) {
	installController(g)
}

func installController(g *gin.Engine) *gin.Engine {
	// authz
	jwtStrategy := newJWTAuth().(auth.JWTStrategy)
	g.POST("/login", jwtStrategy.LoginHandler)
	g.POST("/logout", jwtStrategy.LogoutHandler)
	g.POST("/refresh", jwtStrategy.RefreshHandler)

	// 创建一个空的dbstore
	storeIns, err := mysql.GetMySQLFactoryOr(nil)
	if err != nil {
		return nil
	}

	fmt.Println(storeIns)
	return g
}
