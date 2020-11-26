package model

import (
	"fmt"
	"time"

	"github.com/Treblex/go-echo-demo/server/utils/customtype"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AppointmentLog 预约记录
type AppointmentLog struct {
	BaseControll
	UID             uint                  `json:"uid"`
	AppointmentID   uint                  `json:"appointment_id"`
	Status          int                   `json:"status" gorm:"comment:'0,已取消 1,准备'"`
	AppointmentTime customtype.NumberTime `json:"appointment_time"`
}

type showAppointmentLog struct {
	*BaseControll
	*AppointmentLog
	*Appointment
}

// Pointer Pointer
func (a *AppointmentLog) Pointer() interface{} {
	return &showAppointmentLog{}
}

// PointerList PointerList
func (a *AppointmentLog) PointerList() interface{} {
	return &[]showAppointmentLog{}
}

// TableName TableName
func (a *AppointmentLog) TableName() string {
	return TableName("appointment_log")
}

// Joins Joins
func (a *AppointmentLog) Joins(db *gorm.DB) *gorm.DB {
	db = db.Select("*")
	appointment := &Appointment{}
	db = db.Joins(fmt.Sprintf("left join (select id a_id,`title`,`desc`,`address`,`start_time`,`end_time`,`max_num` from `%s`) t2 on t2.`a_id`=`%s`.`appointment_id`", appointment.TableName(), a.TableName()))
	return db
}

// Add Add
func (a *AppointmentLog) Add(c *gin.Context) {
	log := &AppointmentLog{}

	if err := c.Bind(log); err != nil {
		panic("参数错误")
	}
	if log.AppointmentID == 0 {
		panic("必须输入活动id")
	}

	if hasOne := a.HasOne(log); hasOne {
		panic("已预约,请不要重复预约")
	}

	user := c.MustGet("user").(*User)
	status := 1
	time := customtype.NumberTime{Time: time.Now()}

	log.UID = user.ID
	log.Status = status
	log.AppointmentTime = time

	log.Empty()

	a.DoAdd(c, log)
}
