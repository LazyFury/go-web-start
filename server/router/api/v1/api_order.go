package api

import (
	"github.com/Treblex/go-web-start/server/model"
	"github.com/gin-gonic/gin"
)

var modelOrder model.Order

func order(g *gin.RouterGroup) {
	modelOrder.BaseControll.Model = &modelOrder
	modelOrder.Install(g, "/orders")
}
