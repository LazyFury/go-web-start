package api

import (
	"github.com/Treblex/go-echo-demo/server/model"
	"github.com/gin-gonic/gin"
)

var modelOrder model.Order

func order(g *gin.RouterGroup) {
	modelOrder.BaseControll.Model = &modelOrder
	modelOrder.Install(g, "/orders")
}
