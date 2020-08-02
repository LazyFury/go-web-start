package model

import (
	"EK-Server/util"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
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
		Count int `json:"count"`
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

// Joins 统计
func (a *AdGroup) Joins(db *gorm.DB) *gorm.DB {
	db = db.Select("*")
	ad := &Ad{}
	db = db.Joins(fmt.Sprintf("left join (select count(id) `count`,`group_id` from `%s` group by `group_id`) t1 on t1.`group_id`=`%s`.`id`", ad.TableName(), a.TableName()))
	return db
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

	list, _ := ad.BaseControll.ListWithOutPaging(map[string]interface{}{
		"group_id": id,
	}).(*[]selectAds)
	count := len(*list)
	// fmt.Printf("%v\n\n", reflect.TypeOf(list).Elem().Kind())
	// fmt.Printf("%v\n\n", reflect.ValueOf(list).Elem())

	result := &struct {
		*AdGroup
		*EmptySystemFiled
		Count int          `json:"count"`
		List  *[]selectAds `json:"list"`
	}{
		AdGroup: group,
		List:    list,
		Count:   count,
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
	return a.BaseControll.DoAdd(c, adGroup)
}

// Update Update
func (a *AdGroup) Update(c echo.Context) error {
	adGroup := &AdGroup{}

	if err := c.Bind(adGroup); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	adGroup.Empty()
	return a.BaseControll.DoUpdate(c, adGroup)
}
