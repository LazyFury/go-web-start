package model

import (
	"github.com/Treblex/go-echo-demo/server/utils"
	"github.com/gin-gonic/gin"
)

// GetUserOrLogin GetUser
func GetUserOrLogin(c *gin.Context) *User {
	_user, exists := c.Get("user")
	if exists {
		user, ok := _user.(*User)
		if ok {
			return user
		}
	}
	panic(utils.JSON(utils.AuthedError, "请先登录", nil))
}

// GetUserOrEmpty GetUser
func GetUserOrEmpty(c *gin.Context) *User {
	_user, exists := c.Get("user")
	if exists {
		user, ok := _user.(*User)
		if ok {
			return user
		}
	}
	return &User{}
}
