package model

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/lazyfury/go-web-start/server/utils"
	"github.com/lazyfury/go-web-template/model"

	"gorm.io/gorm"
)

// AdGroup 广告位
type AdGroup struct {
	BaseControll
	// IsSigle  bool   `json:"is_sigle" gorm:"default:false;comment:'true单图,false多图'"`
	Name     string `json:"name" gorm:"unique;not null"`
	Desc     string `json:"desc" gorm:"type:text;comment:'描述';"`
	MaxCount int    `json:"max_count" gorm:"default:1"`
}
type selectAdGroup struct {
	AdGroup
	Count int          `json:"count" gorm:"->"`
	List  *[]selectAds `json:"list,omitempty" gorm:"-"`
}

//Validator Validator
func (a *AdGroup) Validator() error {
	a.Name = strings.Trim(a.Name, " ")
	if a.Name == "" {
		utils.Error("分组标题不可空")
	}
	return nil
}

//Object Object
func (a *AdGroup) Object() interface{} {
	return &selectAdGroup{}
}

//Objects Objects
func (a *AdGroup) Objects() interface{} {
	return &[]selectAdGroup{}
}

//Result Result
func (a *AdGroup) Result(data interface{}) interface{} {
	var obj = reflect.ValueOf(data).Elem().Interface()

	if _obj, ok := obj.(selectAdGroup); ok {
		_obj.List = &[]selectAds{}
		ad := &selectAds{}
		adsModel := DB.GetObjectsOrEmpty(_obj.List, map[string]interface{}{
			"group_id": _obj.ID,
		}, ad.Joins)
		if err := adsModel.All(); err == nil {
			_obj.List, _ = ad.Result(_obj.List).(*[]selectAds)
			return _obj
		}
	}
	return data
}

var _ model.Controller = &AdGroup{}

// TableName TableName
func (a *AdGroup) TableName() string {
	return TableName("ad_groups")
}

// Joins 统计
func (a *AdGroup) Joins(db *gorm.DB) *gorm.DB {
	ad := &Ad{}
	db = db.Joins(fmt.Sprintf("left join (select count(id) `count`,`group_id` from `%s` where `deleted_at`  IS NULL  group by `group_id`) t1 on t1.`group_id`=`%s`.`id`", ad.TableName(), a.TableName()))
	return db
}
