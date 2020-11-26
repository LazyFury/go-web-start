package api

import (
	"github.com/Treblex/go-echo-demo/server/model"
	"github.com/gin-gonic/gin"
)

var modelFeedback model.Feedback

func feedback(g *gin.RouterGroup) {
	modelFeedback.BaseControll.Model = &modelFeedback
	modelFeedback.Install(g, "/feedbacks")
}
