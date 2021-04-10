package controller

import (
	"github.com/lazyfury/go-web-start/server/model"
	"github.com/lazyfury/go-web-template/controller"
)

// NewAdController NewAdController
func NewAdController() *controller.Controller {
	return &controller.Controller{
		DB:    model.DB,
		Model: &model.Ad{},
		Auth:  defaultAuth(),
	}
}

// NewAdGroupController NewAdGroupController
func NewAdGroupController() *controller.Controller {
	return &controller.Controller{
		DB:    model.DB,
		Model: &model.AdGroup{},
		Auth:  defaultAuth(),
	}
}

// NewAdEventController NewAdEventController
func NewAdEventController() *controller.Controller {
	return &controller.Controller{
		DB:    model.DB,
		Model: &model.AdEvent{},
		Auth:  defaultAuth(),
	}
}
