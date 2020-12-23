package model

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/Treblex/go-echo-demo/server/utils"
	"github.com/Treblex/go-web-template/xmodel"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

// ArticlesRec 文章分类
type ArticlesRec struct {
	BaseControll
	Name string `json:"name" gorm:"unique;not null"`
	Key  string `json:"key"`
	IDs  string `json:"article_ids"`
	Desc string `json:"desc"`
}
type showArticleRec struct {
	ArticlesRec
	// *EmptySystemFiled
	List  []Articles `json:"list" gorm:"-"`
	Count int        `json:"count" gorm:"-"`
}

var _ xmodel.Controller = &ArticlesRec{}

// Object Object
func (a *ArticlesRec) Object() interface{} {
	return &showArticleRec{}
}

// Objects Object
func (a *ArticlesRec) Objects() interface{} {
	return &[]showArticleRec{}
}

// Validator Validator
func (a *ArticlesRec) Validator() error {
	a.Name = strings.Trim(a.Name, " ")
	if a.Name == "" {
		return errors.New("请输入推荐位名称")
	}
	return nil
}

// TableName 表名
func (a *ArticlesRec) TableName() string {
	return TableName("article_rec")
}

// Joins  查询相关文章数据
func (a *ArticlesRec) Joins(db *gorm.DB) *gorm.DB {
	return db
}

func (a *ArticlesRec) getArticle(item *showArticleRec) {
	ids := strings.Split(item.IDs, ",")
	article := &Articles{}
	articles := []Articles{}

	row := DB.Table(article.TableName()).Where("id IN (?)", ids).Find(&articles)
	if row.Error == nil && len(articles) > 0 {
		item.List = articles
	}

	l := len(articles)
	item.Count = l
	if l == 0 {
		item.List = []Articles{}
	}

	// fmt.Println(item)
}

// Result 处理结构
func (a *ArticlesRec) Result(data interface{}) interface{} {
	interf := reflect.ValueOf(data).Elem().Interface()
	arr, ok := interf.([]showArticleRec)
	if ok {
		for i := range arr {
			a.getArticle(&arr[i])
		}
		return arr
	}

	item, ok := interf.(showArticleRec)
	if ok {
		a.getArticle(&item)
		return item
	}
	return data
}

// List 分页
func (a *ArticlesRec) List(c *gin.Context) {
	res := a.Result(a.ListWithOutPaging(nil))
	c.JSON(http.StatusOK, utils.JSONSuccess("", res))
}

// Add 添加分类
func (a *ArticlesRec) Add(c *gin.Context) {
	rec := &ArticlesRec{}

	if err := c.Bind(rec); err != nil {
		utils.Error("参数错误")
	}

	rec.Name = strings.Trim(rec.Name, " ")
	if rec.Name == "" {
		utils.Error("名称不可空")
	}

	rec.Empty()
	a.BaseControll.DoAdd(c, rec)
}

// Update 添加分类
func (a *ArticlesRec) Update(c *gin.Context) {
	rec := &ArticlesRec{}

	if err := c.Bind(rec); err != nil {
		utils.Error("参数错误")
	}

	// 为0 清空选择的文章，不为0时需要验证文章可用性
	if rec.IDs != "0" {
		ids := strings.Split(rec.IDs, ",")
		if ids[0] == "" {
			ids = ids[1:]
		}
		if len(ids) > 0 {
			db := DB
			article := &Articles{}
			articles := []Articles{}
			row := db.Table(article.TableName()).Where("id IN (?)", ids).Find(&articles)
			if row.Error != nil {
				utils.Error(row.Error)
			}

			if len(articles) <= 0 {
				utils.Error("选择了无效的文章")
			}
			ids = []string{}
			for _, id := range articles {
				ids = append(ids, fmt.Sprintf("%d", id.ID))
			}

			rec.IDs = strings.Join(ids, ",")
		}
	}

	rec.Empty()
	a.BaseControll.DoUpdate(c, rec)
}

// Delete 删除
func (a *ArticlesRec) Delete(c *gin.Context) {
	db := DB
	id := c.Param("id")
	if id == "" {
		utils.Error("参数错误")
	}
	article := &Articles{}
	if hasArticle := db.Model(article).Where(map[string]interface{}{"cate_id": id}).Find(article).RowsAffected; hasArticle > 0 {
		utils.Error("该推荐位下还有文章，不能删除")
	}
	a.BaseControll.Delete(c)
}
