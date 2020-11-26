package model

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/Treblex/go-echo-demo/server/utils/customtype"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type (
	// Articles 文章模型
	Articles struct {
		BaseControll
		Title          string           `json:"title" gorm:"not null"`
		Desc           string           `json:"desc"`
		Author         string           `json:"author" gorm:"DEFAULT:'佚名'"`
		Content        string           `json:"content" gorm:"type:text"`
		Email          string           `json:"email"`
		Cover          string           `json:"cover" gorm:"DEFAULT:'/static/images/default.jpg'"`
		Tag            customtype.Array `json:"tag" gorm:"type:varchar(255)"`
		LikeCount      int              `json:"like_count"`
		AlreadyLikedIt bool             `json:"already_liked_it" gorm:"-"` //判断当前用户是否点赞

		CateID int `json:"cate_id" gorm:"column:cate_id"`
	}
	selectArticle struct {
		*Articles
		CateName string `json:"cate_name"`
		// A        string `json:"content,omitempty"`
	}
	selectListArticle struct {
		*Articles
		CateName string `json:"cate_name"`
		A        string `json:"content,omitempty"`
	}
)

// NewArticle 新建文章类型
func NewArticle() *Articles {
	a := &Articles{}
	a.BaseControll.Model = a
	return a
}

// PointerList 列表
func (a *Articles) PointerList() interface{} {
	return &[]selectListArticle{}
}

// Pointer 实例
func (a *Articles) Pointer() interface{} {
	return &selectArticle{}
}

// TableName 表名
func (a *Articles) TableName() string {
	return TableName("articles")
}

// Search 搜索
func (a *Articles) Search(db *gorm.DB, key string) *gorm.DB {
	if key != "" {
		db = db.Where("`title` like ?", "%"+key+"%").Or("`desc` like ?", "%"+key+"%").Or("`tag` like ?", "%"+key+"%")
	}
	return db
}

// Joins 查询分类名
func (a *Articles) Joins(db *gorm.DB) *gorm.DB {
	// left join 需要手动拼接字段了,gorm默认都是slect tableName.*
	// db = db.Select("id,title,cate_id,cate_name,created_at,updated_at,author,email,cover,tag,`like`,`desc`")
	cate := &ArticlesCate{}
	db = db.Joins(fmt.Sprintf("left join (select name cate_name,id cid from `%s`) cate on cate.cid=`%s`.`cate_id`", cate.TableName(), a.TableName()))
	return db
}

// Result Result
func (a *Articles) Result(data interface{}, userID uint) interface{} {

	item, ok := reflect.ValueOf(data).Elem().Interface().(selectArticle)
	if ok {
		return item
	}

	return data
}

//List 文章列表
func (a *Articles) List(c *gin.Context) {
	cid := c.Query("cid")
	if cid != "" {
		cateID, err := strconv.Atoi(cid)
		if err == nil && cateID > 0 {
			a.BaseControll.GetList(c, &Articles{CateID: cateID})
			return
		}
	}
	a.BaseControll.GetList(c, nil)
}

// Detail 文章详情
func (a *Articles) Detail(c *gin.Context) {
	a.BaseControll.GetDetail(c, "文章不存在")
}

// Add 添加
func (a *Articles) Add(c *gin.Context) {
	article := &Articles{}

	if err := c.Bind(article); err != nil {
		panic("参数错误")
	}

	article.Title = strings.Trim(article.Title, " ")
	if article.Title == "" {
		panic("文章标题不可空")
	}

	if article.CateID == 0 {
		panic("请选择文章分类")
	}

	article.Empty()
	a.BaseControll.DoAdd(c, article)
}

// Update Update
func (a *Articles) Update(c *gin.Context) {
	article := &Articles{}

	if err := c.Bind(article); err != nil {
		panic("参数错误")
	}

	article.Empty()
	a.BaseControll.DoUpdate(c, article)
}
