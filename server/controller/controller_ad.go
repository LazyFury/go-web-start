package controller

import "github.com/Treblex/go-echo-demo/server/model"

// NewAdController NewAdController
func NewAdController() *Controller {
	return &Controller{
		DB:    model.DB,
		Model: &model.Ad{},
	}
}

// NewAdGroupController NewAdGroupController
func NewAdGroupController() *Controller {
	return &Controller{
		DB:    model.DB,
		Model: &model.AdGroup{},
	}
}

// NewAdEventController NewAdEventController
func NewAdEventController() *Controller {
	return &Controller{
		DB:    model.DB,
		Model: &model.AdEvent{},
	}
}
