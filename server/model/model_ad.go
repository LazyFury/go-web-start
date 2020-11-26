package model

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Treblex/go-echo-demo/server/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Ad 广告位
type Ad struct {
	BaseControll
	Title   string `json:"title"`
	Param   string `json:"param" gorm:"comment:'参数，商品id 分类id url'"` //参数，商品id 分类id url
	EventID uint   `json:"event_id"`
	GroupID uint   `json:"group_id"`
	Image   string `json:"image"`
}
type selectAds struct {
	*Ad
	*EmptySystemFiled
	Y string `json:"event_id,omitempty"`
	// X         string `json:"id,omitempty"`
	Event     string `json:"event"`
	GroupName string `json:"group_name,omitempty"`
	Type      int    `json:"type,omitempty"`
}

// PointerList PointerList
func (a *Ad) PointerList() interface{} {
	return &[]selectAds{}
}

// Pointer Pointer
func (a *Ad) Pointer() interface{} {
	return &selectAds{}
}

// TableName TableName
func (a *Ad) TableName() string {
	return TableName("ads")
}

// Joins 查询
func (a *Ad) Joins(db *gorm.DB) *gorm.DB {
	db = db.Select("`title`,`param`,`event_id`,`code`,IFNULL(`event`,'no_event') `event`,`group_id`,`id`,`image`") //`group_name`,`type`
	// 连接事件
	event := &AdEvent{}
	db = db.Joins(fmt.Sprintf("left join (select `id` e_id,`event` from `%s`) t2 on t2.`e_id`=`%s`.`event_id`", event.TableName(), a.TableName()))
	// 连接分类
	// groupName := TableName("ad_groups")
	// db = db.Joins(fmt.Sprintf("left join (select `id` g_id,`name` group_name,`type` from `%s`) t3 on t3.`g_id`=`%s`.`group_id`", groupName, a.TableName()))
	return db
}

// List 列表
func (a *Ad) List(c *gin.Context) {
	groupID := c.Query("group_id")
	if groupID == "" {
		panic("请选择分组")
	}
	c.JSON(http.StatusOK, utils.JSONSuccess("", a.BaseControll.ListWithOutPaging(map[string]interface{}{
		"group_id": groupID,
	})))
}

// Add Add
func (a *Ad) Add(c *gin.Context) {
	ad := &Ad{}

	if err := c.Bind(ad); err != nil {
		panic("参数错误")
	}

	ad.Title = strings.Trim(ad.Title, " ")
	if ad.Title == "" {
		panic("广告位标题不可空")
	}

	if ad.EventID > 0 {
		event := &AdEvent{}
		event.BaseControll.Model = event
		if !event.HasOne(map[string]interface{}{
			"id": ad.EventID,
		}) {
			panic("事件不存在")
		}
	}

	if ad.GroupID == 0 {
		panic("请选择广告位分组")
	}

	adGourp := &AdGroup{}
	db := DB
	if db.Model(adGourp).Where(map[string]interface{}{
		"id": ad.GroupID,
	}).First(adGourp).Error != nil {
		panic("分组不存在")
	}

	// if adGourp.IsSigle {
	// 	if a.HasOne(map[string]interface{}{
	// 		"group_id": ad.GroupID,
	// 	}) {
	// 		return util.JSONErr(c, nil, "此分组为单图广告位，不可继续添加")
	// 	}
	// }

	ad.Empty()
	a.BaseControll.DoAdd(c, ad)
}

// Update  更新
func (a *Ad) Update(c *gin.Context) {
	ad := &Ad{}

	if err := c.Bind(ad); err != nil {
		panic("参数错误")
	}

	ad.Empty()
	a.BaseControll.DoUpdate(c, ad)
}
