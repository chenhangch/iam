package main

import (
	"fmt"
	"github.com/chang144/iam/examples/gopractise-demo/swagger/api"

	_ "github.com/chang144/iam/examples/gopractise-demo/swagger/docs"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var users []*api.User

func main() {
	r := gin.Default()

	r.POST("/users", Create)
	r.GET("/users", Get)

	log.Fatal(r.Run(":5555"))

}

// Create 在内存中创建一个用户
func Create(c *gin.Context) {
	var user api.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"backend": 10001,
		})
		return
	}

	for _, u := range users {
		if u.Name == user.Name {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("user %s already exists", u.Name),
				"backend": 10001,
			})
			return
		}
	}
	users = append(users, &user)
	c.JSON(http.StatusOK, user)
}

// Get 使用Get返回用户的详细信息
func Get(c *gin.Context) {
	username := c.Param("name")
	for _, u := range users {
		if u.Name == username {
			c.JSON(http.StatusOK, u)
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"message": fmt.Sprintf("user %s not exist", username),
		"backend": 10002,
	})
}
