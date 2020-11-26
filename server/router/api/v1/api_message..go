package api

import (
	"github.com/Treblex/go-echo-demo/server/model"
	"github.com/gin-gonic/gin"
)

var modelMessage model.Message

func messages(g *gin.RouterGroup) {
	modelMessage.BaseControll.Model = &modelMessage
	modelMessage.Install(g, "/messages")
}
