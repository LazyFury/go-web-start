package model

import "github.com/Treblex/go-echo-demo/server/util/customtype"

// AppointmentLog 预约记录
type AppointmentLog struct {
	BaseControll
	UID             uint                  `json:"uid"`
	AppointmentID   uint                  `json:"appointment_id"`
	Status          int                   `json:"status" gorm:"comment:'0,已取消 1,准备'"`
	AppointmentTime customtype.NumberTime `json:"appointment_time"`
}

// Pointer Pointer
func (a *AppointmentLog) Pointer() interface{} {
	return &AppointmentLog{}
}

// PointerList PointerList
func (a *AppointmentLog) PointerList() interface{} {
	return &[]AppointmentLog{}
}

// TableName TableName
func (a *AppointmentLog) TableName() string {
	return TableName("appointment_log")
}
