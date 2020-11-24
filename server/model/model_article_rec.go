package model

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Treblex/go-echo-demo/server/util"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// ArticlesRec 文章分类
type ArticlesRec struct {
	BaseControll
	Name string `json:"name"`
	Key  string `json:"key"`
	IDs  string `json:"article_ids"`
	Desc string `json:"desc"`
}
type showArticleRec struct {
	*ArticlesRec
	// *EmptySystemFiled
	List  []selectArticle `json:"list"`
	Count int             `json:"count"`
}

// NewArticleRec 推荐文章
func NewArticleRec() *ArticlesRec {
	rec := &ArticlesRec{}
	rec.BaseControll.Model = rec
	return rec
}

// PointerList 列表
func (a *ArticlesRec) PointerList() interface{} {
	return &[]showArticleRec{}
}

// Pointer 实例
func (a *ArticlesRec) Pointer() interface{} {
	return &showArticleRec{}
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
	articles := []selectArticle{}

	db := DB.Table(article.TableName())
	row := db.Where("id IN (?)", ids).Find(&articles)
	if row.Error == nil && len(articles) > 0 {
		item.List = articles
	}

	l := len(articles)
	item.Count = l
	if l == 0 {
		item.List = []selectArticle{}
	}

	fmt.Println(item)
}

// Result 处理结构
func (a *ArticlesRec) Result(data interface{}, userID uint) interface{} {
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
func (a *ArticlesRec) List(c echo.Context) error {
	list := a.BaseControll.ListWithOutPaging(nil)
	userID, _ := c.Get("userId").(float64)
	list = a.Result(list, uint(userID))
	return util.JSONSuccess(c, list, "")
}

// Add 添加分类
func (a *ArticlesRec) Add(c echo.Context) error {
	rec := &ArticlesRec{}

	if err := c.Bind(rec); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	rec.Name = strings.Trim(rec.Name, " ")
	if rec.Name == "" {
		return util.JSONErr(c, nil, "名称不可空")
	}

	rec.Empty()
	return a.BaseControll.DoAdd(c, rec)
}

// Update 添加分类
func (a *ArticlesRec) Update(c echo.Context) error {
	rec := &ArticlesRec{}

	if err := c.Bind(rec); err != nil {
		return util.JSONErr(c, err, "参数错误")
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
				return util.JSONErr(c, row.Error, "")
			}

			if len(articles) <= 0 {
				return util.JSONErr(c, nil, "选择了无效的文章")
			}
			ids = []string{}
			for _, id := range articles {
				ids = append(ids, fmt.Sprintf("%d", id.ID))
			}

			rec.IDs = strings.Join(ids, ",")
		}
	}

	rec.Empty()
	return a.BaseControll.DoUpdate(c, rec)
}

// Delete 删除
func (a *ArticlesRec) Delete(c echo.Context) error {
	db := DB
	id := c.Param("id")
	if id == "" {
		return util.JSONErr(c, nil, "参数错误")
	}
	article := &Articles{}
	if hasArticle := db.Model(article).Where(map[string]interface{}{"cate_id": id}).Find(article).RowsAffected; hasArticle > 0 {
		return util.JSONErr(c, nil, "该推荐位下还有文章，不能删除")
	}
	return a.BaseControll.Delete(c)
}
