// Package api 接口定义了用户model
package api

// User 定义了User资源创建和返回的字段
type User struct {
	// 用户的名字
	// Request: true
	Name string `json:"name"`

	// 用户的昵称
	// Request: true
	Nickname string `json:"nickname"`

	// 用户的家庭住址
	Address string `json:"address"`

	Email string `json:"email"`
}
