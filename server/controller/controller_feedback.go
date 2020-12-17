package controller

import (
	"github.com/Treblex/go-echo-demo/server/model"
	"github.com/Treblex/go-web-template/controller"
)

// NewFeedbackController NewFeedbackController
func NewFeedbackController() *controller.Controller {
	return &controller.Controller{
		DB:    model.DB,
		Model: &model.Feedback{},
	}
}
