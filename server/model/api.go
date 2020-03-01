package model

import (
	"github.com/jinzhu/gorm"
)

// API API接口列表
type API struct {
	gorm.Model
	Name string `json:"name"`
	Data string `json:"data" gorm:"type:longtext;"`
	Cid  string `json:"cid"`
}

//添加API
func (a *API) Add() (msg string, err error) {
	db := DB
	//db.NewRecord(c)
	if err = db.Create(a).Error; err != nil {
		msg = "添加失败"
		return
	}
	//db.NewRecord(c)
	msg = "添加成功"
	return
}

// APICate api分类
type APICate struct {
	gorm.Model
	Name string `json:"name"`
	Desc string `json:"desc"`
}

// Add 添加分类
func (c *APICate) Add() (msg string, err error) {
	db := DB
	//db.NewRecord(c)
	if err = db.Create(c).Error; err != nil {
		msg = "添加失败"
		return
	}
	//db.NewRecord(c)
	msg = "添加成功"
	return
}

//获取API分类
func (c *APICate) GetAll() (result []APICate, msg string, err error) {
	db := DB
	if db.Find(&result).Error != nil {
		msg = "查询失败"
		return
	}
	msg = "获取成功"
	return
}

func (c *APICate) Save() (msg string, err error) {
	db := DB

	if err = db.Save(&c).Error; err != nil {
		msg = "保存失败"
		return
	}
	msg = "保存成功"
	return
}
