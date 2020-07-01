package model

import (
	"EK-Server/util/customtype"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type (
	//Goods 商品表
	Goods struct {
		Cid         uint             `json:"cid"`
		Title       string           `json:"title"`
		Description string           `gorm:"type:MEDIUMTEXT" json:"description"`
		Cover       string           `json:"cover"`
		Images      customtype.Array `gorm:"type:MEDIUMTEXT" json:"images" `
		Price       customtype.Money `gorm:"not null" json:"price"`
		Count       int              `json:"count"`
		BaseControll
	}

	//GoodsCate 商品分类表
	GoodsCate struct {
		BaseControll
		Name     string `json:"name"`
		Desc     string `json:"desc"`
		ParentID uint   `json:"parent_id"` //上级
		Cover    string `json:"cover"`     //封面
		Icon     string `json:"icon"`      //图标
		Level    int    `gorm:"DEFAULT:1" json:"level"`
	}
)

// PointerList 列表
func (goods *Goods) PointerList() interface{} {
	return &[]Goods{}
}

// Pointer 实例
func (goods *Goods) Pointer() interface{} {
	return &Goods{}
}

// TableName 表名
func (goods *Goods) TableName() string {
	return TableName("goods")
}

// Search 搜索
func (goods *Goods) Search(db *gorm.DB, key string) *gorm.DB {
	if key != "" {
		return db.Where("`title` like ?", "%"+key+"%").Or("`desc` like ?", "%"+key+"%")
	}
	return db
}

//List 文章列表
func (goods *Goods) List(c echo.Context) error {
	cid := c.QueryParam("cid")
	if cid != "" {
		cateID, err := strconv.Atoi(cid)
		if err == nil && cateID > 0 {
			return goods.BaseControll.GetList(c, &Goods{Cid: uint(cateID)})
		}
	}
	return goods.BaseControll.GetList(c, nil)
}
