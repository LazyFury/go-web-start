package controller

import "github.com/Treblex/go-echo-demo/server/model"

// NewMessageController NewMessageController
func NewMessageController() *Controller {
	return &Controller{
		DB:    model.DB,
		Model: &model.Message{},
	}
}
