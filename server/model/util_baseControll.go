package model

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/lazyfury/go-web-template/model"
	"github.com/lazyfury/go-web-template/tools"
	"gorm.io/gorm"
)

// BaseControll 空方法用户数据模型继承方法
type BaseControll model.Model

// SetCode SetCode
func (b *BaseControll) SetCode() (err error) {
	_uuid, err := tools.UUID()
	b.Code = _uuid
	return
}

// Joins Joins
func (b *BaseControll) Joins(db *gorm.DB) *gorm.DB {
	return db
}

// SetUser SetUser
func (b *BaseControll) SetUser(c *gin.Context, data interface{}) error {

	obj := reflect.ValueOf(data).Elem()
	ref := obj.FieldByNameFunc(func(s string) bool {
		fieldNames := []string{"user_id", "to_user"}
		for _, str := range fieldNames {
			if str == s {
				return true
			}
		}
		return false
	})
	user := GetUserOrLogin(c)
	if ref.CanSet() {
		ref.SetUint(uint64(user.ID))
	}

	return nil
}
