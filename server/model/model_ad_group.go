package model

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Treblex/go-echo-demo/server/utils"
	"github.com/gin-gonic/gin"

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

// PointerList PointerList
func (a *AdGroup) PointerList() interface{} {
	return &[]struct {
		*AdGroup
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
	db = db.Joins(fmt.Sprintf("left join (select count(id) `count`,`group_id` from `%s` where `deleted_at`  IS NULL  group by `group_id`) t1 on t1.`group_id`=`%s`.`id`", ad.TableName(), a.TableName()))
	return db
}

// List 列表
func (a *AdGroup) List(c *gin.Context) {
	c.JSON(http.StatusOK, utils.JSONSuccess("", a.BaseControll.ListWithOutPaging(nil)))
}

// Detail 分组详情
func (a *AdGroup) Detail(c *gin.Context) {
	db := DB
	id := c.Param("id")
	if id == "" {
		panic("参数错误")
	}

	group := &AdGroup{}
	if db.Model(group).Where(map[string]interface{}{
		"id": id,
	}).First(group).Error != nil {
		panic("广告位不存在")
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

	c.JSON(http.StatusOK, utils.JSONSuccess("", result))
}

// Add AdGroupd
func (a *AdGroup) Add(c *gin.Context) {
	adGroup := &AdGroup{}

	if err := c.Bind(adGroup); err != nil {
		panic("参数错误")
	}

	adGroup.Name = strings.Trim(adGroup.Name, " ")
	if adGroup.Name == "" {
		panic("分组标题不可空")
	}

	if a.BaseControll.HasOne(map[string]interface{}{
		"name": adGroup.Name,
	}) {
		panic("已存在相同的分类")
	}

	adGroup.Empty()
	a.BaseControll.DoAdd(c, adGroup)
}

// Update Update
func (a *AdGroup) Update(c *gin.Context) {
	adGroup := &AdGroup{}

	if err := c.Bind(adGroup); err != nil {
		panic("参数错误")
	}

	ad := &Ad{GroupID: adGroup.ID}
	ads := []Ad{}
	db := DB
	if err := db.Table(ad.TableName()).Where(ad).Find(&ads).Error; err == nil {
		if len(ads) > adGroup.MaxCount {
			adGroup.MaxCount = len(ads)
		}
	}

	adGroup.Empty()
	a.BaseControll.DoUpdate(c, adGroup)
}
