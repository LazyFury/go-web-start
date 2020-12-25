package controller

import (
	"github.com/Treblex/go-web-start/server/model"
	"github.com/Treblex/go-web-start/server/utils"
	"github.com/Treblex/go-web-template/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// NewArticleController NewArticleController
func NewArticleController() *ArticleController {
	return &ArticleController{
		Controller: &controller.Controller{
			DB:    model.DB,
			Model: &model.Articles{},
			Auth:  defaultAuth(),
		},
	}
}

// ArticleController ArticleController
type ArticleController struct {
	*controller.Controller
}

// Install Install
func (a *ArticleController) Install(g *gin.RouterGroup, path string) {
	controller.Install(g, a, path)
}

// ListPaging ListPagin
func (a *ArticleController) ListPaging(c *gin.Context) {
	cid := c.DefaultQuery("cid", "")
	a.DefaultListPaging(c, func(db *gorm.DB) *gorm.DB {
		query := map[string]interface{}{}
		if cid != "" {
			query["cate_id"] = cid
		}
		return db.Where(query)
	})
}

// ListAll ListAll
func (a *ArticleController) ListAll(c *gin.Context) {
	utils.Error(utils.NoRoute)
}

// NewArticleRecController NewArticleRecController
func NewArticleRecController() *ArticleRecController {
	return &ArticleRecController{
		Controller: &controller.Controller{
			DB:    model.DB,
			Model: &model.ArticlesRec{},
			Auth:  defaultAuth(),
		},
	}
}

// ArticleRecController ArticleController
type ArticleRecController struct {
	*controller.Controller
}

// Install Install
func (a *ArticleRecController) Install(g *gin.RouterGroup, path string) {
	controller.Install(g, a, path)
}

// Delete 删除
func (a *ArticleRecController) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.Error("请输入推荐位id")
	}

	rec := model.ArticlesRec{}
	if err := a.DB.GetObjectOrNotFound(&rec, map[string]interface{}{
		"id": id,
	}); err != nil {
		utils.Error("推荐位不存在")
	}

	articles := []model.Articles{}
	if err := a.DB.GetObjectsOrEmpty(&articles, nil, func(db *gorm.DB) *gorm.DB {
		return db.Where("id in (?)", rec.IDs)
	}).All(); err != nil {
		utils.Error(err)
	}
	if len(articles) > 0 {
		utils.Error("该推荐位下有文章，不可删除")
	}
	a.Controller.Delete(c)
}

// NewArticleCategoryController NewArticleCategoryController
func NewArticleCategoryController() *ArticleCategoryController {
	return &ArticleCategoryController{
		Controller: &controller.Controller{
			DB:    model.DB,
			Model: &model.ArticlesCate{},
			Auth:  defaultAuth(),
		},
	}
}

// ArticleCategoryController ArticleCategoryController
type ArticleCategoryController struct {
	*controller.Controller
}

// Install Install
func (a *ArticleCategoryController) Install(g *gin.RouterGroup, path string) {
	controller.Install(g, a, path)
}

// Delete 删除
func (a *ArticleCategoryController) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.Error("请输入分类id")
	}

	articles := []model.Articles{}
	if err := a.DB.GetObjectsOrEmpty(&articles, map[string]interface{}{
		"cate_id": id,
	}).All(); err != nil {
		utils.Error(err)
	}
	if len(articles) > 0 {
		utils.Error("该分类下有文章，不可删除")
	}

	tags := []model.ArticlesTag{}
	if err := a.DB.GetObjectsOrEmpty(&tags, map[string]interface{}{
		"cate_id": id,
	}).All(); err != nil {
		utils.Error(err)
	}
	if len(tags) > 0 {
		utils.Error("该分类下有标签，不可删除")
	}
	a.Controller.Delete(c)
}

// NewArticleTagController NewArticleTagController
func NewArticleTagController() *ArticleTagController {
	return &ArticleTagController{
		Controller: &controller.Controller{
			DB:    model.DB,
			Model: &model.ArticlesTag{},
			Auth:  defaultAuth(),
		},
	}
}

// ArticleTagController ArticleTagController
type ArticleTagController struct {
	*controller.Controller
}

// Install Install
func (a *ArticleTagController) Install(g *gin.RouterGroup, path string) {
	controller.Install(g, a, path)
}

// List List
func (a *ArticleTagController) List(c *gin.Context) {
	cid := c.DefaultQuery("cid", "")
	a.DefaultListPaging(c, func(db *gorm.DB) *gorm.DB {
		query := map[string]interface{}{}
		if cid != "" {
			query["cate_id"] = cid
		}
		return db.Where(query)
	})
}
