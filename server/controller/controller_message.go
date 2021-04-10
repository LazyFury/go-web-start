package controller

import (
	"github.com/lazyfury/go-web-start/server/model"
	"github.com/lazyfury/go-web-template/controller"
)

// NewMessageController NewMessageController
func NewMessageController() *controller.Controller {
	return &controller.Controller{
		DB:    model.DB,
		Model: &model.Message{},
		Auth:  authWithFilter("to_user"),
	}
}

// NewMessageTemplateController NewMessageTemplateController
func NewMessageTemplateController() *controller.Controller {
	return &controller.Controller{
		DB:    model.DB,
		Model: &model.MessageTemplate{},
		Auth:  defaultAuth(),
	}
}
