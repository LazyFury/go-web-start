package controller

import "github.com/Treblex/go-echo-demo/server/model"

// NewFeedbackController NewFeedbackController
func NewFeedbackController() *Controller {
	return &Controller{
		DB:    model.DB,
		Model: &model.Feedback{},
	}
}
