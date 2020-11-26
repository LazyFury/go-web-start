package model

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/Treblex/go-echo-demo/server/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
func (a *AdEvent) List(c *gin.Context) {
	c.JSON(http.StatusOK, utils.JSONSuccess("", a.BaseControll.ListWithOutPaging(nil)))
}

// Add AdEventd
func (a *AdEvent) Add(c *gin.Context) {
	adEvent := &AdEvent{}

	if err := c.Bind(adEvent); err != nil {
		panic("参数错误")
	}

	adEvent.Event = strings.Trim(adEvent.Event, " ")
	if adEvent.Event == "" {
		panic("event定义不可空")
	}

	if match, _ := regexp.MatchString("^[A-Za-z]+$", adEvent.Event); !match {
		panic("仅支持英文字符串")
	}

	hasOne := a.BaseControll.HasOne(map[string]interface{}{
		"event": adEvent.Event,
	})
	if hasOne {
		panic("不可添加,已存在相同的事件")
	}

	adEvent.Empty()
	a.BaseControll.DoAdd(c, adEvent)
}

// Update Update
func (a *AdEvent) Update(c *gin.Context) {
	adEvent := &AdEvent{}

	if err := c.Bind(adEvent); err != nil {
		panic("参数错误")
	}

	adEvent.Empty()
	a.BaseControll.DoUpdate(c, adEvent)
}
