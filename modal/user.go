package modal

import (
	"fmt"
	"main/util"
)

// User User
type User struct {
	ID         string         `json:"id"`
	Password   string         `json:"password"`
	Name       string         `json:"name"`
	IP         string         `json:"ip"`
	Ua         string         `json:"ua"`
	CreateTime util.LocalTime `json:"create_time"`
	LoginTime  util.LocalTime `json:"login_time"`
	Status     int            `json:"status"`
}

type searchUser struct {
	ID string `json:"id"`
	// Password   string    `json:"password" gorm:"-"`
	Name       string         `json:"name"`
	IP         string         `json:"ip"`
	Ua         string         `json:"ua"`
	CreateTime util.LocalTime `json:"create_time"`
	LoginTime  util.LocalTime `json:"login_time"`
	Status     int            `json:"status"`
}

// GetAllUser  获取所有用户列表
func (obj *User) GetAllUser(limit int, page int) map[string]interface{} {
	db := util.DB
	// 用户列表
	users := []searchUser{}
	// 初始化数据库对象
	userModal := db.Table("users").Model(&searchUser{}).Omit("password")
	// 总数
	var count int
	// 绑定总数
	userModal.Count(&count)
	// 查询绑定用户列表
	userModal.Offset(limit*(page-1)).Limit(limit).Find(&users).Order("name", false)

	// m := map[string]string{}
	// for i, v := range users {
	// 	fmt.Println(i, v)
	// }

	return map[string]interface{}{
		"count":    count,
		"list":     users,
		"pageSize": limit,
		"page":     page,
	}
}

// UpdateUser 更新用户
func (obj *User) UpdateUser(id string, data *User) (string, error) {
	db := util.DB
	db.Model(&User{ID: id}).Updates(data)
	err := db.Error

	if err != nil {
		return "", err
	}

	fmt.Println(data)

	fmt.Println(db.RowsAffected)
	return "保存成功", nil
}

// AddUser 添加用户
func (user *User) AddUser() (string, error) {
	db := util.DB

	fmt.Println(user)
	db.NewRecord(user) // => 主键为空返回`true`
	db.Create(user)
	return "-1", nil
}

// DelUser 删除用户
func (user *User) DelUser() (string, error) {
	db := util.DB
	db.Delete(user)
	return "1", nil
}
