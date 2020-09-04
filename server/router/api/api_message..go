package api

import (
	"github.com/treblex/go-echo-demo/server/model"

	"github.com/labstack/echo/v4"
)

var modelMessage model.Message

func messages(g *echo.Group) {
	modelMessage.BaseControll.Model = &modelMessage
	modelMessage.Install(g, "/messages")
}
