package model

import (
	"EK-Server/util"

	"github.com/labstack/echo"
)

// Ad 广告位
type Ad struct {
	URL     string `json:"url"`
	EventID int    `json:"event_id"`
	Title   string `json:"title"`
	BaseControll
}

// PointerList PointerList
func (a *Ad) PointerList() interface{} {
	return &[]Ad{}
}

// Pointer Pointer
func (a *Ad) Pointer() interface{} {
	return &Ad{}
}

// TableName TableName
func (a *Ad) TableName() string {
	return TableName("ads")
}

// Add Add
func (a *Ad) Add(c echo.Context) error {
	ad := &Ad{}

	if err := c.Bind(ad); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	ad.Empty()
	return a.BaseControll.Add(c, ad)
}

// Update Update
func (a *Ad) Update(c echo.Context) error {
	ad := &Ad{}

	if err := c.Bind(ad); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	ad.Empty()
	return a.BaseControll.Update(c, ad)
}
