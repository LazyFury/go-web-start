package model

import (
	"fmt"
	"reflect"
	"strings"

	"gorm.io/gorm"

	"github.com/Treblex/go-echo-demo/server/utils/customtype"
	"github.com/gin-gonic/gin"
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
func (a *Appointment) Add(c *gin.Context) {
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

	user, _ := c.MustGet("userId").(*User)

	appointment.UID = user.ID

	appointment.Empty()
	a.DoAdd(c, appointment)
}

// List List
func (a *Appointment) List(c *gin.Context) {
	user := c.MustGet("user").(*User)
	a.GetList(c, map[string]interface{}{
		"uid": user.ID,
	})
}

// Delete Delete
func (a *Appointment) Delete(c *gin.Context) {
	id := c.Param("id")
	user := c.MustGet("user").(*User)

	db := DB
	where := map[string]interface{}{
		"id":  id,
		"uid": user.ID,
	}
	if nofund := db.Table(a.TableName()).Where(where).Find(a.Pointer()).Error != nil; nofund {
		panic("删除失败,没有权限")
	}

	a.DoDelete(c)
}
