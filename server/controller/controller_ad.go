package controller

import (
	"github.com/Treblex/go-echo-demo/server/model"
	"github.com/Treblex/go-web-template/controller"
)

// NewAdController NewAdController
func NewAdController() *controller.Controller {
	return &controller.Controller{
		DB:    model.DB,
		Model: &model.Ad{},
	}
}

// NewAdGroupController NewAdGroupController
func NewAdGroupController() *controller.Controller {
	return &controller.Controller{
		DB:    model.DB,
		Model: &model.AdGroup{},
	}
}

// NewAdEventController NewAdEventController
func NewAdEventController() *controller.Controller {
	return &controller.Controller{
		DB:    model.DB,
		Model: &model.AdEvent{},
	}
}
