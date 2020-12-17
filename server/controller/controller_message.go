package controller

import (
	"github.com/Treblex/go-echo-demo/server/model"
	"github.com/Treblex/go-web-template/controller"
)

// NewMessageController NewMessageController
func NewMessageController() *controller.Controller {
	return &controller.Controller{
		DB:    model.DB,
		Model: &model.Message{},
		Auth:  auth("to_user"),
	}
}

// NewMessageTemplateController NewMessageTemplateController
func NewMessageTemplateController() *controller.Controller {
	return &controller.Controller{
		DB:    model.DB,
		Model: &model.MessageTemplate{},
	}
}
