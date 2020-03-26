package model

import "github.com/jinzhu/gorm"

type (
	//Goods 商品表
	Goods struct {
		gorm.Model
		Title  string `json:"title"`
		Desc   string `json:"desc"`
		Cover  string `json:"cover"`
		Images string `json:"images"`
		Price  string `json:"price"`
	}

	// GoodsList 商品展示表
	GoodsList struct {
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Desc   string `json:"desc"`
		Cover  string `json:"cover"`
		Images string `json:"images"`
		Price  string `json:"price"`
	}

	//GoodsCate 商品分类表
	GoodsCate struct {
		gorm.Model
	}
)
