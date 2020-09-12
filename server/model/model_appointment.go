package model

import (
	"fmt"
	"strings"

	"github.com/Treblex/go-echo-demo/server/util/customtype"

	"github.com/Treblex/go-echo-demo/server/util"
	"github.com/labstack/echo/v4"
)

// Appointment 预约
type Appointment struct {
	BaseControll
	UID          uint                  `json:"uid"`
	TypeID       uint                  `json:"type_id"    gorm:"comment:'1,组局，2，营销活动报名'"` // 类型
	Title        string                `json:"title"      gorm:"NOT NULL"`                //标题
	Desc         string                `json:"desc"       gorm:"type:text;not null"`      // 介绍
	Covers       customtype.Array      `json:"covers" gorm:"type:text"`
	Address      string                `json:"address"`                    // 地址
	StartTime    customtype.NumberTime `json:"start_time" gorm:"not null"` //
	EndTime      customtype.NumberTime `json:"end_time" gorm:"not null"   gorm:"not null"`
	MaxNum       int                   `json:"max_num"    gorm:"default:0"`
	WalkParallel bool                  `json:"walk_parallel" gorm:"default:0"`
}

var _ Model = &Appointment{}

// Pointer Pointer
func (a *Appointment) Pointer() interface{} {
	return &Appointment{}
}

// PointerList PointerList
func (a *Appointment) PointerList() interface{} {
	return &[]Appointment{}
}

// TableName TableName
func (a *Appointment) TableName() string {
	return TableName("appointment")
}

// Add add
func (a *Appointment) Add(c echo.Context) error {
	defer util.APIRcovery(c)
	appointment := &Appointment{}

	if err := c.Bind(appointment); err != nil {
		panic(err)
	}

	appointment.Title = strings.Trim(appointment.Title, " ")
	if appointment.Title == "" {
		panic("标题不可空")
	}
	if appointment.Desc == "" {
		panic("请输入简介")
	}
	if appointment.Address == "" {
		panic("请输入地址")
	}

	userID, _ := c.Get("userId").(float64)

	appointment.UID = uint(userID)

	fmt.Println(userID)

	appointment.Empty()
	return a.DoAdd(c, appointment)
}

// List List
func (a *Appointment) List(c echo.Context) error {
	uid := c.Get("userId").(float64)
	return a.GetList(c, map[string]interface{}{
		"uid": uint(uid),
	})
}

// Delete Delete
func (a *Appointment) Delete(c echo.Context) error {
	id := c.Param("id")
	uid := c.Get("userId").(float64)

	db := DB
	where := map[string]interface{}{
		"id":  id,
		"uid": uid,
	}
	if nofund := db.Table(a.TableName()).Where(where).Find(a.Pointer()).RecordNotFound(); nofund {
		return util.JSONErr(c, nil, "删除失败,没有权限")
	}

	return a.DoDelete(c)
}
