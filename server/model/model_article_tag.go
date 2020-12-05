package model

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/Treblex/go-echo-demo/server/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ArticlesTag 文章Tag
type ArticlesTag struct {
	BaseControll
	Val    string `json:"val" gorm:"not null;index;unique"`
	CateID uint   `json:"cate_id" gorm:"conment:'分类id，暂时公用文章分类'"`
}

type showArticlesTag struct {
	ArticlesTag

	CateName string `json:"cate_name" gorm:"->"`
	Count    int64  `json:"count" gorm:"->"`
}

var _ Controller = &ArticlesTag{}

// Validator Validator
func (a *ArticlesTag) Validator() error {
	if a.CateID == 0 {
		panic("请选择分类")
	}
	a.Val = strings.Trim(a.Val, " ")
	if a.Val == "" {
		panic("请输入标签名称")
	}
	return nil
}

// Object Object
func (a *ArticlesTag) Object() interface{} {
	return &ArticlesTag{}
}

// Objects Objects
func (a *ArticlesTag) Objects() interface{} {
	return &[]showArticlesTag{}
}

// TableName TableName
func (a *ArticlesTag) TableName() string {
	return TableName("article_tags")
}

// Joins 链接
func (a *ArticlesTag) Joins(db *gorm.DB) *gorm.DB {
	db = db.Select("*")

	articlesCate := &ArticlesCate{}
	name := a.TableName()
	cName := articlesCate.TableName()

	db = db.Joins(fmt.Sprintf("left join (select name cate_name,id cid from %s) t1 on t1.`cid`=`%s`.`cate_id`", cName, name))

	return db
}

// TODO:统计文章数量，后期准备优化为定时更新 或者手动更新 或重新设计数据表
// 想要使用左连接查询解决 但是没有找到方案
func (a *ArticlesTag) countArticles(tag *showArticlesTag) {
	var count int64
	db := DB
	article := &Articles{}
	db.Table(article.TableName()).Where("tag like ? AND `deleted_at` IS NULL", "%"+tag.Val+"%").Count(&count)
	tag.Count = count
}

// Result Result
func (a *ArticlesTag) Result(data interface{}) interface{} {
	// TODO:反射获取Interface之前需要判断是否是指针类型
	arr, ok := reflect.ValueOf(data).Elem().Interface().([]showArticlesTag)
	if ok {
		for i := range arr {
			a.countArticles(&arr[i])
		}
		return arr
	}
	return data
}

// List 列表
func (a *ArticlesTag) List(c *gin.Context) {
	where := map[string]interface{}{}
	cid := c.Query("cate_id")
	if cid == "" {
		where = nil
	} else {
		where["cate_id"] = cid
	}

	list := a.ListWithOutPaging(where)
	list = a.Result(list)
	c.JSON(http.StatusOK, utils.JSONSuccess("", list))
}
