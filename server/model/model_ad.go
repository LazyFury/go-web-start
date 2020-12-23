package model

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Treblex/go-echo-demo/server/utils"
	"github.com/Treblex/go-web-template/xmodel"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Ad 广告位
type Ad struct {
	BaseControll
	Title   string `json:"title" gorm:"not null"`
	Param   string `json:"param" gorm:"comment:'参数，商品id 分类id url';not null"` //参数，商品id 分类id url
	EventID uint   `json:"event_id" gorm:"not null"`
	GroupID uint   `json:"group_id" gorm:"not null"`
	Image   string `json:"image"`
}

type selectAds struct {
	Ad
	// X         string `json:"id,omitempty"`
	Event     string `json:"event" gorm:"->"`
	GroupName string `json:"group_name,omitempty" gorm:"->"`
	Type      int    `json:"type,omitempty" gorm:"->"`
}

// PointerList PointerList
func (a *Ad) PointerList() interface{} {
	return &[]selectAds{}
}

var _ xmodel.Controller = &Ad{}

//Result Result
func (a *Ad) Result(data interface{}) interface{} {
	return data
}

//Validator Validator
func (a *Ad) Validator() error {
	a.Title = strings.Trim(a.Title, " ")
	if a.Title == "" {
		utils.Error("广告位标题不可空")
	}

	if a.EventID == 0 {
		utils.Error("请选择广告位事件")
	}

	if a.EventID > 0 {
		event := &AdEvent{}
		if err := DB.GetObjectOrNotFound(event, map[string]interface{}{
			"id": a.EventID,
		}); err != nil {
			utils.Error("事件不存在")
		}
	}

	if a.GroupID == 0 {
		utils.Error("请选择广告位分组")
	}

	adGourp := &AdGroup{}
	if err := DB.GetObjectOrNotFound(adGourp, map[string]interface{}{
		"id": a.GroupID,
	}); err != nil {
		utils.Error("分组不存在")
	}
	return nil
}

//Object Object
func (a *Ad) Object() interface{} {
	return &selectAds{}
}

//Objects Objects
func (a *Ad) Objects() interface{} {
	return &[]selectAds{}
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
		utils.Error("请选择分组")
	}
	c.JSON(http.StatusOK, utils.JSONSuccess("", a.BaseControll.ListWithOutPaging(map[string]interface{}{
		"group_id": groupID,
	})))
}
