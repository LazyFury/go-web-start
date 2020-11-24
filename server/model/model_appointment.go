package model

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Treblex/go-echo-demo/server/util/customtype"
	"gorm.io/gorm"

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

type showAppointment struct {
	*Appointment
	Count     int  `json:"count"`
	Applied   bool `json:"applied"` //已申请
	AmIAuthor bool `json:"am_i_author"`
}

// Pointer Pointer
func (a *Appointment) Pointer() interface{} {
	return &showAppointment{}
}

// PointerList PointerList
func (a *Appointment) PointerList() interface{} {
	return &[]showAppointment{}
}

// TableName TableName
func (a *Appointment) TableName() string {
	return TableName("appointment")
}

// Joins Joins
func (a *Appointment) Joins(db *gorm.DB) *gorm.DB {
	db = db.Select("*")
	log := &AppointmentLog{}
	db = db.Joins(fmt.Sprintf("left join (select count(id) count,appointment_id from `%s` group by `appointment_id`) t1 on `t1`.`appointment_id`=`%s`.`id`", log.TableName(), a.TableName()))
	return db
}

// Result 处理返回值
func (a *Appointment) Result(data interface{}, userID uint) interface{} {
	// 如果是列表
	var _, ok = reflect.ValueOf(data).Elem().Interface().([]showAppointment)

	// 如果是详情
	item, ok := reflect.ValueOf(data).Elem().Interface().(showAppointment)
	if ok {
		item.AmIAuthor = userID == item.UID

		log := &AppointmentLog{}
		log.BaseControll.Model = log
		hasOne := log.HasOne(log)

		item.Applied = hasOne

		return item
	}

	return data
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
	if nofund := db.Table(a.TableName()).Where(where).Find(a.Pointer()).Error != nil; nofund {
		return util.JSONErr(c, nil, "删除失败,没有权限")
	}

	return a.DoDelete(c)
}
