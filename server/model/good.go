package model

import (
	"EK-Server/util"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type (
	//Goods 商品表
	Goods struct {
		gorm.Model
		Cid         int        `json:"cid"`
		Title       string     `json:"title"`
		Description string     `gorm:"type:MEDIUMTEXT" json:"description"`
		Cover       string     `json:"cover"`
		Images      util.Array `gorm:"type:MEDIUMTEXT" json:"images" `
		Price       util.Money `gorm:"not null" json:"price"`
		Count       int        `json:"count"`
	}

	// GoodsList 商品展示表
	GoodsList struct {
		ID          int        `json:"id"`
		Cid         int        `json:"cid"`
		Title       string     `json:"title"`
		Description string     `json:"description"`
		Cover       string     `json:"cover"`
		Images      util.Array `json:"images"`
		Price       util.Money `json:"price"`
		Count       int        `json:"count"`
	}

	//GoodsCate 商品分类表
	GoodsCate struct {
		gorm.Model
		Name     string `json:"name"`
		Desc     string `json:"desc"`
		ParentID int    `json:"parent_id"` //上级
		Cover    string `json:"cover"`     //封面
		Icon     string `json:"icon"`      //图标
		Level    int    `gorm:"DEFAULT:1" json:"level"`
	}

	// GoodsCateList  GoodsCateList
	GoodsCateList struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Desc     string `json:"desc"`
		ParentID int    `json:"parent_id"` //上级
		Cover    string `json:"cover"`     //封面
		Icon     string `json:"icon"`      //图标
		Level    int    `json:"level"`
	}
)

//List 商品列表
func (g Goods) List(c echo.Context) error {
	type Param struct {
		PageParams
		Cid int `json:"cid"`
	}
	page := Param{PageParams: PageParams{Page: 1, Limit: 10, Order: "id_desc"}}

	if err := c.Bind(&page); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	fmt.Printf("post json 参数：%v", page)

	if page.Order != "" {
		page.Order = strings.ReplaceAll(page.Order, "_", " ")
		// page.Order = strings.ReplaceAll(page.Order, ",", " ")
	}
	where := &Goods{}
	if page.Cid > 0 {
		where = &Goods{Cid: page.Cid}
	}

	return util.JSONSuccess(c, DataBaselimit(page.Limit, page.Page, where, &[]GoodsList{}, "goods",
		page.Order), "获取成功")
}
