package controller

import (
	"github.com/Treblex/go-echo-demo/server/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// NewArticleController NewArticleController
func NewArticleController() *ArticleController {
	return &ArticleController{
		Controller: &Controller{
			DB:     model.DB,
			Model:  &model.Articles{},
			isUser: false,
		},
	}
}

// ArticleController ArticleController
type ArticleController struct {
	*Controller
}

// Install Install
func (a *ArticleController) Install(g *gin.RouterGroup, path string) {
	Install(g, a, path)
}

// ListPaging ListPagin
func (a *ArticleController) ListPaging(c *gin.Context) {
	cid := c.DefaultQuery("cid", "")
	a.DefaultListPaging(c, func(db *gorm.DB) *gorm.DB {
		if cid != "" {
			db = db.Where(map[string]interface{}{
				"cate_id": cid,
			})
		}
		return db
	})
}

// NewArticleRecController NewArticleRecController
func NewArticleRecController() *ArticleRecController {
	return &ArticleRecController{
		Controller: &Controller{
			DB:    model.DB,
			Model: &model.ArticlesRec{},
		},
	}
}

// ArticleRecController ArticleController
type ArticleRecController struct {
	*Controller
}

var _ BaseInterface = &ArticleRecController{}

// Detail Detail
func (a *ArticleRecController) Detail(c *gin.Context) {
	a.Controller.Detail(c)
}
