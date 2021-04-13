package model

import (
	"github.com/gin-gonic/gin"
	"github.com/lazyfury/go-web-template/response"
)

// GetUserOrLogin GetUser
func GetUserOrLogin(c *gin.Context) *User {
	_user, exists := c.Get("user")
	if !exists {
		response.Error(response.JSON(response.AuthedError, "请先登录1", nil))
	}
	user, ok := _user.(*User)
	if !ok {
		response.Error(response.JSON(response.AuthedError, "请先登录", nil))
	}
	return user
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
