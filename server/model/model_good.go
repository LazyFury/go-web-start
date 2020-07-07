package model

import (
	"EK-Server/util"
	"EK-Server/util/customtype"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type (
	//Goods 商品表
	Goods struct {
		BaseControll
		Cid         uint             `json:"cid"`
		Title       string           `json:"title"`
		Description string           `gorm:"type:MEDIUMTEXT" json:"description"`
		Cover       string           `json:"cover"`
		Images      customtype.Array `gorm:"type:MEDIUMTEXT" json:"images" `
		Price       customtype.Money `gorm:"not null" json:"price"`
		StockCount  int              `json:"stock_count"`
		OnSale      bool             `json:"on_sale" gorm:"default:1;comment:'是否在售，上下架功能'"`
	}
)

// PointerList 列表
func (g *Goods) PointerList() interface{} {
	return &[]Goods{}
}

// Pointer 实例
func (g *Goods) Pointer() interface{} {
	return &Goods{}
}

// Search 搜索
func (g *Goods) Search(db *gorm.DB, key string) *gorm.DB {
	if key != "" {
		return db.Where("`title` like ?", "%"+key+"%").Or("`description` like ?", "%"+key+"%")
	}
	return db
}

// TableName 表名
func (g *Goods) TableName() string {
	return TableName("goods")
}

//List 文章列表
func (g *Goods) List(c echo.Context) error {
	cid := c.QueryParam("cid")
	if cid != "" {
		cateID, err := strconv.Atoi(cid)
		if err == nil && cateID > 0 {
			return g.BaseControll.GetList(c, &Goods{Cid: uint(cateID)})
		}
	}
	return g.BaseControll.GetList(c, nil)
}

// Add 添加商品
func (g *Goods) Add(c echo.Context) error {
	good := &Goods{}

	if err := c.Bind(good); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	if good.Cid == 0 {
		return util.JSONErr(c, nil, "请选择商品分类")
	}
	good.Title = strings.Trim(good.Title, " ")
	if good.Title == "" {
		return util.JSONErr(c, nil, "商品标题不可空")
	}

	var zeroMoney customtype.Money
	if good.Price == zeroMoney {
		return util.JSONErr(c, nil, "请填写商品价格")
	}

	good.Empty()
	return g.BaseControll.Add(c, good)
}

// Update 添加商品
func (g *Goods) Update(c echo.Context) error {
	good := &Goods{}

	if err := c.Bind(good); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	good.Empty()
	return g.BaseControll.Update(c, good)
}
