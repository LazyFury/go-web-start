package model

import (
	"EK-Server/util"

	"github.com/jinzhu/gorm"
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
