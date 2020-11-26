package controller

import (
	"github.com/Treblex/go-echo-demo/server/model"
	"github.com/gin-gonic/gin"
)

// ArticleController ArticleController
type ArticleController struct {
	*Controller
}

// ListAll ListAll
func (a *ArticleController) ListAll(c *gin.Context) {
	a.Controller.obj = &[]model.Articles{}
	a.Controller.ListAll(c)
}

// Detail Detail
func (a *ArticleController) Detail(c *gin.Context) {
	a.Controller.Detail(c)
}
