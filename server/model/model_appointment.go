package model

import (
	"github.com/Treblex/go-echo-demo/server/util/customtype"
)

// Appointment 预约
type Appointment struct {
	BaseControll
	TypeID    uint                 `json:"type_id"    gorm:"comment:'1,组局，2，营销活动报名'"` // 类型
	Title     string               `json:"title"      gorm:"not null"`                //标题
	Desc      string               `json:"desc"       gorm:"type:text;not null"`      // 介绍
	Address   string               `json:"address"`                                   // 地址
	StartTime customtype.LocalTime `json:"start_time" gorm:"not null"`                //
	EndTime   customtype.LocalTime `json:"end_time"   gorm:"not null"`
	MaxNum    int                  `json:"max_num"    gorm:"default:0"`
}

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
