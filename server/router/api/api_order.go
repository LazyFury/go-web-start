package api

import (
	"github.com/treblex/go-echo-demo/server/model"

	"github.com/labstack/echo/v4"
)

var modelOrder model.Order

func order(g *echo.Group) {
	modelOrder.BaseControll.Model = &modelOrder
	modelOrder.Install(g, "/orders")
}
