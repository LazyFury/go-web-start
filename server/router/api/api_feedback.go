package api

import (
	"github.com/treblex/go-echo-demo/server/model"

	"github.com/labstack/echo/v4"
)

var modelFeedback model.Feedback

func feedback(g *echo.Group) {
	modelFeedback.BaseControll.Model = &modelFeedback
	modelFeedback.Install(g, "/feedbacks")
}
