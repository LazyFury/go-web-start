package model

import (
	"EK-Server/util"
	"strings"

	"github.com/labstack/echo"
)

// ArticlesCate 文章分类
type ArticlesCate struct {
	BaseControll
	Name string `json:"name"`
	Key  string `json:"key"`
	Desc string `json:"desc"`
}
type showArticleCate struct {
	*ArticlesCate
	*EmptySystemFiled
}

// PointerList 列表
func (a *ArticlesCate) PointerList() interface{} {
	return &[]showArticleCate{}
}

// Pointer 实例
func (a *ArticlesCate) Pointer() interface{} {
	return &showArticleCate{}
}

// Add 添加分类
func (a *ArticlesCate) Add(c echo.Context) error {
	cate := &ArticlesCate{}

	if err := c.Bind(cate); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	cate.Name = strings.Trim(cate.Name, " ")
	if cate.Name == "" {
		return util.JSONErr(c, nil, "分类名称不可空")
	}

	cate.Empty()
	return a.BaseControll.Add(c, cate)
}

// Update 添加分类
func (a *ArticlesCate) Update(c echo.Context) error {
	cate := &ArticlesCate{}

	if err := c.Bind(cate); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	cate.Empty()
	return a.BaseControll.Update(c, cate)
}
