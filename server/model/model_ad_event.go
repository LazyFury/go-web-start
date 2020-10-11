package model

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Treblex/go-echo-demo/server/util"
	"github.com/jinzhu/gorm"

	"github.com/labstack/echo/v4"
)

// AdEvent banner事件
type AdEvent struct {
	BaseControll
	Event string `json:"event" gorm:"not null;unique_index;default:'no_event';comment:'banner事件,字符串，唯一'"`
	Desc  string `json:"desc" gorm:""`
}

// PointerList PointerList
func (a *AdEvent) PointerList() interface{} {
	return &[]struct {
		*AdEvent
		// *EmptySystemFiled
		Count int `json:"count"`
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

// Joins Joins
func (a *AdEvent) Joins(db *gorm.DB) *gorm.DB {
	db = db.Select("*")
	ad := &Ad{}
	db = db.Joins(fmt.Sprintf("left join (select count(id) `count`,`event_id` from `%s` group by `event_id`) t1 on t1.`event_id`=`%s`.`id`", ad.TableName(), a.TableName()))
	return db
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

	if match, _ := regexp.MatchString("^[A-Za-z]+$", adEvent.Event); !match {
		return util.JSONErr(c, nil, "仅支持英文字符串")
	}

	hasOne := a.BaseControll.HasOne(map[string]interface{}{
		"event": adEvent.Event,
	})
	if hasOne {
		return util.JSONErr(c, nil, "不可添加,已存在相同的事件")
	}

	adEvent.Empty()
	return a.BaseControll.DoAdd(c, adEvent)
}

// Update Update
func (a *AdEvent) Update(c echo.Context) error {
	adEvent := &AdEvent{}

	if err := c.Bind(adEvent); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	adEvent.Empty()
	return a.BaseControll.DoUpdate(c, adEvent)
}
