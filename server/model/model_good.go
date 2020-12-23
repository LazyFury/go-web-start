package model

import (
	"strconv"
	"strings"

	"github.com/Treblex/go-echo-demo/server/utils"
	"github.com/Treblex/go-echo-demo/server/utils/customtype"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
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
func (g *Goods) List(c *gin.Context) {
	cid := c.Query("cid")
	if cid != "" {
		cateID, err := strconv.Atoi(cid)
		if err == nil && cateID > 0 {
			g.BaseControll.GetList(c, &Goods{Cid: uint(cateID)})
			return
		}
	}
	g.BaseControll.GetList(c, nil)
}

// Detail 商品详情
func (g *Goods) Detail(c *gin.Context) {
	g.BaseControll.GetDetail(c, "商品不存在")
}

// Add 添加商品
func (g *Goods) Add(c *gin.Context) {
	good := &Goods{}

	if err := c.Bind(good); err != nil {
		utils.Error("参数错误")
	}

	if good.Cid == 0 {
		utils.Error("请选择商品分类")
	}
	good.Title = strings.Trim(good.Title, " ")
	if good.Title == "" {
		utils.Error("商品标题不可空")
	}

	var zeroMoney customtype.Money
	if good.Price == zeroMoney {
		utils.Error("请填写商品价格")
	}

	good.Empty()
	g.BaseControll.DoAdd(c, good)
}

// Update 添加商品
func (g *Goods) Update(c *gin.Context) {
	good := &Goods{}

	if err := c.Bind(good); err != nil {
		utils.Error("参数错误")
	}

	good.Empty()
	g.BaseControll.DoUpdate(c, good)
}
