package cate

import (
	"EK-Server/config"
	"EK-Server/model"
	"EK-Server/util"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// Init Init
func Init(g *echo.Group) {
	baseURL := "/cate"
	cate := g.Group(baseURL)
	cate.POST("/add", add)

	cate.GET("/list", list)
}

type catelist struct {
	model.GoodsCateList
	Tmenu []catelist `json:"tmenu"`
}

func list(c echo.Context) error {
	tableName := config.Global.TablePrefix + "_goods_cates"

	list := []catelist{}

	db := model.DB

	db.Table(tableName).Where(map[string]interface{}{"parent_id": 0, "level": 1}).Find(&list)

	// fmt.Println(list)
	for i, item := range list {
		list[i].Tmenu = getTmenu(&item, db, tableName)
	}
	return util.JSONSuccess(c, list, "获取成功")
}

// 循环获取自分类
func getTmenu(item *catelist, db *gorm.DB, tableName string) (tmenu []catelist) {
	parentID := item.ID
	tmenu = []catelist{}
	db.Table(tableName).Where(&model.GoodsCate{ParentID: parentID}).Find(&tmenu)
	if len(tmenu) > 0 {
		for i, menuItem := range tmenu {
			tmenu[i].Tmenu = getTmenu(&menuItem, db, tableName)
		}
	}
	return
}

func add(c echo.Context) error {
	db := model.DB
	cate := &model.GoodsCate{}
	// 绑定json
	if err := c.Bind(cate); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}
	cate.Name = strings.Trim(cate.Name, " ") //过滤空格
	if cate.Name == "" {
		return util.JSONErr(c, nil, "分类名称不可空")
	}
	// 查询分类是否存在 parentID为空时是一级分类
	if cate.ParentID > 0 {
		parent := &model.GoodsCate{Model: gorm.Model{ID: uint(cate.ParentID)}}
		empty := db.First(parent).RecordNotFound()
		// fmt.Println(empty)
		if empty {
			return util.JSONErr(c, nil, "上级分类不存在")
		}
		cate.Level = parent.Level + 1
	}
	// 限制层级
	if cate.Level > 4 {
		return util.JSONErr(c, nil, "已经是底层分类，不可添加子分类")
	}
	// 禁止同名
	if repeat := db.Where(&model.GoodsCate{Name: cate.Name}).First(&model.GoodsCate{}).RecordNotFound(); !repeat {
		return util.JSONErr(c, nil, "已存在相同分类")
	}
	db.NewRecord(cate)
	row := db.Create(cate)
	if row.RowsAffected >= 1 {
		return util.JSONSuccess(c, nil, "添加成功")
	}

	return util.JSONSuccess(c, nil, "添加失败")
}
