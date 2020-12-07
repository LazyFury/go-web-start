package controller

import (
	"github.com/Treblex/go-echo-demo/server/model"
	"github.com/Treblex/go-echo-demo/server/utils"
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
		query := map[string]interface{}{}
		if cid != "" {
			query["cate_id"] = cid
		}
		return db.Where(query)
	})
}

// ListAll ListAll
func (a *ArticleController) ListAll(c *gin.Context) {
	panic(utils.NoRoute)
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

// Install Install
func (a *ArticleRecController) Install(g *gin.RouterGroup, path string) {
	Install(g, a, path)
}

// Delete 删除
func (a *ArticleRecController) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		panic("请输入推荐位id")
	}

	rec := model.ArticlesRec{}
	if err := a.DB.GetObjectOrNotFound(&rec, map[string]interface{}{
		"id": id,
	}); err != nil {
		panic("推荐位不存在")
	}

	articles := []model.Articles{}
	if err := a.DB.GetObjectsOrEmpty(&articles, nil, func(db *gorm.DB) *gorm.DB {
		return db.Where("id in (?)", rec.IDs)
	}).All(); err != nil {
		panic(err)
	}
	if len(articles) > 0 {
		panic("该推荐位下有文章，不可删除")
	}
	a.Controller.Delete(c)
}

// NewArticleCategoryController NewArticleCategoryController
func NewArticleCategoryController() *ArticleCategoryController {
	return &ArticleCategoryController{
		Controller: &Controller{
			DB:    model.DB,
			Model: &model.ArticlesCate{},
		},
	}
}

// ArticleCategoryController ArticleCategoryController
type ArticleCategoryController struct {
	*Controller
}

// Install Install
func (a *ArticleCategoryController) Install(g *gin.RouterGroup, path string) {
	Install(g, a, path)
}

// Delete 删除
func (a *ArticleCategoryController) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		panic("请输入分类id")
	}

	articles := []model.Articles{}
	if err := a.DB.GetObjectsOrEmpty(&articles, map[string]interface{}{
		"cate_id": id,
	}).All(); err != nil {
		panic(err)
	}
	if len(articles) > 0 {
		panic("该分类下有文章，不可删除")
	}

	tags := []model.ArticlesTag{}
	if err := a.DB.GetObjectsOrEmpty(&tags, map[string]interface{}{
		"cate_id": id,
	}).All(); err != nil {
		panic(err)
	}
	if len(tags) > 0 {
		panic("该分类下有标签，不可删除")
	}
	a.Controller.Delete(c)
}

// NewArticleTagController NewArticleTagController
func NewArticleTagController() *ArticleTagController {
	return &ArticleTagController{
		Controller: &Controller{
			DB:    model.DB,
			Model: &model.ArticlesTag{},
		},
	}
}

// ArticleTagController ArticleTagController
type ArticleTagController struct {
	*Controller
}

// Install Install
func (a *ArticleTagController) Install(g *gin.RouterGroup, path string) {
	Install(g, a, path)
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
