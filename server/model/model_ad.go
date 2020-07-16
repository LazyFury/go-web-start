package model

import (
	"EK-Server/util"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// Ad 广告位
type Ad struct {
	BaseControll
	Title   string `json:"title"`
	Param   string `json:"param" gorm:"comment:'参数，商品id 分类id url'"` //参数，商品id 分类id url
	EventID uint   `json:"event_id"`
	GroupID uint   `json:"group_id"`
}

// PointerList PointerList
func (a *Ad) PointerList() interface{} {
	return &[]struct {
		*Ad
		*EmptySystemFiled
		Y         string `json:"event_id,omitempty"`
		X         string `json:"id,omitempty"`
		Event     string `json:"event"`
		GroupName string `json:"group_name,omitempty"`
		Type      int    `json:"type,omitempty"`
	}{}
}

// Pointer Pointer
func (a *Ad) Pointer() interface{} {
	return &Ad{}
}

// TableName TableName
func (a *Ad) TableName() string {
	return TableName("ads")
}

// Joins 查询
func (a *Ad) Joins(db *gorm.DB) *gorm.DB {
	db = db.Select("`title`,`param`,`event_id`,IFNULL(`event`,'no_event') `event`,`group_id`") //`group_name`,`type`
	// 连接事件
	eventName := TableName("ad_events")
	db = db.Joins(fmt.Sprintf("left join (select `id` e_id,`event` from `%s`) t2 on t2.`e_id`=`%s`.`event_id`", eventName, a.TableName()))
	// 连接分类
	// groupName := TableName("ad_groups")
	// db = db.Joins(fmt.Sprintf("left join (select `id` g_id,`name` group_name,`type` from `%s`) t3 on t3.`g_id`=`%s`.`group_id`", groupName, a.TableName()))
	return db
}

// List 列表
func (a *Ad) List(c echo.Context) error {
	groupID := c.QueryParam("group_id")
	if groupID == "" {
		return util.JSONErr(c, nil, "请选择分组")
	}
	return util.JSONSuccess(c, a.BaseControll.ListWithOutPaging(map[string]interface{}{
		"group_id": groupID,
	}), "")
}

// Add Add
func (a *Ad) Add(c echo.Context) error {
	ad := &Ad{}

	if err := c.Bind(ad); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	ad.Title = strings.Trim(ad.Title, " ")
	if ad.Title == "" {
		return util.JSONErr(c, nil, "广告位标题不可空")
	}

	if ad.GroupID == 0 {
		return util.JSONErr(c, nil, "请选择广告位分组")
	}

	adGourp := &AdGroup{}
	adGourp.BaseControll.Model = adGourp
	if !adGourp.BaseControll.HasOne(map[string]interface{}{
		"id": ad.GroupID,
	}) {
		return util.JSONErr(c, nil, "分组不存在")
	}

	ad.Empty()
	return a.BaseControll.Add(c, ad)
}

// Update  更新
func (a *Ad) Update(c echo.Context) error {
	ad := &Ad{}

	if err := c.Bind(ad); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	ad.Empty()
	return a.BaseControll.Update(c, ad)
}
