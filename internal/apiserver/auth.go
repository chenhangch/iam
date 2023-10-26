package apiserver

import (
	"encoding/base64"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/chang144/iam/internal/pkg/middleware"
	"github.com/chang144/iam/internal/pkg/middleware/auth"
	"github.com/chang144/iam/pkg/log"
	"github.com/gin-gonic/gin"
	v1 "github.com/hangcodebug/go-api-module/iam/apiserver/v1"
	"github.com/spf13/viper"
	"net/http"
	"strings"
	"time"
)

const (
	// APIServerAudience defines the value of jwt audience field.
	APIServerAudience = "iam.api.ch.com"

	// APIServerIssuer defines the value of jwt issuer field.
	APIServerIssuer = "iam-apiserver"
)

type loginInfo struct {
	Username string `form:"username" json:"username" binding:"username"`
	Password string `form:"password" json:"password" binding:"required,password"`
}

func newBasicAuth() middleware.AuthStrategy {
	return auth.NewBasicStrategy(func(username, password string) bool {
		// TODO: fetch user from database

		// TODO: Compare password

		// TODO: Update Login time
		return true
	})
}

func newJWTAuth() middleware.AuthStrategy {
	ginjwt, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            viper.GetString("jwt.Realm"),
		SigningAlgorithm: "HS256",
		Key:              []byte(viper.GetString("jwt.key")),
		Timeout:          viper.GetDuration("jwt.timeout"),
		MaxRefresh:       viper.GetDuration("jwt.max-refresh"),
		Authenticator:    authenticator(),
		Authorizator:     authorizator(),
		PayloadFunc:      payloadFunc(),
		Unauthorized:     nil,
		LoginResponse:    loginResponse(),
		LogoutResponse: func(c *gin.Context, code int) {
			c.JSON(http.StatusOK, nil)
		},
		RefreshResponse: refreshResponse(),
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)

			return claims[jwt.IdentityKey]
		},
		IdentityKey:           middleware.UsernameKey,
		TokenLookup:           "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:         "Bearer",
		TimeFunc:              time.Now,
		HTTPStatusMessageFunc: nil,
		PrivKeyFile:           "",
		PubKeyFile:            "",
		SendCookie:            true,
		CookieMaxAge:          0,
		SecureCookie:          false,
		CookieHTTPOnly:        false,
		CookieDomain:          "",
		SendAuthorization:     false,
		DisabledAbort:         false,
		CookieName:            "",
		CookieSameSite:        0,
	})

	return auth.NewJWTStrategy(*ginjwt)
}

func authorizator() func(data interface{}, c *gin.Context) bool {
	return func(data interface{}, c *gin.Context) bool {
		if _, ok := data.(string); ok {
			log.L(c).Info("user %s is authenticated.")

			return true
		}
		return false
	}
}

func payloadFunc() func(data interface{}) jwt.MapClaims {
	return func(data interface{}) jwt.MapClaims {
		claims := jwt.MapClaims{
			"iss": APIServerIssuer,
			"aud": APIServerAudience,
		}
		if u, ok := data.(*v1.User); ok {
			claims[jwt.IdentityKey] = u.Name
			claims["sub"] = u.Name
		}

		return claims
	}
}

func loginResponse() func(*gin.Context, int, string, time.Time) {
	return func(c *gin.Context, code int, token string, expire time.Time) {
		c.JSON(http.StatusOK, gin.H{
			"token":  token,
			"expire": expire.Format(time.RFC3339),
		})
	}
}

func newAutoAuth() middleware.AuthStrategy {
	return auth.NewAutoStrategy(newBasicAuth().(auth.BasicStrategy), newJWTAuth().(auth.JWTStrategy))
}

// authenticatior 从http Authorization header 中获取用户名和密码、 并校验密码是否合法
func authenticator() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var login loginInfo
		var err error
		if c.Request.Header.Get("Authorization") != "" {
			login, err = parseWithHeader(c)
		} else {
			login, err = parseWithBody(c)
		}
		if err != nil {
			return "", jwt.ErrFailedAuthentication
		}
		fmt.Print(login)
		// TODO: get the user information by the login username

		// TODO: compare the login password with the user password

		// TODO: updata the login time

		return nil, nil
	}
}

func parseWithHeader(c *gin.Context) (loginInfo, error) {
	auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		log.Errorf("get basic string from Authorization header failed")

		return loginInfo{}, jwt.ErrFailedAuthentication
	}

	payload, err := base64.StdEncoding.DecodeString(auth[1])
	if err != nil {
		log.Errorf("decode basic string: %s", err.Error())

		return loginInfo{}, jwt.ErrFailedAuthentication
	}

	pair := strings.SplitN(string(payload), ":", 2)
	if len(pair) != 2 {
		log.Errorf("parse payload failed")

		return loginInfo{}, jwt.ErrFailedAuthentication
	}

	return loginInfo{
		Username: pair[0],
		Password: pair[1],
	}, nil
}

func parseWithBody(c *gin.Context) (loginInfo, error) {
	var login loginInfo
	if err := c.ShouldBindJSON(&login); err != nil {
		log.Errorf("parse login parameters: %s", err.Error())

		return loginInfo{}, jwt.ErrFailedAuthentication
	}

	return login, nil
}

func refreshResponse() func(c *gin.Context, code int, token string, expire time.Time) {
	return func(c *gin.Context, code int, token string, expire time.Time) {
		c.JSON(http.StatusOK, gin.H{
			"token":  token,
			"expire": expire.Format(time.RFC3339),
		})
	}
}
