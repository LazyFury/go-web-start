package controller

import (
	"github.com/lazyfury/go-web-start/server/model"
	"github.com/lazyfury/go-web-template/controller"
)

// NewFeedbackController NewFeedbackController
func NewFeedbackController() *controller.Controller {
	return &controller.Controller{
		DB:    model.DB,
		Model: &model.Feedback{},
		Auth:  defaultAuth(),
	}
}
