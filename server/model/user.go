package model

import (
	"EK-Server/util"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

// User 用户更新
type User struct {
	ID         int            `json:"id"`
	Password   string         `json:"password"`
	Name       string         `json:"name" gorm:"unique"`
	Email      string         `json:"email"`
	IP         string         `json:"ip"`
	Ua         string         `json:"ua"`
	CreateTime util.LocalTime `json:"create_time"`
	LoginTime  util.LocalTime `json:"login_time"`
	Status     int            `json:"status"`
	AddTime    util.NumberTime
}

// SearchUser	 用户列表显示
type searchUser struct {
	ID         int             `json:"id"`
	Email      string          `json:"email"`
	Name       string          `json:"name"`
	IP         string          `json:"ip"`
	Ua         string          `json:"ua"`
	CreateTime util.LocalTime  `json:"create_time"`
	LoginTime  util.LocalTime  `json:"login_time"`
	Status     int             `json:"status"`
	AddTime    util.NumberTime `json:"AddTime"`
	Bool       bool
}

// WechatOauth 微信用户登陆
type WechatOauth struct {
	gorm.Model
	UID          int    `json:"uid"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
	Nickname     string `json:"nickname"`
	Sex          int    `json:"sex"`
	Headimgurl   string `json:"headimgurl"`
	Province     string `json:"province"`
	City         string `json:"city"`
	Country      string `json:"country"`
	Unionid      string `json:"unionid"`
	Subscribe    int    `json:"subscribe"`
}

// Find 查找用户
func (u *User) Find() error {
	db := DB
	if db.Where(u).First(u).RecordNotFound() {
		return errors.New("用户不存在")
	}
	return nil
}

// GetAllUser  获取所有用户列表
func (u *User) GetAllUser(limit int, page int) map[string]interface{} {
	return DataBaselimit(limit, page, map[string]interface{}{}, &[]searchUser{}, "users", "id desc")
}

// UpdateUser 更新用户
func (u *User) UpdateUser(id int, data *User) error {
	db := DB

	// uid, _ := strconv.Atoi(id)
	// 使用id查找用户
	user := &User{ID: id}
	err := user.Find()
	if err != nil {
		return err
	}
	row := db.Model(&User{ID: id}).Updates(data)
	if row.Error != nil {
		return row.Error
	}
	if row.RowsAffected <= 0 {
		return errors.New("没有更改")
	}
	return nil
}

// AddUser 添加用户
func (u *User) AddUser() (string, error) {
	db := DB

	fmt.Println(u)
	db.NewRecord(u) // => 主键为空返回`true`
	row := db.Create(u)

	if row.Error != nil {
		return "添加失败,用户已存在", nil
	}

	if row.RowsAffected <= 0 {
		return "添加失败，没有更改", errors.New("")
	}

	return "添加成功", nil
}

// DelUser 删除用户
func (u *User) DelUser() (interface{}, error) {
	db := DB
	row := db.Delete(u)
	// row := db.Exec("DELETE FROM "+config.Global.TablePrefix+"`users` WHERE `id` = ?", u.ID)
	if row.Error != nil {
		return row.Error, errors.New("删除失败")
	}
	if row.RowsAffected <= 0 {
		return nil, errors.New("删除失败,数据不存在")
	}
	return nil, nil
}
