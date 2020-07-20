package model

import "EK-Server/util/customtype"

// Order 订单
type Order struct {
	BaseControll
	Status      int                   `json:"status"`
	PayTime     *customtype.LocalTime `json:"pay_time"`
	ConfirmTime *customtype.LocalTime `json:"confirm_time"`
}
