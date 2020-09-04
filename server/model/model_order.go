package model

import "github.com/treblex/go-echo-demo/server/util/customtype"

// Order 订单
type Order struct {
	BaseControll
	Status      int                   `json:"status"`
	PayTime     *customtype.LocalTime `json:"pay_time"`
	ConfirmTime *customtype.LocalTime `json:"confirm_time"`
}

// PointerList PointerList
func (o *Order) PointerList() interface{} {
	return &[]Order{}
}

// Pointer Pointer
func (o *Order) Pointer() interface{} {
	return &Order{}
}

// TableName TableName
func (o *Order) TableName() string {
	return TableName("order")
}

// IsPublic 个人数据
func (o *Order) IsPublic() bool { return false }
