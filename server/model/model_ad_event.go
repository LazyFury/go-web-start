package model

import (
	"EK-Server/util"
	"strings"

	"github.com/labstack/echo"
)

// AdEvent banner事件
type AdEvent struct {
	BaseControll
	Event string `json:"event" gorm:"not null;unique_index;default:'no_event';comment:'banner事件,字符串，唯一'"`
}

// PointerList PointerList
func (a *AdEvent) PointerList() interface{} {
	return &[]struct {
		*AdEvent
		*EmptySystemFiled
	}{}
}

// Pointer Pointer
func (a *AdEvent) Pointer() interface{} {
	return &AdEvent{}
}

// TableName TableName
func (a *AdEvent) TableName() string {
	return TableName("ad_events")
}

// List 列表
func (a *AdEvent) List(c echo.Context) error {
	return util.JSONSuccess(c, a.BaseControll.ListWithOutPaging(nil), "")
}

// Add AdEventd
func (a *AdEvent) Add(c echo.Context) error {
	adEvent := &AdEvent{}

	if err := c.Bind(adEvent); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	adEvent.Event = strings.Trim(adEvent.Event, " ")
	if adEvent.Event == "" {
		return util.JSONErr(c, nil, "event定义不可空")
	}

	hasOne := a.BaseControll.HasOne(map[string]interface{}{
		"event": adEvent.Event,
	})
	if hasOne {
		return util.JSONErr(c, nil, "不可添加,已存在相同的事件")
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
