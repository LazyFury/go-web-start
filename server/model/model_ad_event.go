package model

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/lazyfury/go-web-template/model"
	"github.com/lazyfury/go-web-template/response"
	"gorm.io/gorm"
)

// AdEvent banner事件
type AdEvent struct {
	BaseControll
	Event string `json:"event" gorm:"not null;unique_index;default:'no_event';comment:'banner事件,字符串，唯一'"`
	Desc  string `json:"desc" gorm:""`
}

// Validator Validator
func (a *AdEvent) Validator() error {

	a.Event = strings.Trim(a.Event, " ")
	if a.Event == "" {
		response.Error("event定义不可空")
	}

	if match, _ := regexp.MatchString("^[A-Za-z]+$", a.Event); !match {
		response.Error("仅支持英文字符串")
	}

	// 不等于0时为编辑 否则为新增
	if err := DB.GetObjectOrNotFound(&AdEvent{}, map[string]interface{}{
		"event": a.Event,
	}, func(db *gorm.DB) *gorm.DB {
		if a.ID == 0 {
			return db
		}
		return db.Not(map[string]interface{}{
			"id": a.ID,
		})
	}); err == nil {
		response.Error("不可添加,已存在相同的事件")
	}

	return nil
}

// Object Object
func (a *AdEvent) Object() interface{} {
	return &AdEvent{}
}

// Objects Objects
func (a *AdEvent) Objects() interface{} {
	return &[]struct {
		AdEvent
		// *EmptySystemFiled
		Count int `json:"count" gorm:"->"`
	}{}
}

// Result Result
func (a *AdEvent) Result(data interface{}) interface{} {
	return data
}

var _ model.Controller = &AdEvent{}

// TableName TableName
func (a *AdEvent) TableName() string {
	return TableName("ad_events")
}

// Joins Joins
func (a *AdEvent) Joins(db *gorm.DB) *gorm.DB {
	ad := &Ad{}
	db = db.Joins(fmt.Sprintf("left join (select count(id) `count`,`event_id` from `%s` group by `event_id`) t1 on t1.`event_id`=`%s`.`id`", ad.TableName(), a.TableName()))
	return db
}
