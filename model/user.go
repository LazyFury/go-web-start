package model

import (
	"errors"
	"fmt"
	"main/util"
)

// User User
type User struct {
	ID         string         `json:"id"`
	Password   string         `json:"password"`
	Name       string         `json:"name" gorm:"unique"`
	Email      string         `json:"email"`
	IP         string         `json:"ip"`
	Ua         string         `json:"ua"`
	CreateTime util.LocalTime `json:"create_time"`
	LoginTime  util.LocalTime `json:"login_time"`
	Status     int            `json:"status"`
}

// SearchUser	 SearchUser
type searchUser struct {
	ID         string         `json:"id"`
	Email      string         `json:"email"`
	Name       string         `json:"name"`
	IP         string         `json:"ip"`
	Ua         string         `json:"ua"`
	CreateTime util.LocalTime `json:"create_time"`
	LoginTime  util.LocalTime `json:"login_time"`
	Status     int            `json:"status"`
}

// Find 查找用户
func (u *User) Find() (interface{}, error) {
	db := util.DB
	if db.Where(u).First(u).RecordNotFound() {
		return "", nil
	}
	return nil, errors.New("用户名已存在")
}

// GetAllUser  获取所有用户列表
func (u *User) GetAllUser(limit int, page int) map[string]interface{} {
	return util.DataBaselimit(limit, page, &searchUser{}, &[]searchUser{}, "users")
}

// UpdateUser 更新用户
func (u *User) UpdateUser(id string, data *User) (string, error) {
	db := util.DB
	if db.First(&User{ID: id}).RecordNotFound() {
		return "", errors.New("数据不存在")
	}
	row := db.Model(&User{ID: id}).Updates(data)
	if row.Error != nil {
		return "", row.Error
	}
	if row.RowsAffected <= 0 {
		return "", errors.New("没有更改")
	}
	return "", nil
}

// AddUser 添加用户
func (u *User) AddUser() (string, error) {
	db := util.DB

	fmt.Println(u)
	db.NewRecord(u) // => 主键为空返回`true`
	row := db.Create(u)

	if row.Error != nil {
		return "添加失败", row.Error
	}

	if row.RowsAffected <= 0 {
		return "添加失败，没有更改", errors.New("")
	}

	return "添加成功", nil
}

// DelUser 删除用户
func (u *User) DelUser() (interface{}, error) {
	db := util.DB
	// db.Delete(u)
	row := db.Exec("DELETE FROM `users` WHERE `id` = ?", u.ID)
	if row.Error != nil {
		return row.Error, errors.New("删除失败")
	}
	if row.RowsAffected <= 0 {
		return nil, errors.New("删除失败,数据不存在")
	}
	return nil, nil
}
