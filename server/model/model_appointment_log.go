package model

import (
	"fmt"
	"time"

	"github.com/Treblex/go-echo-demo/server/util"
	"github.com/Treblex/go-echo-demo/server/util/customtype"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
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
func (a *AppointmentLog) Add(c echo.Context) error {
	log := &AppointmentLog{}

	if err := c.Bind(log); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}
	if log.AppointmentID == 0 {
		return util.JSONErr(c, nil, "必须输入活动id")
	}

	if hasOne := a.HasOne(log); hasOne {
		return util.JSONErr(c, nil, "已预约,请不要重复预约")
	}

	uid := c.Get("userId").(float64)
	status := 1
	time := customtype.NumberTime{Time: time.Now()}

	log.UID = uint(uid)
	log.Status = status
	log.AppointmentTime = time

	log.Empty()

	return a.DoAdd(c, log)
}
