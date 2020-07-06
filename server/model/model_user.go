package model

import (
	"EK-Server/util"
	"EK-Server/util/customtype"
	"EK-Server/util/sha"
	"EK-Server/util/wechat"
	"errors"
	"strings"
	"time"

	"github.com/labstack/echo"
)

// User 用户更新
type User struct {
	BaseControll
	Password  string               `json:"password" gorm:"not null"`
	Name      string               `json:"name" gorm:"unique;not null"`
	Email     string               `json:"email"`
	IP        string               `json:"ip"`
	Ua        string               `json:"ua"`
	LoginTime customtype.LocalTime `json:"login_time"`
	Status    int                  `json:"status"`
	IsAdmin   bool                 `json:"is_admin" gorm:"default:0"`
}

// WechatOauth 微信用户登陆
type WechatOauth struct {
	BaseControll
	UID int `json:"user_id"`
	*wechat.UserInfo
}

// PointerList 列表
func (u *User) PointerList() interface{} {
	type tmp struct {
		*User
		Password string `json:"-"`
	}
	return &[]tmp{}
}

// Pointer 实例
func (u *User) Pointer() interface{} {
	return &User{}
}

// TableName 表名
func (u *User) TableName() string {
	return TableName("users")
}

// Add 添加
func (u *User) Add(c echo.Context) error {
	user := new(User)

	if err := c.Bind(user); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	user.Name = strings.Trim(user.Name, " ")
	if user.Name == "" {
		return util.JSONErr(c, nil, "用户名不可空")
	}
	if user.Password == "" {
		return util.JSONErr(c, nil, "用户密码不可空")
	}

	user.Password = sha.EnCode(user.Password)

	req := c.Request()
	ua := req.UserAgent()
	ip := util.ClientIP(c)
	user.IP = ip
	user.Ua = ua
	user.LoginTime = customtype.LocalTime{Time: time.Now()}
	user.Status = 1

	user.Empty()
	return u.BaseControll.Add(c, user)
}

// RegController AddUser
func (u *User) RegController(c echo.Context) error {

	return u.Add(c)
}

// Update 更新
func (u *User) Update(c echo.Context) error {
	user := new(User)

	if err := c.Bind(user); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	_u := &User{BaseControll: BaseControll{ID: uint(user.ID)}}
	err := _u.HasUser()
	if err != nil {
		return util.JSONErr(c, nil, err.Error())
	}

	user.Name = strings.Trim(user.Name, " ")
	if user.Password != "" {
		user.Password = sha.EnCode(user.Password)
	}

	user.Empty()
	return u.BaseControll.Update(c, user)
}

// HasUser 查找用户
func (u *User) HasUser() error {
	db := DB
	if db.Where(u).First(u).RecordNotFound() {
		return errors.New("用户不存在")
	}
	return nil
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
