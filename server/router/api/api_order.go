package api

import (
	"EK-Server/model"

	"github.com/labstack/echo"
)

var modelOrder model.Order

func order(g *echo.Group) {
	modelOrder.BaseControll.Model = &modelOrder
	modelOrder.Install(g, "/orders")
}
