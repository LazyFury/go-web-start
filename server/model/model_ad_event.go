package model

import (
	"EK-Server/util"

	"github.com/labstack/echo"
)

// AdEvent banner事件
type AdEvent struct {
	Event string `json:"event" gorm:"not null;unique_index;comment:'banner事件,字符串，唯一'"`
	BaseControll
}

// PointerList PointerList
func (a *AdEvent) PointerList() interface{} {
	return &[]AdEvent{}
}

// Pointer Pointer
func (a *AdEvent) Pointer() interface{} {
	return &AdEvent{}
}

// TableName TableName
func (a *AdEvent) TableName() string {
	return TableName("AdEvents")
}

// AdEventd AdEventd
func (a *AdEvent) AdEventd(c echo.Context) error {
	adEvent := &AdEvent{}

	if err := c.Bind(adEvent); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	adEvent.Empty()
	return a.BaseControll.Add(c, adEvent)
}

// Update Update
func (a *AdEvent) Update(c echo.Context) error {
	adEvent := &AdEvent{}

	if err := c.Bind(adEvent); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	adEvent.Empty()
	return a.BaseControll.Update(c, adEvent)
}
