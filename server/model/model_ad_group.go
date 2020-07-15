package model

import (
	"EK-Server/util"

	"github.com/labstack/echo"
)

const (
	// Sigle 单图
	Sigle int = 1
	// Multi 多图
	Multi int = 2
)

// AdGroup 广告位
type AdGroup struct {
	Type int    `json:"type" gorm:"comment:'1单图,2多图'"`
	Name string `json:"name"`
	BaseControll
}

// PointerList PointerList
func (a *AdGroup) PointerList() interface{} {
	return &[]AdGroup{}
}

// Pointer Pointer
func (a *AdGroup) Pointer() interface{} {
	return &AdGroup{}
}

// TableName TableName
func (a *AdGroup) TableName() string {
	return TableName("ad_groups")
}

// Add AdGroupd
func (a *AdGroup) Add(c echo.Context) error {
	adGroup := &AdGroup{}

	if err := c.Bind(adGroup); err != nil {
		return util.JSONErr(c, err, "参数错误")
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
