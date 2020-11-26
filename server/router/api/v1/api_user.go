package api

import (
	"github.com/Treblex/go-echo-demo/server/model"
	"github.com/gin-gonic/gin"
)

var modelUser model.User

// 用户API
func user(g *gin.RouterGroup) {
	modelUser.BaseControll.Model = &modelUser
	modelUser.Install(g, "/users")
}
