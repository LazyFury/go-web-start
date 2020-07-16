package model

import (
	"EK-Server/util"
	"strings"

	"github.com/labstack/echo"
)

// AdGroup 广告位
type AdGroup struct {
	BaseControll
	IsSigle bool   `json:"is_sigle" gorm:"default:false;comment:'true单图,false多图'"`
	Name    string `json:"name" gorm:"unique;not null"`
	Desc    string `json:"desc" gorm:"type:text;comment:'描述';"`
}

// PointerList PointerList
func (a *AdGroup) PointerList() interface{} {
	return &[]struct {
		*AdGroup
		*EmptySystemFiled
	}{}
}

// Pointer Pointer
func (a *AdGroup) Pointer() interface{} {
	return &AdGroup{}
}

// TableName TableName
func (a *AdGroup) TableName() string {
	return TableName("ad_groups")
}

// Detail 分组详情
func (a *AdGroup) Detail(c echo.Context) error {
	db := DB
	id := c.Param("id")
	if id == "" {
		return util.JSONErr(c, nil, "参数错误")
	}

	group := &AdGroup{}
	if db.Model(group).Where(map[string]interface{}{
		"id": id,
	}).First(group).RecordNotFound() {
		return util.JSONErr(c, nil, "广告位不存在")
	}

	ad := &Ad{}
	ad.BaseControll.Model = ad

	list := ad.BaseControll.ListWithOutPaging(map[string]interface{}{
		"group_id": id,
	})
	result := map[string]interface{}{
		"name":     group.Name,
		"is_sigle": group.IsSigle,
		"list":     list,
	}

	return util.JSONSuccess(c, result, "")
}

// Add AdGroupd
func (a *AdGroup) Add(c echo.Context) error {
	adGroup := &AdGroup{}

	if err := c.Bind(adGroup); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	adGroup.Name = strings.Trim(adGroup.Name, " ")
	if adGroup.Name == "" {
		return util.JSONErr(c, nil, "分组标题不可空")
	}

	if a.BaseControll.HasOne(map[string]interface{}{
		"name": adGroup.Name,
	}) {
		return util.JSONErr(c, nil, "已存在相同的分类")
	}

	adGroup.Empty()
	return a.BaseControll.Add(c, adGroup)
}

// Update Update
func (a *AdGroup) Update(c echo.Context) error {
	adGroup := &AdGroup{}

	if err := c.Bind(adGroup); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	adGroup.Empty()
	return a.BaseControll.Update(c, adGroup)
}
