package model

import (
	"EK-Server/util"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

//GoodsCate 商品分类表
type GoodsCate struct {
	BaseControll
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	ParentID uint   `gorm:"DEFAULT:0" json:"parent_id"` //上级
	Cover    string `json:"cover"`                      //封面
	Icon     string `json:"icon"`                       //图标
	Level    int    `gorm:"DEFAULT:1" json:"level"`
}

// 接口返回的列表 隐藏部分属性
type showGoodsCate struct {
	*GoodsCate
	*EmptySystemFiled
}

// 遍历所有子分类的结构体
type catelist struct {
	*GoodsCate
	*EmptySystemFiled
	Tmenu []catelist `json:"tmenu"`
}

// PointerList 列表
func (cate *GoodsCate) PointerList() interface{} {
	return &[]showGoodsCate{}
}

// Pointer 实例
func (cate *GoodsCate) Pointer() interface{} {
	return &showGoodsCate{}
}

// TableName 表名
func (cate *GoodsCate) TableName() string {
	return TableName("good-cates")
}

// List 列表
func (cate *GoodsCate) List(c echo.Context) error {
	db := DB
	list := []catelist{}

	db.Table(cate.TableName()).Where(map[string]interface{}{"parent_id": 0, "level": 1}).Find(&list)

	for i, item := range list {
		list[i].Tmenu = cate.getCateTmenu(&item, db)
	}
	return util.JSONSuccess(c, list, "获取成功")
}

// 循环获取自分类
func (cate *GoodsCate) getCateTmenu(item *catelist, db *gorm.DB) (tmenu []catelist) {
	parentID := item.ID
	tmenu = []catelist{}
	db.Table(cate.TableName()).Where(&GoodsCate{ParentID: parentID}).Find(&tmenu)
	if len(tmenu) > 0 {
		for i, menuItem := range tmenu {
			tmenu[i].Tmenu = cate.getCateTmenu(&menuItem, db)
		}
	}
	return
}

// Add 添加
func (cate *GoodsCate) Add(c echo.Context) error {
	db := DB
	_cate := &GoodsCate{}

	if err := c.Bind(_cate); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	_cate.Name = strings.Trim(_cate.Name, " ")
	if _cate.Name == "" {
		return util.JSONErr(c, nil, "分类名称不可空")
	}

	_cate.Level = 1 //禁止手动设置level
	_cate.Empty()
	// 查询分类是否存在 parentID为空时是一级分类
	if _cate.ParentID > 0 {
		cateParent := &GoodsCate{BaseControll: BaseControll{ID: uint(_cate.ParentID)}}
		empty := db.First(cateParent).RecordNotFound()
		// fmt.Println(empty)
		if empty {
			return util.JSONErr(c, nil, "上级分类不存在")
		}
		_cate.Level = cateParent.Level + 1
	}
	// 限制层级
	if _cate.Level > 3 {
		return util.JSONErr(c, nil, "最多3级分类，不可添加子分类")
	}

	// 禁止同名
	if repeat := db.Where(&GoodsCate{Name: _cate.Name}).Find(&GoodsCate{}).RecordNotFound(); !repeat {
		return util.JSONErr(c, nil, "已存在相同分类")
	}

	return cate.BaseControll.Add(c, _cate)
}

// Update 更新
func (cate *GoodsCate) Update(c echo.Context) error {
	_cate := &GoodsCate{}

	if err := c.Bind(_cate); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	_cate.ParentID = 0
	_cate.Level = 0
	_cate.Empty()
	return cate.BaseControll.Update(c, _cate)
}

// Delete 删除
func (cate *GoodsCate) Delete(c echo.Context) error {
	db := DB
	id := c.Param("id")
	if id == "" {
		return util.JSONErr(c, nil, "参数错误")
	}

	if hasGoods := db.Model(cate.Pointer()).Where(map[string]interface{}{"cid": id}).Find(cate.Pointer()).RowsAffected; hasGoods > 0 {
		return util.JSONErr(c, nil, "分类下有商品，无法删除")
	}
	if hasCates := db.Model(cate.Pointer()).Where(map[string]interface{}{"parent_id": id}).Find(cate.Pointer()).RowsAffected; hasCates > 0 {
		return util.JSONErr(c, nil, "分类下有其他分类，无法删除")
	}

	return cate.BaseControll.Delete(c)
}
